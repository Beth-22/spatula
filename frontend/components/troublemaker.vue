<template>
  <IconifyIcon :icon="iconName" :class="className" />
</template>

<script setup>
import { Icon as IconifyIcon } from '@iconify/vue'

const props = defineProps({
  name: {
    type: String,
    required: true
  },
  className: {
    type: String,
    default: 'h-4 w-4'
  }
})

const iconName = props.name
</script>