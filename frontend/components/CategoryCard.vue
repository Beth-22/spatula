<template>
  <div 
    class="relative rounded-xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer h-80 group"
    @click="navigateToCategory"
  >
    <!-- Category Image Background -->
    <div class="absolute inset-0">
      <img 
        :src="category.image" 
        :alt="category.name"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
      />
      <!-- Gradient Overlay -->
      <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/30 to-transparent"></div>
    </div>

    <!-- Content -->
    <div class="relative z-10 h-full flex flex-col justify-end p-6 text-white">
      <h3 class="font-bold text-2xl mb-2">{{ category.name }}</h3>
      <p class="text-white/90 text-lg">{{ category.recipeCount }} recipes</p>
    </div>

    <!-- Hover Overlay -->
    <div class="absolute inset-0 bg-primary-500/10 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
  </div>
</template>

<script setup>
const props = defineProps({
  category: {
    type: Object,
    required: true
  }
})

const navigateToCategory = () => {
  navigateTo(`/recipes?category=${props.category.slug}`)
}
</script>