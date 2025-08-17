<script>
    // Placeholder for custom colors to match the theme
    const colors = {
      primary: '#60a5fa', // Blue-400
      foreground: '#111827', // Gray-900 (used for the footer's background)
      background: '#f9fafb', // Gray-50 (used for the footer's text color)
    };
  </script>
  
  <style>
    /* Custom CSS for the template */
    .bg-gradient-primary {
      background-image: linear-gradient(to right, #4c51bf, #667eea);
    }
  
    /*
      The original component uses `text-background` and `bg-foreground`
      which are typically reversed in dark mode themes.
      I've defined the colors here to match that intent.
    */
    .bg-foreground {
      background-color: var(--foreground, #111827);
    }
  
    .text-background {
      color: var(--background, #f9fafb);
    }
  
    .text-background\/80 {
      color: rgba(249, 250, 251, 0.8);
    }
  
    .text-background\/60 {
      color: rgba(249, 250, 251, 0.6);
    }
  
    .border-background\/20 {
      border-color: rgba(249, 250, 251, 0.2);
    }
  
    .hover\:text-primary:hover {
      color: var(--primary, #60a5fa);
    }
  </style>
  
  <footer class="text-background py-16">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
        <!-- Brand -->
        <div class="lg:col-span-2">
          <div class="flex items-center space-x-2 mb-6">
            <div class="w-10 h-10 bg-gradient-primary rounded-xl flex items-center justify-center">
              <!-- Heart Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-6 h-6 text-white"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
            </div>
            <span class="text-2xl font-bold text-white">
              Med<span class="text-primary">Akcess</span>
            </span>
          </div>
          <p class="text-background/80 mb-6 max-w-md">
            Your trusted AI medical assistant providing reliable healthcare guidance
            as your first line of medical support.
          </p>
          <div class="space-y-2">
            <div class="flex items-center gap-3 text-background/80">
              <!-- Mail Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></svg>
              <span>support@medakcess.cloud.co.zm</span>
            </div>
            <div class="flex items-center gap-3 text-background/80">
              <!-- Phone Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
              <span>1-800-MED-ACCESS</span>
            </div>
            <div class="flex items-center gap-3 text-background/80">
              <!-- MapPin Icon -->
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
              <span>Healthcare Innovation Center, Medical District</span>
            </div>
          </div>
        </div>
  
        <!-- Quick Links -->
        <!-- <div>
          <h3 class="font-semibold text-background mb-4">Quick Links</h3>
          <ul class="space-y-2">
            <li><a href="#features" class="text-background/80 hover:text-primary transition-colors">Features</a></li>
            <li><a href="#symptom-checker" class="text-background/80 hover:text-primary transition-colors">Symptom Checker</a></li>
            <li><a href="#" class="text-background/80 hover:text-primary transition-colors">Find Doctors</a></li>
            <li><a href="#" class="text-background/80 hover:text-primary transition-colors">Health Records</a></li>
          </ul>
        </div> -->
  
        <!-- Support -->
        <!-- <div>
          <h3 class="font-semibold text-background mb-4">Support</h3>
          <ul class="space-y-2">
            <li><a href="#" class="text-background/80 hover:text-primary transition-colors">Help Center</a></li>
            <li><a href="#" class="text-background/80 hover:text-primary transition-colors">Privacy Policy</a></li>
            <li><a href="#" class="text-background/80 hover:text-primary transition-colors">Terms of Service</a></li>
            <li><a href="#" class="text-background/80 hover:text-primary transition-colors">DPA Compliance</a></li>
          </ul>
        </div> -->
      </div>
  
      <div class="border-t border-background/20 mt-12 pt-8 text-center">
        <div class="flex flex-col sm:flex-row justify-between items-center gap-4">
          <p class="text-background/60">
            © 2025 MedAkcess. All rights reserved.
          </p>
          <!-- <div class="flex items-center gap-6">
            <span class="text-background/60 text-sm">Certified Medical AI</span>
            <span class="text-background/60 text-sm">DPA Compliant</span>
            <span class="text-background/60 text-sm">ISO 27001</span>
          </div> -->
        </div>
      </div>
    </div>
  </footer>
  