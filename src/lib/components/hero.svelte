<script>
	import { goto } from "$app/navigation";

    // Placeholder image URL, replace with your actual image path
    const heroImage = "https://medakcess-your-ai-doctor.lovable.app/assets/hero-medical-BXcg5MlF.jpg";
  </script>
  
  <style>
    /* Custom CSS for the template, replicating the original styles */
    .bg-gradient-hero {
      background: linear-gradient(135deg, #001f3f 0%, #0073e6 100%);
    }
  
    /* Keyframes for animations */
    @keyframes fade-in {
      from { opacity: 0; transform: translateY(20px); }
      to { opacity: 1; transform: translateY(0); }
    }
  
    @keyframes slide-in-right {
      from { opacity: 0; transform: translateX(20px); }
      to { opacity: 1; transform: translateX(0); }
    }
  
    @keyframes float-animation {
      0% { transform: translateY(0); }
      50% { transform: translateY(-10px); }
      100% { transform: translateY(0); }
    }
  
    @keyframes bounce-in {
      0%, 20%, 50%, 80%, 100% { transform: translateY(0); }
      40% { transform: translateY(-15px); }
      60% { transform: translateY(-7px); }
    }
  
    /* Apply animations and other styles */
    .animate-fade-in {
      animation: fade-in 1s ease-out;
    }
  
    .animate-slide-in-right {
      animation: slide-in-right 1s ease-out;
    }
  
    .float-animation {
      animation: float-animation 4s ease-in-out infinite;
    }
  
    .animate-bounce-in {
      animation: bounce-in 1s cubic-bezier(0.68, -0.55, 0.27, 1.55);
    }
  
    .shadow-strong {
      box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
    }
  </style>
  
  <section class="min-h-screen bg-gradient-hero flex items-center justify-center px-4 sm:px-6 lg:px-8 pt-16">
    <div class="container mx-auto">
      <div class="grid lg:grid-cols-2 gap-12 items-center">
        <!-- Left Content -->
        <div class="text-center lg:text-left animate-fade-in">
          <div class="inline-flex items-center px-4 py-2 bg-blue-500/20 rounded-full text-blue-300 font-medium text-sm mb-6">
            <!-- Zap Icon -->
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4 mr-2"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>
            AI-Powered Medical Assistant
          </div>
  
          <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold text-white mb-6 leading-tight">
            Your Trusted
            <span class="block text-cyan-400"> Medical Companion</span>
          </h1>
  
          <p class="text-xl text-white/90 mb-8 max-w-2xl">
            MedAkcess provides instant, reliable medical guidance as your first line of healthcare support.
            Get professional insights powered by advanced AI technology.
          </p>
  
          <div class="flex flex-col sm:flex-row gap-4 mb-12">
            <button class="bg-cyan-400 text-gray-900 font-semibold text-lg px-8 py-4 rounded-full shadow-lg hover:bg-cyan-500 transition-colors duration-300" onclick={() => goto('/chat')}>
              Start Consultation
              <!-- ArrowRight Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-5 h-5 ml-2 inline-block"><line x1="5" x2="19" y1="12" y2="12"/><polyline points="12 5 19 12 12 19"/></svg>
            </button>
            <button class="border-2 border-white text-white font-semibold text-lg px-8 py-4 rounded-full shadow-lg hover:bg-white hover:text-blue-900 transition-colors duration-300" onclick={() => goto('#learn-more')} >
              Learn More
            </button>
          </div>
  
          <!-- Trust Indicators -->
          <div class="grid grid-cols-3 gap-8 max-w-md mx-auto lg:mx-0">
            <div class="text-center">
              <!-- Shield Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-8 h-8 text-white mb-2 mx-auto"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>
              <p class="text-white/80 text-sm font-medium">DPA Compliant</p>
            </div>
            <div class="text-center">
              <!-- Users Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-8 h-8 text-white mb-2 mx-auto"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>
              <p class="text-white/80 text-sm font-medium">Doctor Verified</p>
            </div>
            <div class="text-center">
              <!-- Zap Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-8 h-8 text-white mb-2 mx-auto"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>
              <p class="text-white/80 text-sm font-medium">24/7 Available</p>
            </div>
          </div>
        </div>
  
        <!-- Right Content - Hero Image -->
        <div class="relative animate-slide-in-right">
          <div class="card-hero float-animation">
            <img src={heroImage} alt="Modern medical AI technology" class="w-full h-auto rounded-2xl shadow-strong" />
  
            <!-- Floating Stats -->
            <div class="absolute -top-4 -right-4 bg-white/10 backdrop-blur-sm rounded-xl p-4 shadow-strong animate-bounce-in">
              <div class="text-center">
                <div class="text-2xl font-bold text-cyan-400">99.9%</div>
                <div class="text-sm text-gray-300">Accuracy</div>
              </div>
            </div>
  
            <div class="absolute -bottom-4 -left-4 bg-white/10 backdrop-blur-sm rounded-xl p-4 shadow-strong animate-bounce-in" style="animation-delay: 0.2s;">
              <div class="text-center">
                <div class="text-2xl font-bold text-green-400">24/7</div>
                <div class="text-sm text-gray-300">Support</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  