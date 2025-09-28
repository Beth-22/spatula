<template>
  <div class="min-h-screen bg-white">
    <!-- Loading State -->
    <div v-if="loading" class="min-h-screen flex items-center justify-center">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto mb-4"></div>
        <p class="text-gray-600">Loading recipe...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="min-h-screen flex items-center justify-center">
      <div class="text-center">
        <div class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
          <span class="text-2xl">❌</span>
        </div>
        <h2 class="text-xl font-semibold text-gray-900 mb-2">Recipe Not Found</h2>
        <p class="text-gray-600 mb-4">{{ error }}</p>
        <button @click="goBack" class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600">
          Go Back
        </button>
      </div>
    </div>

    <!-- Recipe Content - ChefSteps Style -->
    <div v-else-if="recipe" class="min-h-screen bg-white">
      <!-- Hero Section with Carousel -->
      <div class="relative h-96 bg-gray-900">
        <!-- Carousel -->
        <div class="absolute inset-0 overflow-hidden">
          <div class="relative w-full h-full">
            <!-- Show all images with proper conditional rendering -->
            <template v-for="(image, index) in carouselImages" :key="image.id">
              <img 
                :src="image.image_url"
                :alt="recipe.title"
                class="absolute top-0 left-0 w-full h-full object-cover opacity-90 transition-opacity duration-500"
                :class="{ 'opacity-100 z-10': currentImageIndex === index, 'opacity-0 z-0': currentImageIndex !== index }"
              />
            </template>
            
            <!-- Fallback if no images -->
            <div v-if="carouselImages.length === 0" class="w-full h-full bg-gray-700 flex items-center justify-center">
              <span class="text-white text-2xl">No Image Available</span>
            </div>
          </div>
          <div class="absolute inset-0 bg-black/40"></div>
        </div>

        <!-- Carousel Indicators -->
        <div v-if="carouselImages.length > 1" class="absolute bottom-4 left-1/2 transform -translate-x-1/2 z-20 flex space-x-2">
          <button
            v-for="(image, index) in carouselImages"
            :key="image.id"
            @click="currentImageIndex = index"
            class="w-3 h-3 rounded-full transition-all duration-300"
            :class="currentImageIndex === index ? 'bg-white' : 'bg-white/50'"
          />
        </div>

        <!-- Carousel Navigation -->
        <div v-if="carouselImages.length > 1" class="absolute inset-0 z-10 flex items-center justify-between px-4">
          <button 
            @click="prevImage"
            class="p-2 rounded-full bg-black/30 text-white hover:bg-black/50 transition-all"
          >
            <Icon icon="mdi:chevron-left" class="h-6 w-6" />
          </button>
          <button 
            @click="nextImage"
            class="p-2 rounded-full bg-black/30 text-white hover:bg-black/50 transition-all"
          >
            <Icon icon="mdi:chevron-right" class="h-6 w-6" />
          </button>
        </div>

        <!-- Navigation -->
        <nav class="relative z-10">
          <div class="max-w-6xl mx-auto px-6 py-6">
            <div class="flex justify-between items-center">
              <button @click="goBack" class="flex items-center text-white hover:text-gray-200 transition-colors">
                <Icon icon="mdi:arrow-left" class="h-5 w-5 mr-2" />
                <span class="text-lg font-medium">Back to Recipes</span>
              </button>
              <div class="flex items-center space-x-4">
                <!-- Like Button in Header -->
                <button 
                  @click="handleLike"
                  :disabled="likeLoading"
                  class="flex items-center space-x-2 text-white hover:text-gray-200 transition-colors p-2"
                  :class="likeLoading ? 'opacity-50 cursor-not-allowed' : ''"
                >
                  <Heart
                    :fill="isLiked ? 'red' : 'none'"
                    class="w-5 h-5 transition-all duration-200"
                    :class="isLiked ? 'text-red-500 scale-110' : 'text-white'"
                  />
                  <span class="font-medium">{{ likeCount }}</span>
                </button>
                
                <button class="text-white hover:text-gray-200 transition-colors p-2">
                  <Icon icon="mdi:bookmark-outline" class="h-5 w-5" />
                </button>
                <button class="text-white hover:text-gray-200 transition-colors p-2">
                  <Icon icon="mdi:share-variant" class="h-5 w-5" />
                </button>
              </div>
            </div>
          </div>
        </nav>

        <!-- Recipe Title Centered -->
        <div class="relative z-10 max-w-4xl mx-auto px-6 h-full flex flex-col items-center justify-center text-center">
          <h1 class="text-5xl lg:text-6xl font-light text-white mb-4 leading-tight">{{ recipe.title }}</h1>
          <div class="flex items-center space-x-4 text-gray-200">
            <Icon icon="mdi:chef-hat" class="h-6 w-6" />
            <span class="text-xl">{{ recipe.user?.username || 'Unknown Chef' }}</span>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="max-w-6xl mx-auto px-6 py-12">
        <!-- Description & Recipe Details Side by Side -->
        <div class="grid grid-cols-1 lg:grid-cols-4 gap-12 mb-16">
          <!-- Description - Left Column -->
          <div class="lg:col-span-3">
            <div class="prose prose-lg max-w-none">
              <p class="text-gray-700 text-lg leading-relaxed">{{ recipe.description }}</p>
            </div>
          </div>

          <!-- Recipe Details Card - Right Column -->
          <div class="lg:col-span-1">
            <div class="bg-gray-50 rounded-lg p-6 border border-gray-200 sticky top-8">
              <h3 class="text-xl font-medium text-gray-900 mb-6">Recipe Details</h3>
              <div class="space-y-4">
                <!-- Prep Time -->
                <div class="flex items-center justify-between" v-if="recipe.prep_time">
                  <span class="text-gray-600">Prep time</span>
                  <span class="font-medium text-gray-900">{{ formatTime(recipe.prep_time) }}</span>
                </div>
                
                <!-- Rating -->
                <div class="flex items-center justify-between">
                  <span class="text-gray-600">Rating</span>
                  <div class="flex items-center space-x-1">
                    <span class="font-medium text-gray-900">{{ averageRating }}</span>
                    <Icon icon="mdi:star" class="h-4 w-4 text-yellow-400" />
                    <span class="text-gray-500 text-sm">({{ ratingCount }})</span>
                  </div>
                </div>

                <!-- Like Count -->
                <div class="flex items-center justify-between">
                  <span class="text-gray-600">Likes</span>
                  <div class="flex items-center space-x-1">
                    <span class="font-medium text-gray-900">{{ likeCount }}</span>
                    <Icon icon="mdi:heart" class="h-4 w-4 text-red-400" />
                  </div>
                </div>
                
                <!-- Chef/Owner -->
                <div class="flex items-center justify-between pt-3 border-t border-gray-200">
                  <span class="text-gray-600">Chef</span>
                  <span class="font-medium text-gray-900">{{ recipe.user?.username || 'Unknown' }}</span>
                </div>
                
                <!-- Category -->
                <div class="flex items-center justify-between" v-if="recipe.category">
                  <span class="text-gray-600">Category</span>
                  <span class="font-medium text-gray-900 capitalize">{{ recipe.category.name }}</span>
                </div>

                <!-- Like Button in Details Card -->
                <div class="pt-3 border-t border-gray-200">
                  <button 
                    @click="handleLike"
                    :disabled="likeLoading"
                    class="w-full flex items-center justify-center space-x-2 bg-white border border-gray-300 rounded-lg px-4 py-2 hover:bg-gray-50 transition-colors"
                    :class="likeLoading ? 'opacity-50 cursor-not-allowed' : ''"
                  >
                    <Heart
                      :fill="isLiked ? 'red' : 'none'"
                      class="w-5 h-5 transition-all duration-200"
                      :class="isLiked ? 'text-red-500 scale-110' : 'text-gray-600'"
                    />
                    <span class="font-medium text-gray-700">
                      {{ isLiked ? 'Liked' : 'Like this recipe' }}
                    </span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Ingredients Section -->
        <div class="mb-16">
          <h2 class="text-3xl font-light text-gray-900 mb-8 text-center">Ingredients</h2>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8 max-w-4xl mx-auto">
            <div 
              v-for="ingredient in visibleIngredients" 
              :key="ingredient.id" 
              class="flex items-center justify-between py-3 border-b border-gray-100"
            >
              <span class="text-gray-700 text-lg">{{ ingredient.name }}</span>
              <span class="text-gray-900 font-medium text-lg">{{ ingredient.quantity }} {{ ingredient.unit || '' }}</span>
            </div>
          </div>

          <!-- Premium Notice -->
          <div 
            v-if="recipe.is_premium && !hasPurchased && visibleIngredients.length < recipe.ingredients.length"
            class="mt-8 p-6 bg-orange-50 rounded-lg border border-orange-200 max-w-4xl mx-auto text-center"
          >
            <p class="text-orange-800 text-lg font-medium mb-2">Unlock all {{ recipe.ingredients.length }} ingredients by purchasing this recipe</p>
            <button class="bg-orange-500 text-white px-6 py-3 rounded-lg text-lg font-medium hover:bg-orange-600 transition-colors">
              Purchase for ${{ recipe.price }}
            </button>
          </div>
        </div>

        <!-- Steps Section -->
        <div class="mb-16">
          <h2 class="text-3xl font-light text-gray-900 mb-12 text-center">Steps</h2>
          
          <div class="space-y-12 max-w-4xl mx-auto">
            <div 
              v-for="step in recipe.steps" 
              :key="step.id" 
              class="flex space-x-8"
            >
              <!-- Step Number -->
              <div class="flex-shrink-0">
                <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center">
                  <span class="text-gray-600 font-bold text-2xl">{{ step.step_number }}</span>
                </div>
              </div>
              
              <!-- Step Content -->
              <div class="flex-1">
                <p class="text-gray-700 text-xl leading-relaxed">{{ step.instruction }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Rate This Recipe Section -->
        <div class="border-t border-gray-200 pt-16 mb-16">
          <div class="text-center">
            <h2 class="text-3xl font-light text-gray-900 mb-4">Rate this Recipe</h2>
            <p class="text-gray-600 text-xl mb-8">How did it turn out?</p>

            <!-- Star Rating -->
            <div class="flex justify-center items-center space-x-2 mb-6">
              <button
                v-for="star in 5"
                :key="star"
                @click="toggleRating(star)"
                class="transition-transform duration-200 hover:scale-110 text-4xl"
                :class="getStarClass(star)"
                @mouseenter="hoverRating = star"
                @mouseleave="hoverRating = 0"
              >
                <span v-if="shouldShowFilledStar(star)">★</span>
                <span v-else>☆</span>
              </button>
            </div>

            <!-- Rating Text -->
            <div class="text-gray-600">
              <span v-if="currentRating > 0" class="text-xl font-medium">
                You rated this {{ currentRating }} star{{ currentRating > 1 ? 's' : '' }}
                <button @click="removeRating" class="text-orange-500 hover:text-orange-700 ml-2 text-sm underline">
                  (remove rating)
                </button>
              </span>
              <span v-else class="text-xl">Click a star to rate</span>
            </div>

            <!-- Current Average Rating -->
            <div v-if="ratingCount > 0" class="mt-6 text-gray-600">
              <p class="text-lg">Average rating: {{ averageRating }} stars ({{ ratingCount }} ratings)</p>
            </div>
          </div>
        </div>

        <!-- Comments Section -->
        <div class="border-t border-gray-200 pt-16">
          <CommentsSection :recipe-id="recipe.id" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import CommentsSection from '~/components/Comments/CommentSection.vue'
import { Heart } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const { id } = route.params

// Composables
const { getRecipeById } = useRecipeApi()
const { setRating, deleteRating, getRatingStatus } = useRatings()
const { toggleLike, getLikeStatus, getLikesCount } = useLikes()
const { isAuthenticated } = useAuth()

// Reactive data
const loading = ref(true)
const error = ref('')
const recipe = ref(null)
const currentImageIndex = ref(0)
const hasPurchased = ref(false)
const currentRating = ref(0)
const hoverRating = ref(0)
const averageRating = ref(0)
const ratingCount = ref(0)
const carouselInterval = ref(null)

// Like functionality
const isLiked = ref(false)
const likeCount = ref(0)
const likeLoading = ref(false)

// Computed properties
const carouselImages = computed(() => {
  if (!recipe.value || !recipe.value.recipe_images) return []
  return recipe.value.recipe_images || []
})

const visibleIngredients = computed(() => {
  if (!recipe.value || !recipe.value.ingredients) return []
  if (recipe.value.is_premium && !hasPurchased.value) {
    return recipe.value.ingredients.slice(0, 5)
  }
  return recipe.value.ingredients
})

// Lifecycle
onMounted(async () => {
  await loadRecipeData()
})

onUnmounted(() => {
  if (carouselInterval.value) {
    clearInterval(carouselInterval.value)
  }
})

// Methods
const loadRecipeData = async () => {
  try {
    loading.value = true
    error.value = ''
    
    const recipeData = await getRecipeById(id)
    if (!recipeData) {
      error.value = 'Recipe not found'
      return
    }
    recipe.value = recipeData
    
    console.log('Recipe data loaded:', recipe.value)
    
    // Check if this is a premium recipe that needs redirect
    if (recipe.value.is_premium) {
      await checkPurchaseStatus()
      
      // Redirect to premium page if not purchased
      if (!hasPurchased.value) {
        await navigateTo(`/recipes/premium/${id}`)
        return
      }
    }
    
    // Start carousel if there are multiple images
    if (carouselImages.value.length > 1) {
      startCarousel()
    }
    
    // Load rating information
    await loadRatingData()
    
    // Load like information
    await loadLikeData()
    
  } catch (err) {
    console.error('Error loading recipe:', err)
    error.value = 'Failed to load recipe. Please try again.'
  } finally {
    loading.value = false
  }
}

const checkPurchaseStatus = async () => {
  try {
    const response = await $fetch(`/actions/recipe/access?recipe_id=${id}`)
    hasPurchased.value = response.has_access
    console.log('Purchase status:', hasPurchased.value)
  } catch (err) {
    console.error('Error checking purchase status:', err)
    hasPurchased.value = false
  }
}

const loadRatingData = async () => {
  try {
    const ratingStatus = await getRatingStatus(recipe.value.id)
    currentRating.value = ratingStatus.userRating
    averageRating.value = ratingStatus.averageRating
    ratingCount.value = ratingStatus.ratingCount
    
    console.log('Rating data loaded:', {
      userRating: currentRating.value,
      averageRating: averageRating.value,
      ratingCount: ratingCount.value
    })
  } catch (err) {
    console.error('Error loading rating data:', err)
    // Continue without rating data
  }
}

const loadLikeData = async () => {
  try {
    // Get like count (public data)
    const count = await getLikesCount(recipe.value.id)
    likeCount.value = count
    
    // Get like status if authenticated
    if (isAuthenticated.value) {
      const likeStatus = await getLikeStatus(recipe.value.id)
      isLiked.value = likeStatus.liked
      likeCount.value = likeStatus.likeCount
    }
    
    console.log('Like data loaded:', {
      isLiked: isLiked.value,
      likeCount: likeCount.value
    })
  } catch (err) {
    console.error('Error loading like data:', err)
    // Continue without like data
  }
}

const handleLike = async () => {
  if (!isAuthenticated.value) {
    navigateTo('/login')
    return
  }

  likeLoading.value = true
  try {
    const result = await toggleLike(recipe.value.id)
    isLiked.value = result.liked
    likeCount.value = result.likeCount
  } catch (error) {
    console.error('Failed to toggle like:', error)
  } finally {
    likeLoading.value = false
  }
}

const startCarousel = () => {
  carouselInterval.value = setInterval(() => {
    nextImage()
  }, 5000)
}

const nextImage = () => {
  if (carouselImages.value.length > 0) {
    currentImageIndex.value = (currentImageIndex.value + 1) % carouselImages.value.length
  }
}

const prevImage = () => {
  if (carouselImages.value.length > 0) {
    currentImageIndex.value = currentImageIndex.value === 0 ? carouselImages.value.length - 1 : currentImageIndex.value - 1
  }
}

const rateRecipe = async (rating) => {
  console.log('Rating recipe:', rating)
  
  if (!isAuthenticated.value) {
    navigateTo('/login')
    return
  }

  try {
    console.log('Calling setRating...')
    const result = await setRating(recipe.value.id, rating)
    console.log('Rating result:', result)
    
    // Update local state
    currentRating.value = rating
    averageRating.value = result.averageRating
    ratingCount.value = result.ratingCount
    
  } catch (error) {
    console.error('Failed to set rating:', error)
    alert('Failed to submit rating. Please try again.')
  }
}

const removeRating = async () => {
  if (!isAuthenticated.value) {
    navigateTo('/login')
    return
  }

  try {
    console.log('Removing rating...')
    const result = await deleteRating(recipe.value.id)
    console.log('Rating removed:', result)
    
    // Update local state
    currentRating.value = 0
    averageRating.value = result.averageRating
    ratingCount.value = result.ratingCount
    
  } catch (error) {
    console.error('Failed to remove rating:', error)
    alert('Failed to remove rating. Please try again.')
  }
}

const toggleRating = async (star) => {
  if (!isAuthenticated.value) {
    navigateTo('/login')
    return
  }

  if (currentRating.value === star) {
    await removeRating()
    return
  }

  await rateRecipe(star)
}

const getStarClass = (star) => {
  if (hoverRating >= star || currentRating >= star) {
    return 'text-yellow-500'
  }
  return 'text-gray-300'
}

const shouldShowFilledStar = (star) => {
  return hoverRating >= star || currentRating >= star
}

const goBack = () => {
  router.back()
}

const formatTime = (minutes) => {
  if (!minutes) return '0m'
  const hours = Math.floor(minutes / 60)
  const mins = minutes % 60
  return hours > 0 ? `${hours}h ${mins}m` : `${mins}m`
}

useHead({
  title: recipe.value ? `${recipe.value.title} - Spatula` : 'Recipe - Spatula'
})
</script>

<style scoped>
.transition-opacity {
  transition: opacity 0.5s ease-in-out;
}

.text-4xl {
  font-size: 2.5rem;
  line-height: 1;
}

.sticky {
  position: sticky;
  top: 2rem;
}
</style>