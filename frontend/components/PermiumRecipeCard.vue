<template>
  <div 
    class="relative bg-white rounded-xl shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer group overflow-hidden h-96 w-64 mx-auto"
    @click="handleCardClick"
  >
    <!-- Recipe Image Background -->
    <div class="absolute inset-0">
      <img 
        :src="getRecipeImage(recipe)" 
        :alt="recipe.title"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/30 to-transparent"></div>
    </div>

    <!-- Premium Crown -->
    <div v-if="recipe.is_premium" 
      class="absolute top-3 left-3 bg-black/30 backdrop-blur-sm rounded-full p-2 shadow-lg z-20 border border-white/20"
    >
      <Icon icon="mdi:crown" class="h-4 w-4 text-yellow-400" />
    </div>

    <!-- Price Badge -->
    <div v-if="recipe.is_premium" class="absolute top-3 right-3 bg-gradient-to-r from-orange-500 to-red-500 text-white text-sm px-3 py-1 rounded-full font-semibold shadow-lg z-20">
      ${{ recipe.price }}
    </div>

    <!-- Content Bottom -->
    <div class="absolute bottom-0 left-0 right-0 p-4 text-white z-10">
      <h3 class="font-bold text-lg mb-2 line-clamp-2 leading-tight">{{ recipe.title }}</h3>
      
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <!-- Rating (commented out since we don't have the data) -->
          <!-- <div class="flex items-center">
            <Icon icon="mdi:star" class="h-3 w-3 text-yellow-400 mr-1" />
            <span class="text-xs font-medium">{{ getAverageRating(recipe) }}</span>
          </div> -->
          <!-- Likes (commented out since we don't have the data) -->
          <!-- <div class="flex items-center">
            <Icon icon="mdi:heart" class="h-3 w-3 text-red-400 mr-1" />
            <span class="text-xs font-medium">{{ getLikesCount(recipe) }}</span>
          </div> -->
        </div>
        
        <!-- Simple Button -->
        <button 
          @click.stop="handleButtonClick"
          :class="purchased ? 'bg-green-500 hover:bg-green-600' : 'bg-gradient-to-r from-orange-500 to-red-500 hover:from-orange-600 hover:to-red-600'"
          class="text-white text-xs px-3 py-2 rounded-lg font-semibold transition-all transform hover:scale-105 shadow-lg"
        >
          {{ purchased ? 'Bought Recipe' : 'Buy Recipe' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  recipe: {
    type: Object,
    required: true
  },
  purchased: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['purchase'])

// Get recipe image - handles both old and new schema
const getRecipeImage = (recipe) => {
  // If recipe has featured_image (old schema)
  if (recipe.featured_image) {
    return recipe.featured_image
  }
  
  // If recipe has recipe_images array (new schema)
  if (recipe.recipe_images && recipe.recipe_images.length > 0) {
    return recipe.recipe_images[0].image_url
  }
  
  // Fallback image
  return '/images/recipe-placeholder.jpg'
}

// Handle card click
const handleCardClick = () => {
  if (props.purchased) {
    navigateTo(`/recipes/${props.recipe.id}`)
  } else if (props.recipe.is_premium) {
    emit('purchase', props.recipe)
  } else {
    navigateTo(`/recipes/${props.recipe.id}`)
  }
}

// Handle button click
const handleButtonClick = () => {
  if (props.purchased) {
    navigateTo(`/recipes/${props.recipe.id}`)
  } else {
    emit('purchase', props.recipe)
  }
}
</script>