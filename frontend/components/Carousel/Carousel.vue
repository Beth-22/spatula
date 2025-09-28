<template>
  <div class="relative w-full overflow-hidden">
    <!-- Full Width Container -->
    <div class="flex justify-between items-center w-full max-w-screen-2xl mx-auto px-8 md:px-16 lg:px-24 mb-0">
      
      <!-- Left Recipe -->
      <div class="flex-1 flex justify-end mr-8 md:mr-16 lg:mr-24">
        <CarouselRecipe 
          :recipe="getRecipeAt(currentIndex - 1)"
          size="medium"
          position="left"
          @click="prevRecipe"
        />
      </div>
      
      <!-- Left Arrow -->
      <div class="mx-4 md:mx-8 lg:mx-12">
        <CarouselArrow 
          direction="left"
          @click="prevRecipe"
        />
      </div>

      <!-- Center Recipe -->
<div class="mx-4 md:mx-8 lg:mx-12 flex-shrink-0 relative z-30">  <!-- ← Add relative z-30 here -->
  <CarouselRecipe 
    :recipe="recipes[currentIndex]"
    size="large"
    position="center"
    :is-active="true"
  />
</div>

      <!-- Right Arrow -->
      <div class="mx-4 md:mx-8 lg:mx-12">
        <CarouselArrow 
          direction="right"
          @click="nextRecipe"
        />
      </div>

      <!-- Right Recipe -->
      <div class="flex-1 flex justify-start ml-8 md:ml-16 lg:ml-24">
        <CarouselRecipe 
          :recipe="getRecipeAt(currentIndex + 1)"
          size="medium"
          position="right"
          @click="nextRecipe"
        />
      </div>
    </div>

   
  </div>
</template>

<script setup>
const props = defineProps({
  recipes: {
    type: Array,
    required: true
  }
})

const currentIndex = ref(0)

const getRecipeAt = (index) => {
  const length = props.recipes.length
  if (index < 0) return props.recipes[length + index]
  if (index >= length) return props.recipes[index - length]
  return props.recipes[index]
}

const nextRecipe = () => {
  if (currentIndex.value < props.recipes.length - 1) {
    currentIndex.value++
  } else {
    currentIndex.value = 0
  }
}

const prevRecipe = () => {
  if (currentIndex.value > 0) {
    currentIndex.value--
  } else {
    currentIndex.value = props.recipes.length - 1
  }
}
</script>

<style scoped>
/* Ensure full width utilization */
.max-w-screen-2xl {
  max-width: 1536px;
}
</style>