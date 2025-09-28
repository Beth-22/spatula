<template>
  <footer class="bg-green-900 text-white py-12">
    <div class="max-w-7xl mx-auto px-4">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <!-- Logo Section -->
        <div>
          <div class="flex items-center mb-4">
            <SpatulaIcon class="h-8 w-8 text-primary-400" />
            <span class="ml-2 text-xl font-serif font-bold">SEASONED</span>
          </div>
          <p class="text-gray-400">Discover and share amazing recipes with food lovers worldwide.</p>
        </div>

        <!-- Quick Links -->
        <div>
          <h4 class="font-semibold mb-4">QUICK LINKS</h4>
          <ul class="space-y-2 text-gray-400">
            <li><NuxtLink to="/" class="hover:text-white transition-colors">Home</NuxtLink></li>
            <li><NuxtLink to="/recipes" class="hover:text-white transition-colors">Recipes</NuxtLink></li>
            <li><NuxtLink to="/shop" class="hover:text-white transition-colors">Shop</NuxtLink></li>
            <li><NuxtLink to="/about" class="hover:text-white transition-colors">About</NuxtLink></li>
          </ul>
        </div>

        <!-- Categories -->
        <div>
          <h4 class="font-semibold mb-4">CATEGORIES</h4>
          <ul class="space-y-2 text-gray-400">
            <li><a href="#" class="hover:text-white transition-colors">Asian Cuisine</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Italian Food</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Desserts</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Healthy Eating</a></li>
          </ul>
        </div>

        <!-- Contact -->
        <div>
          <h4 class="font-semibold mb-4">CONTACT</h4>
          <ul class="space-y-2 text-gray-400">
            <li class="flex items-center">
              <Icon name="lucide:mail" class="h-4 w-4 mr-2" />
              hello@spatula.com
            </li>
            <li class="flex items-center">
              <Icon name="lucide:phone" class="h-4 w-4 mr-2" />
              +1 (555) 123-4567
            </li>
            <li class="flex items-center">
              <Icon name="lucide:map-pin" class="h-4 w-4 mr-2" />
              Food Street, Culinary City
            </li>
          </ul>
        </div>
      </div>

      <!-- Bottom Bar -->
      <div class="border-t border-green-800 mt-8 pt-8 text-center text-gray-400">
        <p>&copy; 2024 Spatula. All rights reserved. Made with <span class="text-red-500">❤️</span> for food lovers.</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import SpatulaIcon from './SpatulaIcon.vue'
</script>