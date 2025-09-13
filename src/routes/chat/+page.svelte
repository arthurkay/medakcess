<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import { Textarea } from "$lib/components/ui/textarea";
  import { Card } from "$lib/components/ui/card";
  import { Send, Bot, User, ArrowLeft } from "@lucide/svelte";
  import { goto } from "$app/navigation";
  import { toast } from "svelte-sonner";
  import { onMount } from "svelte";

   let chat: { role: string; content: string }[] = $state([]); // Strongly type chat
  let input = $state('');
  let bottom: HTMLDivElement; // Reference for scrolling
  let currentAssistantMessage = $state(''); // This will hold the streaming message
  let isLoading = $state(false); // Enable loading state for better UX

  async function sendMessage() {
    if (!input.trim()) return; // Prevent sending empty messages

    chat.push({ role: "user", content: input });
    input = ''; // Clear input after sending

    isLoading = true;
    currentAssistantMessage = ''; // Reset current message

    // Simulate AI response with streaming
    const response = await fetch('/api/chat', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(chat)
    });

    if (!response.ok) {
      toast.error('Error fetching AI response:');
      isLoading = false;
      return;
    }

    const reader = response.body?.getReader();
          if (!reader) {
              toast.error("Failed to get reader from response body.");
              isLoading = false;
              return;
          }

          const decoder = new TextDecoder();
          let buffer = ''; // Use a buffer to handle partial JSON lines

          while (true) {
            const { done, value } = await reader.read();
            if (done) break;

            buffer += decoder.decode(value, { stream: true });

            const lines = buffer.split('\n');
            buffer = lines.pop() || '';

            for (const line of lines) {
              if (line.trim().startsWith('data:')) {
                // Remove the "data: " prefix to parse the JSON
                const data = line.substring(5).trim();
                if (data === '[DONE]') {
                  // This signifies the end of the stream
                  break;
                }
                try {
                  const json = JSON.parse(data);
                  // OpenAI's stream sends 'content' within 'delta'
                  // The 'delta' object can be empty at the end of a message
                  if (json.choices && json.choices[0] && json.choices[0].delta) {
                    const delta = json.choices[0].delta;
                    if (delta.content) {
                      currentAssistantMessage += delta.content;
                    }
                  }
                } catch (e) {
                  console.warn("Could not parse chunk as JSON:", data, e);
                }
              }
            }
          }

          // After the stream is complete, add the full assistant message to chat
          chat.push({ role: "assistant", content: currentAssistantMessage });
          currentAssistantMessage = ''; // Clear the streaming message for the next turn
          isLoading = false; // Reset loading state
  }

  onMount(() => {
    // Optionally, you can load initial chat history from local storage or an API
    //const savedChat = localStorage.getItem('chat');
    //if (savedChat) {
    //  chat = JSON.parse(savedChat);
    //}
    chat.push({ role: "assistant", content: "Hello! I'm your MedAkcess AI assistant. I'm here to help you understand your symptoms and provide general medical guidance. Please remember that I cannot replace professional medical advice. How can I assist you today?" }); // Initial message
    bottom.scrollIntoView({ behavior: "smooth" }); // Scroll to bottom on mount
  })

  $effect(() => {
    console.log(`messages changed:`, isLoading);
    bottom.scrollIntoView({ behavior: "smooth" });
  });

  const handleKeyPress = (e: KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      sendMessage();
    }
  };
</script>

<div class="min-h-screen bg-gradient-to-br from-background via-background to-primary/5">
  <div class="sticky top-0 z-40 bg-background/80 backdrop-blur-md border-b border-border">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center space-x-4">
          <button onclick={() => goto('/')} class="flex items-center space-x-2 text-foreground hover:text-primary transition-colors">
            <ArrowLeft class="w-5 h-5" />
            <span>Back to Home</span>
          </button>
        </div>
        <h1 class="text-xl font-semibold text-foreground">Conversational Medical AI</h1>
        <div class="w-24"></div>
      </div>
    </div>
  </div>

  <div class="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <div class="max-w-4xl mx-auto">
      <Card class="mb-6 p-4 bg-accent/50 border-warning/20">
        <div class="flex items-start space-x-3">
          <div class="w-6 h-6 bg-warning/20 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
            <span class="text-warning text-sm font-bold">!</span>
          </div>
          <div class="text-sm text-muted-foreground">
            <strong class="text-foreground">Medical Disclaimer:</strong> This AI assistant provides general health information only and cannot diagnose conditions or replace professional medical advice. Always consult qualified healthcare professionals for medical concerns.
          </div>
        </div>
      </Card>

      <Card class="h-[600px] flex flex-col shadow-elegant">
        <div class="flex-1 overflow-y-auto p-6 space-y-4">
          {#each chat as message}
            <div class={`flex ${message.role === "user" ? "justify-end" : "justify-start"} animate-fade-in`}>
              <div class={`max-w-[80%] flex items-start space-x-3 ${message.role === "user" ? "flex-row-reverse space-x-reverse" : ""}`}>
                <div class={`w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0 ${message.role === "user" ? "bg-primary text-white" : "bg-gradient-primary text-white"}`}>
                  {#if message.role === "user"}
                    <User class="w-5 h-5" />
                  {:else}
                    <Bot class="w-5 h-5" />
                  {/if}
                </div>

                <div class={`rounded-lg p-4 ${message.role === "user" ? "bg-primary text-white" : "bg-accent text-foreground border border-border"}`}>
                  <p class="text-sm leading-relaxed">{message.content}</p>
                  <!-- <span class={`text-xs mt-2 block ${message.role === "user" ? "text-primary-foreground/70" : "text-muted-foreground"}`}>
                    {message.timestamp.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                  </span> -->
                </div>
              </div>
            </div>
          {/each}

          {#if isLoading}
            <div class="flex justify-start animate-fade-in">
              <div class="flex items-start space-x-3">
                <div class="w-10 h-10 bg-gradient-primary rounded-full flex items-center justify-center">
                  <Bot class="w-5 h-5 text-white" />
                </div>
                <div class="bg-accent rounded-lg p-4 border border-border">
                  <div class="flex space-x-1">
                    <div class="w-2 h-2 bg-muted-foreground rounded-full animate-pulse"></div>
                    <div class="w-2 h-2 bg-muted-foreground rounded-full animate-pulse" style="animation-delay: 0.2s;"></div>
                    <div class="w-2 h-2 bg-muted-foreground rounded-full animate-pulse" style="animation-delay: 0.4s;"></div>
                  </div>
                  {#if currentAssistantMessage}
                    <p class="text-sm leading-relaxed">{currentAssistantMessage}</p>
                  {:else}
                    <p class="text-sm leading-relaxed">Thinking...</p>
                  {/if}
                </div>
              </div>
            </div>
          {/if}
          <div class="h-6" bind:this={bottom}></div> <!-- Spacer for smooth scrolling -->

        </div>

        <div class="border-t border-border p-6">
          <div class="flex space-x-4">
            <Textarea
              bind:value={input}
              onkeydown={handleKeyPress}
              placeholder="Describe your symptoms or ask a medical question..."
              class="min-h-[60px] resize-none focus:ring-primary/20"
              disabled={isLoading}
            />
            <Button onclick={sendMessage} disabled={!input.trim() || isLoading} class="btn-medical px-6 h-[60px]">
              <Send class="w-5 h-5" />
            </Button>
          </div>
          <p class="text-xs text-muted-foreground mt-2">
            Press Enter to send, Shift+Enter for new line
          </p>
        </div>
      </Card>
    </div>
  </div>
</div>