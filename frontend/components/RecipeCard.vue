<!-- components/RecipeCard.vue -->
<template>
  <div class="relative bg-white rounded-xl shadow-lg transition-all duration-500 cursor-pointer group overflow-hidden h-96 mx-auto"
       :class="isHovered ? 'w-80 scale-105' : 'w-64'"
       @mouseenter="isHovered = true"
       @mouseleave="isHovered = false"
       @click="navigateToRecipeDetail">
    
    <!-- Recipe Image Background -->
    <div class="absolute inset-0">
      <img
        :src="recipeImage"
        :alt="recipe.title"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/30 to-transparent"></div>
    </div>

    <!-- Premium Crown (Transparent) -->
    <div
      v-if="recipe.is_premium"
      class="absolute top-3 left-3 bg-black/30 backdrop-blur-sm rounded-full p-2 shadow-lg z-20 border border-white/20"
    >
      <Crown class="w-4 h-4 text-yellow-400" />
    </div>

    <!-- Bookmark Button (Orange when bookmarked) -->
    <button
      class="absolute top-3 right-3 rounded-full p-2 shadow-lg z-20 transition-all duration-200 border backdrop-blur-sm"
      :class="isBookmarked
        ? 'bg-orange-500 border-orange-600 hover:bg-orange-600' 
        : 'bg-black/30 border-white/20 hover:bg-black/40'"
      @click.stop="toggleBookmark"
      :disabled="bookmarkLoading"
    >
      <Bookmark
        :fill="isBookmarked ? 'orange' : 'none'"
        class="w-4 h-4 text-white stroke-current transition-colors duration-200"
      />
    </button>

    <!-- Content Bottom - Default State -->
    <div class="absolute bottom-0 left-0 right-0 p-4 text-white z-10 transition-opacity duration-300"
         :class="isHovered ? 'opacity-0' : 'opacity-100'">
      <h3 class="font-bold text-lg mb-2 line-clamp-2 leading-tight">{{ recipe.title }}</h3>

      <!-- Creator Info -->
      <div class="flex items-center mb-3 text-sm">
        <ChefHat class="w-4 h-4 mr-2 text-orange-300" />
        <span class="font-medium">{{ recipe.user?.username || 'Chef' }}</span>
      </div>

      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <!-- Like Button -->
          <button
            class="flex items-center hover:scale-110 transition-transform duration-200"
            @click.stop="handleLike"
            :disabled="likeLoading"
          >
            <Heart
              :fill="isLiked ? 'red' : 'none'"
              class="w-5 h-5 mr-1 stroke-current transition-all duration-200"
              :class="isLiked ? 'text-red-500' : 'text-white/80'"
            />
            <span class="text-xs font-medium">{{ likeCount }}</span>
          </button>

          <!-- Rating -->
          <div class="flex items-center">
            <Star class="w-4 h-4 text-yellow-400 mr-1" />
            <span class="text-xs font-medium">{{ averageRating.toFixed(1) }}</span>
          </div>
        </div>

        <!-- Time -->
        <span
          class="text-xs font-medium bg-black/40 px-2 py-1 rounded backdrop-blur-sm border border-white/20 flex items-center"
        >
          <Timer class="w-3 h-3 mr-1" />
          {{ recipe.prep_time }}m
        </span>
      </div>
    </div>

    <!-- Hover Overlay with Expanded Details -->
    <div
      class="absolute inset-0 bg-gradient-to-br from-black/95 to-black/90 backdrop-blur-sm transition-all duration-500 flex flex-col justify-between p-6 text-white z-20"
      :class="isHovered ? 'opacity-100' : 'opacity-0 pointer-events-none'"
    >
      <!-- Top Section: Image and Basic Info -->
      <div class="flex space-x-4 mb-4">
        <!-- Recipe Image Thumbnail -->
        <div class="flex-shrink-0 w-24 h-24 rounded-lg overflow-hidden">
          <img
            :src="recipeImage"
            :alt="recipe.title"
            class="w-full h-full object-cover"
          />
        </div>
        
        <!-- Basic Info -->
        <div class="flex-1 min-w-0">
          <h3 class="font-bold text-lg mb-1 line-clamp-2">{{ recipe.title }}</h3>
          <div class="flex items-center text-sm text-gray-300 mb-2">
            <ChefHat class="w-4 h-4 mr-2" />
            <span>By {{ recipe.user?.username || 'Chef' }}</span>
          </div>
          <div class="flex items-center space-x-4 text-xs">
            <div class="flex items-center">
              <Star class="w-4 h-4 text-yellow-400 mr-1" />
              <span>{{ averageRating.toFixed(1) }} ({{ ratingsCount }})</span>
            </div>
            <div class="flex items-center">
              <Heart 
                :fill="isLiked ? 'red' : 'none'" 
                class="w-4 h-4 mr-1"
                :class="isLiked ? 'text-red-500' : 'text-gray-300'"
              />
              <span>{{ likeCount }} likes</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Middle Section: Description -->
      <div class="mb-4">
        <p class="text-sm text-gray-200 leading-relaxed line-clamp-3">
          {{ truncatedDescription }}
        </p>
      </div>

      <!-- Bottom Section: Actions and Metadata -->
      <div class="space-y-3">
        <!-- Metadata -->
        <div class="flex items-center justify-between text-xs">
          <div class="flex items-center space-x-3">
            <span class="flex items-center bg-white/10 px-2 py-1 rounded">
              <Timer class="w-3 h-3 mr-1" />
              {{ recipe.prep_time }}m
            </span>
            <span v-if="recipe.category?.name" class="bg-white/10 px-2 py-1 rounded">
              {{ recipe.category.name }}
            </span>
          </div>
          
          <span
            v-if="recipe.is_premium"
            class="bg-yellow-500/20 text-yellow-300 px-2 py-1 rounded-full text-xs font-semibold"
          >
            ${{ recipe.price }}
          </span>
          <span
            v-else
            class="bg-green-500/20 text-green-300 px-2 py-1 rounded-full text-xs font-semibold"
          >
            Free
          </span>
        </div>

        <!-- Action Buttons -->
        <div class="flex space-x-2">
          <button
            class="flex-1 bg-orange-500 text-white px-4 py-2 rounded-lg font-semibold hover:bg-orange-600 transition-colors duration-200 text-center flex items-center justify-center"
            @click.stop="navigateToRecipeDetail"
          >
            <Eye class="w-4 h-4 mr-2" />
            View Recipe
          </button>
          
          <!-- Like Button in Expanded View -->
          <button
            class="flex items-center justify-center w-10 h-10 bg-white/10 hover:bg-white/20 rounded-lg transition-colors duration-200"
            @click.stop="handleLike"
            :disabled="likeLoading"
            :title="isLiked ? 'Unlike recipe' : 'Like recipe'"
          >
            <Heart
              :fill="isLiked ? 'red' : 'none'"
              class="w-5 h-5"
              :class="isLiked ? 'text-red-500 scale-110' : 'text-white'"
            />
          </button>
          
          <!-- Bookmark Button in Expanded View -->
          <button
            class="flex items-center justify-center w-10 h-10 bg-white/10 hover:bg-white/20 rounded-lg transition-colors duration-200"
            @click.stop="toggleBookmark"
            :disabled="bookmarkLoading"
            :title="isBookmarked ? 'Remove bookmark' : 'Bookmark recipe'"
          >
            <Bookmark
              :fill="isBookmarked ? 'orange' : 'none'"
              class="w-5 h-5"
              :class="isBookmarked ? 'text-orange-400 scale-110' : 'text-white'"
            />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from "vue"
