<script setup>
import { HeartIcon as HeartSolid } from "@heroicons/vue/24/solid"
import { HeartIcon as HeartOutline } from "@heroicons/vue/24/outline"
</script>

<template>
  <div class="flex gap-6">
    <!-- Filled red heart -->
    <HeartSolid class="w-10 h-10 text-red-500" />

    <!-- Outlined red heart -->
    <HeartOutline class="w-10 h-10 text-red-500" />
  </div>
</template>
