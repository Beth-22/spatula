<template>
  <div class="min-h-screen bg-[#F5F5F7]">
    <Navigation />
    <Banner />
     <!-- Divider just above categories (CreatorsSection) -->
  <!-- Divider just above categories -->
<div class="relative flex justify-center -mt-8 mb-4">
  <img 
    src="/assets/images/divider.png" 
    alt="Divider" 
    class="w-90 mx:w-64 lg:w-90 h-auto"
  />
</div>


        <CreatorsSection /> 


    <Footer />
  </div>
</template>

<script setup>
import ArrowSketch from "@/components/ArrowSketch.vue";

import FeaturedRecipes from "@/components/Carousel/FeaturedRecipes.vue";

// Components will be auto-imported by  <ArrowSketch :size="200" color="#0B840DFF" :strokeWidth="10" :animate="true" :rotate="10" />   Nuxt 3
</script>
