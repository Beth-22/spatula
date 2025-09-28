<template>
  <button 
    @click="$emit('click')"
    class="bg-white rounded-full p-4 shadow-2xl hover:shadow-3xl transition-all duration-300 hover:scale-110 border-2 border-gray-200 z-20 flex items-center justify-center"
    :class="customClass"
  >
    <Icon 
      :name="direction === 'left' ? 'lucide:chevron-left' : 'lucide:chevron-right'" 
      class="h-8 w-8 text-primary-600" 
    />
  </button>
</template>

<script setup>
defineProps({
  direction: {
    type: String,
    default: 'right',
    validator: (value) => ['left', 'right'].includes(value)
  },
  customClass: {
    type: String,
    default: ''
  }
})

defineEmits(['click'])
</script>