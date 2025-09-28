<template>
  <div class="flex items-center space-x-1">
    <!-- Stars -->
    <button
      v-for="star in 5"
      :key="star"
      @click="handleRating(star)"
      :disabled="loading || !isAuthenticated"
      class="transition-all duration-200 hover:scale-110"
      :class="[disabled ? 'opacity-50 cursor-not-allowed' : '']"
    >
      <Icon
        :icon="star <= currentRating ? 'mdi:star' : 'mdi:star-outline'"
        class="h-5 w-5"
        :class="star <= currentRating ? 'text-yellow-400' : 'text-gray-300'"
      />
    </button>
    
    <!-- Rating info -->
    <div class="ml-2 text-sm text-gray-600">
      <span v-if="ratingCount > 0" class="font-medium">{{ averageRating }}</span>
      <span v-else class="text-gray-400">No ratings</span>
      <span class="text-xs"> ({{ ratingCount }})</span>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  recipeId: {
    type: String,
    required: true
  },
  initialUserRating: {
    type: Number,
    default: 0
  },
  initialAverageRating: {
    type: Number,
    default: 0
  },
  initialRatingCount: {
    type: Number,
    default: 0
  }
})

const { isAuthenticated } = useAuth()
const { setRating, getRatingStatus, loading } = useRatings()

const currentRating = ref(props.initialUserRating)
const averageRating = ref(props.initialAverageRating)
const ratingCount = ref(props.initialRatingCount)
const disabled = computed(() => loading.value || !isAuthenticated.value)

// Load initial rating status
onMounted(async () => {
  const status = await getRatingStatus(props.recipeId)
  currentRating.value = status.userRating
  averageRating.value = status.averageRating
  ratingCount.value = status.ratingCount
})

const handleRating = async (rating) => {
  if (!isAuthenticated.value) {
    navigateTo('/login')
    return
  }

  try {
    const result = await setRating(props.recipeId, rating)
    currentRating.value = rating
    averageRating.value = result.averageRating
    ratingCount.value = result.ratingCount
  } catch (error) {
    console.error('Failed to set rating:', error)
  }
}

// Watch for authentication changes
watch(isAuthenticated, async (newVal) => {
  const status = await getRatingStatus(props.recipeId)
  currentRating.value = status.userRating
  averageRating.value = status.averageRating
  ratingCount.value = status.ratingCount
})
</script>