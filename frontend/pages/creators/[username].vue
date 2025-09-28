<!-- pages/creators/[username].vue -->
<template>
  <div class="min-h-screen bg-gray-50">
    <Navigation />
    
    <!-- Creator Header -->
    <div class="bg-orange-500 border-b">
      <div class="max-w-7xl mx-auto px-4 py-12">
        <div class="flex items-center space-x-6">
          <!-- Creator Avatar -->
          <div class="w-24 h-24 rounded-full bg-gradient-to-br from-orange-400 to-red-500 flex items-center justify-center">
            <span class="text-white text-3xl font-bold">{{ creator?.username?.charAt(0)?.toUpperCase() }}</span>
          </div>
          
          <!-- Creator Info -->
          <div>
            <h1 class="text-4xl font-serif font-bold text-gray-900 mb-2"> {{ creator?.username }}</h1>
            <p class="text-gray-600 text-lg">{{ creator?.recipe_count || 0 }} recipes</p>
            <p class="text-gray-500 text-sm">Joined {{ formatDate(creator?.created_at) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Recipes Grid -->
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Loading State -->
      <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <RecipeCardSkeleton v-for="n in 6" :key="n" />
      </div>

      <!-- Recipes -->
      <div v-else-if="creator?.recipes?.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <RecipeCard 
          v-for="recipe in creator.recipes"
          :key="recipe.id"
          :recipe="recipe"
          :show-bookmark="true"
        />
      </div>

      <!-- Empty State -->
      <div v-else-if="!loading && creator" class="text-center py-16">
        <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <UIcon name="i-heroicons-clipboard-document-list" class="h-12 w-12 text-gray-400" />
        </div>
        <h3 class="text-2xl font-semibold text-gray-900 mb-2">No recipes yet</h3>
        <p class="text-gray-600">This creator hasn't published any recipes yet.</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-16">
        <div class="bg-red-50 border border-red-200 rounded-lg p-6 max-w-md mx-auto">
          <UIcon name="i-heroicons-exclamation-triangle" class="w-12 h-12 text-red-600 mx-auto mb-4" />
          <h3 class="text-xl font-semibold text-gray-900 mb-2">Creator not found</h3>
          <p class="text-gray-600 mb-4">{{ error }}</p>
          <button 
            @click="$router.push('/')"
            class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors"
          >
            Back to Home
          </button>
        </div>
      </div>
    </div>

    <Footer />
  </div>
</template>

<script setup>
const route = useRoute()
const { username } = route.params

// Composables
const { getCreatorByUsername } = useCreators()

// Reactive data
const loading = ref(true)
const error = ref('')
const creator = ref(null)

// Load creator data
const loadCreator = async () => {
  try {
    loading.value = true
    error.value = ''
    
    const creatorData = await getCreatorByUsername(username)
    
    if (!creatorData) {
      error.value = 'Creator not found'
      return
    }
    
    creator.value = creatorData
    console.log('Creator data loaded:', creator.value)
    console.log('Recipes count:', creator.value.recipes?.length)
  } catch (err) {
    console.error('Error loading creator:', err)
    error.value = err.message || 'Failed to load creator. Please try again.'
  } finally {
    loading.value = false
  }
}

// Format date
const formatDate = (dateString) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long'
  })
}

// Set page title
useHead({
  title: creator.value ? ` ${creator.value.username} - Seasoned` : 'Creator - Seasoned'
})

// Load data on mount
onMounted(() => {
  loadCreator()
})

// Watch for route changes
watch(() => route.params.username, (newUsername) => {
  if (newUsername) {
    loadCreator()
  }
})
</script>