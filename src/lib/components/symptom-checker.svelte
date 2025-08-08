  <script lang="ts">
    import { toast } from "svelte-sonner";
  
    // State variables for the form inputs
    let symptoms = "";
    let duration = "";
    let severity = "";
    let isAnalyzing = false;
    let currentAssistantMessage = '';
    let results: { confidence: any; condition: any; urgency: any; recommendations: any; whenToSeekCare: any; } | null = null;
  
    // A helper function to return the SVG icon based on its name
    function getIcon(name: string) {
      switch (name) {
        case "Search":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>`;
        case "AlertCircle":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12" y1="16" y2="16"/></svg>`;
        case "CheckCircle":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 11.08V12a10 10 0 1 1-5.93-8.98"/><path d="M9 11l3 3L22 4"/></svg>`;
        case "Clock":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>`;
        default:
          return "";
      }
    }
  
    // The analysis function, simulating an API call
    const handleAnalysis = async () => {
      isAnalyzing = true;

      const response = await fetch("/api/generate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          "message": symptoms,
          "duration": duration,
          "severity": severity
        })
      })

      if (!response.ok) {
              const errorText = await response.text();
              toast.error("Error: " + (response.statusText || errorText));
              return;
          }

          const reader = response.body?.getReader();
          if (!reader) {
              toast.error("Failed to get reader from response body.");
              return;
          }

          const decoder = new TextDecoder();
          let buffer = ''; // Use a buffer to handle partial JSON lines

          while (true) {
              const { done, value } = await reader.read();
              if (done) break;

              buffer += decoder.decode(value, { stream: true });

              // Process lines from the buffer
              const lines = buffer.split('\n');
              // Keep the last (potentially incomplete) line in the buffer
              buffer = lines.pop() || '';

              for (const line of lines) {
                  if (line.trim()) {
                      try {
                          const json = JSON.parse(line);
                          // Ollama's stream typically sends 'content' within 'message'
                          // or directly as 'response' depending on the API endpoint.
                          // Assuming `json.message.content` as per your observation.
                          if (json.response) {
                              currentAssistantMessage += json.response;
                          }
                      } catch (e: unknown) {
                          // Log the error but don't stop the stream for a single malformed line
                          console.warn("Could not parse chunk as JSON:", line, e);
                      }
                  }
              }
          }
          //currentAssistantMessage = '';
      
      // Simulate API call with a delay
      setTimeout(() => {
        results = {
          condition: "Upper Respiratory Infection",
          confidence: 85,
          urgency: "Low",
          recommendations: [
            "Rest and stay hydrated",
            "Consider over-the-counter pain relievers",
            "Monitor symptoms for 48-72 hours",
            "Consult a doctor if symptoms worsen"
          ],
          whenToSeekCare: "Seek immediate care if experiencing difficulty breathing, high fever (>101.5°F), or severe headache"
        };
        isAnalyzing = false;
      }, 2000);
    };
  </script>
  
  <style>
    /* Custom CSS for the template, replicating the original styles */
    :root {
      --primary: #60a5fa; /* Blue-400 */
      --success: #4ade80; /* Green-400 */
      --warning: #fde047; /* Yellow-400 */
      --foreground: #e5e7eb; /* Gray-200 */
      --muted: #4b5563; /* Gray-600 */
      --muted-foreground: #9ca3af; /* Gray-400 */
      --background: #111827; /* Gray-900 */
    }
  
    .bg-muted\/30 {
      background-color: rgba(75, 85, 99, 0.3);
    }
    .bg-primary\/10 {
      background-color: rgba(96, 165, 250, 0.1);
    }
    .bg-success\/10 {
      background-color: rgba(74, 222, 128, 0.1);
    }
    .bg-warning\/10 {
      background-color: rgba(253, 224, 71, 0.1);
    }
    .border-warning\/20 {
      border-color: rgba(253, 224, 71, 0.2);
    }
  
    .text-primary { color: var(--primary); }
    .text-success { color: var(--success); }
    .text-warning { color: var(--warning); }
    .text-foreground { color: var(--foreground); }
    .text-muted-foreground { color: var(--muted-foreground); }
  
    /* Keyframes for animations */
    @keyframes fade-in {
      from { opacity: 0; transform: translateY(20px); }
      to { opacity: 1; transform: translateY(0); }
    }
    @keyframes slide-in-right {
      from { opacity: 0; transform: translateX(20px); }
      to { opacity: 1; transform: translateX(0); }
    }
    @keyframes spin {
      from { transform: rotate(0deg); }
      to { transform: rotate(360deg); }
    }
  
    /* Apply animations and other styles */
    .animate-fade-in { animation: fade-in 1s ease-out; }
    .animate-slide-in-right { animation: slide-in-right 0.6s ease-out; }
    .animate-spin { animation: spin 1s linear infinite; }
  
    .card-medical {
      background-color: rgba(255, 255, 255, 0.05);
      border-radius: 12px;
      padding: 32px;
      box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
      backdrop-filter: blur(8px);
    }
  
    .btn-medical {
      background-color: var(--primary);
      color: var(--background);
      font-weight: 600;
      padding: 0.75rem 1.5rem;
      border-radius: 9999px;
      transition: background-color 0.3s;
    }
    .btn-medical:hover {
      background-color: #3b82f6; /* A slightly darker blue for hover */
    }
  </style>
  
  <section class="py-20 bg-muted/30" id="symptom-checker">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-12 animate-fade-in">
        <h2 class="text-3xl sm:text-4xl font-bold text-foreground mb-4">
          AI Symptom Checker
        </h2>
        <p class="text-xl text-muted-foreground max-w-2xl mx-auto">
          Describe your symptoms and get preliminary insights powered by medical AI
        </p>
      </div>
  
      <div class="max-w-4xl mx-auto grid lg:grid-cols-2 gap-8">
        <!-- Input Form -->
        <div class="card-medical animate-slide-in-right">
          <div class="space-y-6">
            <div class="card-header">
              <h3 class="card-title text-xl font-semibold flex items-center gap-2">
                {@html getIcon("Search")}
                Describe Your Symptoms
              </h3>
              <p class="card-description text-sm text-muted-foreground">
                Please provide detailed information about what you're experiencing
              </p>
            </div>
            
            <div class="space-y-6 card-content">
              <div>
                <label class="text-sm font-medium text-foreground mb-2 block">
                  What symptoms are you experiencing?
                </label>
                <textarea
                  placeholder="Describe your symptoms in detail (e.g., headache, fever, cough, fatigue...)"
                  bind:value={symptoms}
                  class="min-h-[100px] w-full p-2 bg-muted/20 text-foreground border border-muted/50 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-primary"
                />
              </div>
              
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="text-sm font-medium text-foreground mb-2 block">
                    Duration
                  </label>
                  <input
                    type="text"
                    placeholder="e.g., 2 days"
                    bind:value={duration}
                    class="w-full p-2 bg-muted/20 text-foreground border border-muted/50 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-primary"
                  />
                </div>
                <div>
                  <label class="text-sm font-medium text-foreground mb-2 block">
                    Severity (1-10)
                  </label>
                  <input
                    type="text"
                    placeholder="e.g., 7"
                    bind:value={severity}
                    class="w-full p-2 bg-muted/20 text-foreground border border-muted/50 rounded-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-primary"
                  />
                </div>
              </div>
  
              <button
                class="btn-medical w-full flex items-center justify-center gap-2"
                on:click={handleAnalysis}
                disabled={!symptoms || isAnalyzing}
              >
                {#if isAnalyzing}
                  {@html getIcon("Clock")}
                  Analyzing...
                {:else}
                  {@html getIcon("Search")}
                  Analyze Symptoms
                {/if}
              </button>
              
              <div class="text-xs text-muted-foreground text-center flex items-center justify-center gap-1">
                {@html getIcon("AlertCircle")}
                This is not a substitute for professional medical advice
              </div>
            </div>
          </div>
        </div>
        
        <!-- Results -->
        <div class="card-medical animate-slide-in-right" style="animation-delay: 0.2s;">
          <div class="card-header">
            <h3 class="card-title text-xl font-semibold flex items-center gap-2">
              {@html getIcon("CheckCircle")}
              Analysis Results
            </h3>
            <p class="card-description text-sm text-muted-foreground">
              AI-powered preliminary assessment
            </p>
          </div>
          <div class="card-content mt-6">
            {#if currentAssistantMessage != ''}
              <div class="text-sm text-muted-foreground mb-4">
                {currentAssistantMessage}
              </div>
            {/if}
            {#if !results}
              <div class="text-center py-12 text-muted-foreground">
                {@html getIcon("Search")}
                <p class="mt-4">Enter your symptoms above to get an AI analysis</p>
              </div>
            <!--  -->
            {/if}
          </div>
        </div>
      </div>
    </div>
  </section>