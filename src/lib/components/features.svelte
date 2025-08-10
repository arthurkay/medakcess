<script lang="ts">
	import { goto } from "$app/navigation";

    // Svelte's equivalent of a data array to drive the component's content
    const features = [
      {
        icon: "Stethoscope",
        title: "Symptom Analysis",
        description: "AI-powered symptom checker that provides preliminary diagnosis and recommendations based on your inputs.",
        color: "text-primary"
      },
      {
        icon: "Brain",
        title: "Smart Insights",
        description: "Advanced machine learning algorithms analyze medical data to provide intelligent health insights.",
        color: "text-success"
      },
      {
        icon: "Shield",
        title: "Secure & Private",
        description: "Your medical data is protected with enterprise-grade security and HIPAA compliance.",
        color: "text-warning"
      },
      {
        icon: "Clock",
        title: "24/7 Availability",
        description: "Access medical guidance anytime, anywhere. No appointments needed for initial consultations.",
        color: "text-primary"
      },
      {
        icon: "Users",
        title: "Doctor Network",
        description: "Connect with verified healthcare professionals when you need human expertise.",
        color: "text-success"
      },
      {
        icon: "FileText",
        title: "Health Records",
        description: "Maintain comprehensive digital health records with easy sharing to healthcare providers.",
        color: "text-warning"
      }
    ];
  
    // A helper function to return the SVG icon based on its name
    function getIcon(name: string) {
      switch (name) {
        case "Stethoscope":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z"/><path d="M8 12l2 2"/><path d="M14 12l-2 2"/><path d="M2 17h20"/><path d="M11 3v1a2 2 0 0 0-2 2v1h6v-1a2 2 0 0 0-2-2v-1"/><path d="M12 17a6 6 0 0 0-6 6h12a6 6 0 0 0-6-6z"/></svg>`;
        case "Brain":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z"/><path d="M11 3v1a2 2 0 0 0-2 2v1h6v-1a2 2 0 0 0-2-2v-1"/><path d="M12 17a6 6 0 0 0-6 6h12a6 6 0 0 0-6-6z"/></svg>`;
        case "Shield":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>`;
        case "Clock":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>`;
        case "Users":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>`;
        case "FileText":
          return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><polyline points="14 2 14 8 20 8"/><line x1="16" x2="8" y1="13" y2="13"/><line x1="16" x2="8" y1="17" y2="17"/><line x1="10" x2="8" y1="9" y2="9"/></svg>`;
        default:
          return "";
      }
    }
  </script>
  
  <style>
    /* Custom CSS for the template, replicating the original styles */
    .bg-primary\/10 {
      background-color: rgba(96, 165, 253, 0.1);
    }
    .text-primary {
      color: #60a5fa;
    }
    .text-success {
      color: #4ade80;
    }
    .text-warning {
      color: #fde047;
    }
    .text-muted-foreground {
      color: #9ca3af;
    }
    .text-foreground {
      color: #e5e7eb;
    }
    .bg-background {
      background-color: #111827;
    }
  
    /* Keyframes for animations */
    @keyframes fade-in {
      from { opacity: 0; transform: translateY(20px); }
      to { opacity: 1; transform: translateY(0); }
    }
  
    /* Apply animations and other styles */
    .animate-fade-in {
      animation: fade-in 1s ease-out;
    }
  
    .card-feature {
      background-color: rgba(255, 255, 255, 0.05);
      border-radius: 12px;
      padding: 32px;
      box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
      backdrop-filter: blur(8px);
      transition: transform 0.3s ease-in-out;
    }
    .card-feature:hover {
      transform: translateY(-10px);
    }
    .icon-medical {
      width: 48px;
      height: 48px;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: rgba(255, 255, 255, 0.1);
      border-radius: 9999px;
      padding: 12px;
    }
    .icon-medical svg {
      width: 24px;
      height: 24px;
    }
  
    .card-medical {
      background-color: rgba(255, 255, 255, 0.05);
      border-radius: 12px;
      padding: 32px;
      box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
      backdrop-filter: blur(8px);
    }
  
    /* Custom button styles, replicating the original */
    .btn-medical {
      background-color: #60a5fa;
      color: #111827;
      font-weight: 600;
      padding: 0.75rem 1.5rem;
      border-radius: 9999px;
      transition: background-color 0.3s;
    }
  
    .btn-medical:hover {
      background-color: #3b82f6; /* A slightly darker blue for hover */
    }
  
    .btn-medical-outline {
      background-color: transparent;
      color: #60a5fa;
      border: 2px solid #60a5fa;
      font-weight: 600;
      padding: 0.75rem 1.5rem;
      border-radius: 9999px;
      transition: background-color 0.3s, color 0.3s;
    }
  
    .btn-medical-outline:hover {
      background-color: #60a5fa;
      color: #111827;
    }
  </style>
  
  <section id="features" class="py-20 bg-background">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Section Header -->
      <div class="text-center mb-16 animate-fade-in">
        <div class="inline-flex items-center px-4 py-2 bg-primary/10 rounded-full text-primary font-medium text-sm mb-6">
          Advanced Features
        </div>
        <h2 class="text-3xl sm:text-4xl lg:text-5xl font-bold text-foreground mb-6">
          Comprehensive Medical
          <span class="block text-primary"> AI Solutions</span>
        </h2>
        <p class="text-xl text-muted-foreground max-w-3xl mx-auto">
          Experience the future of healthcare with our advanced AI-powered features
          designed to support both patients and healthcare professionals.
        </p>
      </div>
  
      <!-- Features Grid -->
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {#each features as feature, index}
          <div
            class="card-feature group animate-fade-in"
            style="animation-delay: {index * 0.1}s"
          >
            <div class="icon-medical {feature.color} mb-6 group-hover:scale-110 transition-transform duration-300">
              {@html getIcon(feature.icon)}
            </div>
            
            <h3 class="text-xl font-semibold text-foreground mb-4">
              {feature.title}
            </h3>
            
            <p class="text-muted-foreground leading-relaxed">
              {feature.description}
            </p>
          </div>
        {/each}
      </div>
  
      <!-- CTA Section -->
      <div class="text-center mt-16 animate-fade-in" id="learn-more">
        <div class="card-medical max-w-4xl mx-auto text-center">
          <h3 class="text-2xl sm:text-3xl font-bold text-foreground mb-4">
            Ready to experience the future of healthcare?
          </h3>
          <p class="text-muted-foreground mb-8 text-lg">
            Join thousands of users who trust MedAccess for their healthcare needs.
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <button class="btn-medical text-lg px-8 py-4" onclick={() => goto('/chat')} >
              Start Free Trial
            </button>
            <button class="btn-medical-outline text-lg px-8 py-4" onclick={() => goto('#learn-more')}  >
              Book Consultation
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>