<template>
  <div class="relative group">
    <div class="relative" :class="isActive ? 'w-64 h-64' : 'w-52 h-52'">
      <!-- Circular Image Container -->
      <div 
        class="w-full h-full rounded-full overflow-hidden border-4 border-white shadow-2xl transition-all duration-500"
        :class="isActive ? 'border-4' : 'border-2'"
      >
        <img 
          :src="recipe.image" 
          :alt="recipe.name"
          class="w-full h-full object-cover"
        />
      </div>
      
     <!-- Rotating Text Circle -->
<div
  class="absolute inset-0 z-50"
  v-if="isActive"
>
  <svg
    class="w-full h-full transform group-hover:rotate-180 transition-transform duration-1000"
    viewBox="0 0 100 100"
  >
    <defs>
      <path :id="'circlePath' + index" d="M 50, 50 m -45, 0 a 45,45 0 1,1 90,0 a 45,45 0 1,1 -90,0" />
    </defs>
    <text class="text-[5px] font-bold fill-primary-600">
      <textPath :href="'#circlePath' + index" startOffset="0%">
        {{ recipe.name.toUpperCase() }} • {{ recipe.name.toUpperCase() }} •
      </textPath>
    </text>
  </svg>
</div>


      <!-- Premium Crown -->
      <div 
        v-if="recipe.premium" 
        class="absolute -top-2 -right-2 bg-yellow-400 rounded-full p-2 shadow-lg"
        :class="isActive ? 'p-2' : 'p-1'"
      >
        <Icon name="lucide:crown" :class="isActive ? 'h-5 w-5' : 'h-4 w-4'" class="text-white" />
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  recipe: {
    type: Object,
    required: true
  },
  index: {
    type: Number,
    required: true
  },
  isActive: {
    type: Boolean,
    default: false
  }
})
</script>

<style scoped>
text {
  font-size: 5px;
  font-weight: bold;
  letter-spacing: 1px;
}
</style>