<script lang="ts">
  import { fade } from 'svelte/transition';
</script>

<div class="prose mx-auto px-4 py-8" in:fade>
  <h1 class="text-title-xl text-primary">About Us</h1>
  
  <div class="mt-6">
    <p class="text-lg">
      Welcome to our about page. This is where you can share information about your project,
      company, or personal background.
    </p>
    
    <h2 class="text-title-lg mt-8 mb-4">Our Mission</h2>
    <p>
      Add your mission statement or key objectives here. The page is styled using Tailwind's
      typography plugin and your custom design system.
    </p>
    
    <h2 class="text-title-lg mt-8 mb-4">Contact</h2>
    <p>
      Feel free to reach out to us at <a href="mailto:example@email.com" class="text-secondary hover:underline">
      example@email.com</a>
    </p>
  </div>
</div> 