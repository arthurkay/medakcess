<script>
    import { fade } from 'svelte/transition';
  
    // Svelte's reactive state for toggling the mobile menu
    let isMenuOpen = false;
  
    // Function to toggle the menu and close it when a link is clicked
    function toggleMenu() {
      isMenuOpen = !isMenuOpen;
    }
  
    function closeMenu() {
      isMenuOpen = false;
    }
  
    // Placeholder for custom colors to match the theme
    const colors = {
      primary: '#60a5fa', // Blue-400
      accent: '#4ade80',  // Green-400
      foreground: '#f9fafb', // Gray-50
      background: '#111827', // Gray-900
      border: '#374151', // Gray-700
    };
  </script>
  
  <style>
    /* Custom CSS for the template */
    .bg-gradient-primary {
      background-image: linear-gradient(to right, #4c51bf, #667eea);
    }
  
    .shadow-soft {
      box-shadow: 0 4px 14px 0 rgba(0, 0, 0, 0.1);
    }
  
    /* Keyframes for the mobile menu animation */
    @keyframes fade-in {
      from { opacity: 0; transform: translateY(-20px); }
      to { opacity: 1; transform: translateY(0); }
    }
  
    .animate-fade-in {
      animation: fade-in 0.3s ease-out;
    }
  
    /* Custom button styles, replicating the original */
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
  
    /* Custom ghost button style */
    .btn-ghost {
      background-color: transparent;
      color: var(--foreground);
      transition: color 0.3s;
    }
  
    .btn-ghost:hover {
      color: var(--primary);
    }
  </style>
  
  <header class="fixed top-0 left-0 right-0 z-50 bg-gray-900/80 backdrop-blur-md border-b border-gray-700">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo -->
        <div class="flex items-center space-x-2">
          <div class="w-10 h-10 bg-gradient-primary rounded-xl flex items-center justify-center shadow-soft">
            <!-- Heart Icon -->
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-6 h-6 text-white"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
          </div>
          <span class="text-2xl font-bold text-gray-50">
            Med<span class="text-blue-400">Akcess</span>
          </span>
        </div>
  
        <!-- Desktop Navigation -->
       <!--  <nav class="hidden md:flex items-center space-x-8">
          <a href="#features" class="text-gray-50 hover:text-blue-400 transition-colors">
            Features
          </a>
          <a href="/conversation" class="text-gray-50 hover:text-blue-400 transition-colors">
            AI Chat
          </a>
          <a href="#about" class="text-gray-50 hover:text-blue-400 transition-colors">
            About
          </a>
          <a href="#contact" class="text-gray-50 hover:text-blue-400 transition-colors">
            Contact
          </a>
        </nav> -->
  
        <!-- Desktop CTA -->
        <div class="hidden md:flex items-center space-x-4">
          <button class="btn-ghost text-lg px-6 py-2">
            Sign In
          </button>
          <button class="btn-medical text-lg px-6 py-2">
            Get Started
          </button>
        </div>
  
        <!-- Mobile Menu Button -->
        <!-- <button on:click={toggleMenu} class="md:hidden p-2 rounded-lg hover:bg-gray-700 transition-colors">
          {#if isMenuOpen} -->
            <!-- X Icon -->
           <!--  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-6 h-6 text-gray-50"><line x1="18" x2="6" y1="6" y2="18"/><line x1="6" x2="18" y1="6" y2="18"/></svg>
          {:else} -->
            <!-- Menu Icon -->
          <!--   <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-6 h-6 text-gray-50"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
          {/if}
        </button> -->
      </div>
  
      <!-- Mobile Navigation -->
      {#if isMenuOpen}
        <div class="md:hidden py-4 animate-fade-in" transition:fade>
          <nav class="flex flex-col space-y-4">
            <a on:click={closeMenu} href="#features" class="text-gray-50 hover:text-blue-400 transition-colors py-2">
              Features
            </a>
            <a on:click={closeMenu} href="/conversation" class="text-gray-50 hover:text-blue-400 transition-colors py-2">
              AI Chat
            </a>
            <a on:click={closeMenu} href="#about" class="text-gray-50 hover:text-blue-400 transition-colors py-2">
              About
            </a>
            <a on:click={closeMenu} href="#contact" class="text-gray-50 hover:text-blue-400 transition-colors py-2">
              Contact
            </a>
            <div class="flex flex-col space-y-2 pt-4">
              <button class="btn-ghost justify-start">
                Sign In
              </button>
              <button class="btn-medical justify-start">
                Get Started
              </button>
            </div>
          </nav>
        </div>
      {/if}
    </div>
  </header>