<template>
  <button
    @click="handleLike"
    :disabled="loading || !isAuthenticated"
    class="flex items-center space-x-1 transition-all duration-200 hover:scale-110"
    :class="[disabled ? 'opacity-50 cursor-not-allowed' : '']"
  >
    <Icon
      :icon="liked ? 'mdi:heart' : 'mdi:heart-outline'"
      class="h-5 w-5 transition-all duration-200"
      :class="liked ? 'text-red-500 scale-110' : 'text-gray-600'"
    />
    <span class="text-sm font-medium text-gray-700">{{ likeCount }}</span>
  </button>
</template>

<script setup>
const props = defineProps({
  recipeId: {
    type: String,
    required: true
  },
  initialLiked: {
    type: Boolean,
    default: false
  },
  initialLikeCount: {
    type: Number,
    default: 0
  }
})

const { isAuthenticated } = useAuth()
const { toggleLike, getLikeStatus, loading } = useLikes()

const liked = ref(props.initialLiked)
const likeCount = ref(props.initialLikeCount)
const disabled = computed(() => loading.value || !isAuthenticated.value)

// Load initial like status
onMounted(async () => {
  if (isAuthenticated.value) {
    const status = await getLikeStatus(props.recipeId)
    liked.value = status.liked
    likeCount.value = status.likeCount
  }
})

const handleLike = async () => {
  if (!isAuthenticated.value) {
    // Redirect to login or show message
    navigateTo('/login')
    return
  }

  try {
    const result = await toggleLike(props.recipeId)
    liked.value = result.liked
    likeCount.value = result.likeCount
  } catch (error) {
    console.error('Failed to toggle like:', error)
  }
}

// Watch for authentication changes
watch(isAuthenticated, async (newVal) => {
  if (newVal) {
    const status = await getLikeStatus(props.recipeId)
    liked.value = status.liked
    likeCount.value = status.likeCount
  } else {
    liked.value = false
  }
})
</script>