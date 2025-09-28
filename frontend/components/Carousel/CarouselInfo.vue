<template>
  <div class="text-center mt-20 mb-0 px-4">
    <h3 class="text-3xl md:text-3xl font-serif font-bold text-gray-900 mb-6">
      {{ recipe.name }}
    </h3>
    <p class="text-gray-600 text-xl md:text-1xl mb-8 max-w-4xl mx-auto leading-relaxed">
      {{ recipe.description }}
    </p>
    <div class="flex justify-center items-center space-x-12 text-md text-gray-600">
      <span class="flex items-center">
        <Icon name="lucide:clock" class="h-7 w-7 mr-3" /> 
        {{ recipe.time }}
      </span>
     
      <span class="flex items-center text-yellow-600 font-semibold">
        <Icon name="lucide:star" class="h-7 w-7 mr-3" /> 
        {{ recipe.rating }}
      </span>
      <span v-if="recipe.premium" class="bg-yellow-400 text-white px-4 py-2 rounded-full text-base font-medium">
        PREMIUM RECIPE
      </span>
    </div>
  </div>
</template>

<script setup>
defineProps({
  recipe: {
    type: Object,
    required: true
  }
})
</script>