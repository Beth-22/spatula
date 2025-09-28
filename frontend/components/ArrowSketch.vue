<script setup lang="ts">
defineProps({
  size: { type: Number, default: 120 },         // px
  color: { type: String, default: '#8b5cf6' },  // lavender
  strokeWidth: { type: Number, default: 6 },
  rotate: { type: Number, default: 0 }          // degrees rotation
});
</script>

<template>
  <svg
    :width="size"
    :height="Math.round(size * 0.6)"
    viewBox="0 0 140 84"
    xmlns="http://www.w3.org/2000/svg"
    :style="{ transform: `rotate(${rotate}deg)` }"
    aria-hidden="true"
  >
    <!-- smooth arrow curve -->
    <path
      d="M10 60 C40 10, 90 10, 130 40"
      :stroke="color"
      :stroke-width="strokeWidth"
      stroke-linecap="round"
      stroke-linejoin="round"
      fill="none"
    />
    <!-- arrow head -->
    <path
      d="M130 40 L118 28 M130 40 L118 52"
      :stroke="color"
      :stroke-width="strokeWidth"
      stroke-linecap="round"
      stroke-linejoin="round"
      fill="none"
    />
  </svg>
</template>