import { navigateTo } from "#app"
import { Heart, Bookmark, Crown, ChefHat, Star, Timer, Eye } from "lucide-vue-next"

const props = defineProps({
  recipe: {
    type: Object,
    required: true,
  },
  showBookmark: {
    type: Boolean,
    default: true
  },
  initialBookmarked: {
    type: Boolean,
    default: false
  }
})

const { isAuthenticated } = useAuth()
const { toggleLike, getLikeStatus, getLikesCount } = useLikes()
const { toggleBookmark: toggleBookmarkApi, getBookmarkStatus } = useBookmarks()
const { getRatingsData } = useRecipeStats()

// Reactive states
const isHovered = ref(false)
const isLiked = ref(false)
const isBookmarked = ref(props.initialBookmarked)
const likeCount = ref(0)
const likeLoading = ref(false)
const bookmarkLoading = ref(false)
const averageRating = ref(0)
const ratingsCount = ref(0)
const statsLoading = ref(true)

// Compute recipe image
const recipeImage = computed(() => {
  if (props.recipe.recipe_images?.length > 0) {
    return props.recipe.recipe_images[0].image_url
  }
  return "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&h=300&q=80"
})

// Truncate description to max 20 words
const truncatedDescription = computed(() => {
  const description = props.recipe.description || 'A delicious recipe waiting for you to explore.'
  const words = description.split(' ')
  if (words.length > 20) {
    return words.slice(0, 20).join(' ') + '...'
  }
  return description
})

