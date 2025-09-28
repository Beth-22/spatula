<script setup>
import { Icon } from '@iconify/vue'
</script>

<template>
  <!-- Existing -->
  <Icon icon="mdi:home" class="text-xl text-green-600" />
  <Icon icon="fa6-solid:user" class="text-xl text-gray-700" />

  <!-- New -->
  <Icon icon="mdi:star" class="text-xl text-yellow-500" />        <!-- Star -->
  <Icon icon="mdi:bookmark" class="text-xl text-blue-500" />     <!-- Bookmark -->
  <Icon icon="mdi:arrow-right" class="text-xl text-red-500" />   <!-- Arrow -->
  <Icon icon="fa6-solid:star" class="text-yellow-500" />
<Icon icon="fa6-solid:bookmark" class="text-blue-500" />
<Icon icon="fa6-solid:arrow-right" class="text-red-500" />
<div className="text-purple-600 drop-shadow-md">
    <Github size={60} />

  <Github size={50} strokeWidth={1.5} className="animate-pulse" />
</div>
  <Icon 
  icon="mdi:heart" 
  class="text-pink-400 hover:text-red-500 transition-transform transform hover:scale-125"
  width="60" 
  height="60" 
/>



</template>