// Navigation
const navigateToRecipeDetail = () => {
  navigateTo(`/recipes/${props.recipe.id}`)
}

// Load recipe statistics
const loadRecipeStats = async () => {
  try {
    statsLoading.value = true
    const [likesCount, ratingsData] = await Promise.all([
      getLikesCount(props.recipe.id),
      getRatingsData(props.recipe.id),
    ])

    likeCount.value = likesCount
    averageRating.value = ratingsData.averageRating
    ratingsCount.value = ratingsData.ratingsCount
  } catch (error) {
    console.error("Error loading recipe stats:", error)
  } finally {
    statsLoading.value = false
  }
}

// Like functionality
const handleLike = async (e) => {
  e.stopPropagation()

  if (!isAuthenticated.value) {
    navigateTo("/login")
    return
  }

  likeLoading.value = true
  try {
    const result = await toggleLike(props.recipe.id)
    isLiked.value = result.liked
    likeCount.value = result.likeCount
  } catch (error) {
    console.error("Failed to toggle like:", error)
  } finally {
    likeLoading.value = false
  }
}

// Bookmark functionality
const toggleBookmark = async () => {
  if (!props.showBookmark) return
  
  if (!isAuthenticated.value) {
    navigateTo("/login")
    return
  }

  bookmarkLoading.value = true
  try {
    const result = await toggleBookmarkApi(props.recipe.id)
    isBookmarked.value = result.bookmarked
    
    // Emit event for parent component (bookmarks page)
    if (!result.bookmarked) {
      emit('bookmark-removed', props.recipe.id)
    }
  } catch (error) {
    console.error("Failed to toggle bookmark:", error)
  } finally {
    bookmarkLoading.value = false
  }
}

// Load initial data
onMounted(async () => {
  await loadRecipeStats()

  if (isAuthenticated.value) {
    const [likeStatus, bookmarkStatus] = await Promise.all([
      getLikeStatus(props.recipe.id),
      props.showBookmark ? getBookmarkStatus(props.recipe.id) : Promise.resolve({ bookmarked: false })
    ])
    
    isLiked.value = likeStatus.liked
    likeCount.value = likeStatus.likeCount
    isBookmarked.value = bookmarkStatus.bookmarked
  }
})

// Watch authentication changes
watch(isAuthenticated, async (newVal) => {
  if (newVal) {
    const [likeStatus, bookmarkStatus] = await Promise.all([
      getLikeStatus(props.recipe.id),
      props.showBookmark ? getBookmarkStatus(props.recipe.id) : Promise.resolve({ bookmarked: false })
    ])
    
    isLiked.value = likeStatus.liked
    likeCount.value = likeStatus.likeCount
    isBookmarked.value = bookmarkStatus.bookmarked
  } else {
    isLiked.value = false
    isBookmarked.value = false
  }
})

// Define emit for bookmark removal
const emit = defineEmits(['bookmark-removed'])
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

button {
  transition: all 0.2s ease-in-out;
}

.group:hover {
  transform: translateY(-4px);
}
</style>