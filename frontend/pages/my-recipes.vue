<!-- pages/my-recipes.vue -->
<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 to-amber-50">
    <Navigation />
    
    <!-- Header -->
    <div class="bg-gradient-to-r from-orange-500 to-red-500 text-white">
      <div class="max-w-7xl mx-auto px-4 py-16">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-4xl font-serif font-bold mb-2">My Recipes</h1>
            <p class="text-xl opacity-90">Manage and edit your culinary creations</p>
          </div>
          
          <!-- Create New Recipe Button -->
          <button
            @click="navigateToUpload"
            class="flex items-center gap-2 bg-white text-orange-600 px-6 py-3 rounded-xl hover:bg-orange-50 transition-all font-semibold shadow-lg"
          >
            <Icon v-if="isIconReady" icon="mdi:plus" class="h-5 w-5" />
            <span v-else>+</span>
            Create New Recipe
          </button>
        </div>
        
        <!-- Stats -->
        <div class="grid grid-cols-3 gap-8 max-w-2xl mt-8">
          <div class="text-center">
            <div class="text-2xl font-bold">{{ stats.total }}</div>
            <div class="text-orange-100 text-sm">Total Recipes</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold">{{ stats.premium }}</div>
            <div class="text-orange-100 text-sm">Premium</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold">{{ stats.likes }}</div>
            <div class="text-orange-100 text-sm">Total Likes</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 py-12">
      <!-- Error Message -->
      <div v-if="errorMessage" class="mb-6 bg-red-50 border border-red-200 rounded-xl p-4">
        <div class="flex items-center">
          <Icon v-if="isIconReady" icon="mdi:alert-circle-outline" class="h-5 w-5 text-red-500 mr-2" />
          <span v-else class="text-red-500 mr-2">⚠️</span>
          <span class="text-red-700">{{ errorMessage }}</span>
        </div>
      </div>

      <!-- Not Authenticated Message -->
      <div v-if="!isAuthenticated && !loading" class="text-center py-16">
        <div class="w-24 h-24 bg-white rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
          <Icon v-if="isIconReady" icon="mdi:lock-outline" class="h-12 w-12 text-gray-400" />
          <span v-else class="text-2xl">🔒</span>
        </div>
        <h3 class="text-2xl font-semibold text-gray-900 mb-2">Authentication Required</h3>
        <p class="text-gray-600 mb-6">Please log in to view your recipes</p>
        <button 
          @click="navigateToLogin"
          class="bg-gradient-to-r from-orange-500 to-red-500 text-white px-8 py-3 rounded-xl hover:shadow-lg transition-all font-semibold"
        >
          Go to Login
        </button>
      </div>

      <!-- Content for authenticated users -->
      <div v-else>
        <!-- Filter & Sort Bar -->
        <div class="flex flex-wrap items-center justify-between gap-4 mb-8">
          <div class="text-lg font-semibold text-gray-900">
            My Recipes <span class="text-gray-600">({{ filteredRecipes.length }})</span>
          </div>
          
          <div class="flex gap-2">
            <!-- Sort Button -->
            <div class="relative">
              <button
                @click="showSortDropdown = !showSortDropdown"
                class="flex items-center gap-2 px-4 py-3 bg-white border border-gray-200 rounded-xl hover:border-gray-300 transition-colors shadow-sm"
              >
                <Icon v-if="isIconReady" icon="mdi:sort" class="h-5 w-5 text-gray-600" />
                <span v-else>↕️</span>
                <span class="text-gray-700">Sort</span>
              </button>
              
              <div v-if="showSortDropdown" class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-xl shadow-lg z-20">
                <button
                  v-for="option in sortOptions"
                  :key="option.value"
                  @click="selectSort(option.value)"
                  class="w-full px-4 py-3 text-left hover:bg-gray-50 transition-colors first:rounded-t-xl last:rounded-b-xl"
                  :class="selectedSort === option.value ? 'bg-orange-50 text-orange-700' : 'text-gray-700'"
                >
                  {{ option.label }}
                </button>
              </div>
            </div>
            
            <!-- Filter Toggle Button -->
            <button
              @click="showFilters = true"
              class="flex items-center gap-2 px-4 py-3 bg-orange-500 text-white rounded-xl hover:bg-orange-600 transition-colors shadow-lg"
            >
              <Icon v-if="isIconReady" icon="mdi:filter" class="h-5 w-5" />
              <span v-else>🔧</span>
              <span>Filters</span>
            </button>
          </div>
        </div>

        <!-- Recipes Grid -->
        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
          <div v-for="n in 8" :key="n" class="bg-white rounded-xl shadow-lg overflow-hidden h-96 w-64 mx-auto animate-pulse">
            <div class="h-48 bg-gray-300"></div>
            <div class="p-4">
              <div class="h-4 bg-gray-300 rounded mb-2"></div>
              <div class="h-3 bg-gray-300 rounded w-3/4 mb-4"></div>
              <div class="flex justify-between items-center mb-3">
                <div class="flex space-x-2">
                  <div class="h-3 w-8 bg-gray-300 rounded"></div>
                  <div class="h-3 w-8 bg-gray-300 rounded"></div>
                </div>
                <div class="h-3 w-12 bg-gray-300 rounded"></div>
              </div>
              <div class="flex justify-between">
                <div class="h-3 w-16 bg-gray-300 rounded"></div>
                <div class="h-3 w-10 bg-gray-300 rounded"></div>
              </div>
            </div>
          </div>
        </div>
        
        <div v-else-if="filteredRecipes.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
          <MyRecipeCard 
            v-for="recipe in sortedRecipes"
            :key="recipe.id"
            :recipe="recipe"
            @edit="handleEdit"
            @delete="handleDelete"
          />
        </div>

        <!-- Empty States -->
        <div v-else-if="!loading" class="text-center py-16">
          <div v-if="userRecipes.length === 0">
            <div class="w-24 h-24 bg-white rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
              <Icon v-if="isIconReady" icon="mdi:chef-hat" class="h-12 w-12 text-gray-400" />
              <span v-else class="text-2xl">👨‍🍳</span>
            </div>
            <h3 class="text-2xl font-semibold text-gray-900 mb-2">No recipes yet</h3>
            <p class="text-gray-600 mb-6">Start sharing your culinary creations with the community</p>
            <button 
              @click="navigateToUpload"
              class="bg-gradient-to-r from-orange-500 to-red-500 text-white px-8 py-3 rounded-xl hover:shadow-lg transition-all font-semibold"
            >
              Create Your First Recipe
            </button>
          </div>
          
          <div v-else>
            <div class="w-24 h-24 bg-white rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
              <Icon v-if="isIconReady" icon="mdi:filter-off" class="h-12 w-12 text-gray-400" />
              <span v-else class="text-2xl">🚫</span>
            </div>
            <h3 class="text-2xl font-semibold text-gray-900 mb-2">No recipes match your filters</h3>
            <p class="text-gray-600">Try adjusting your filters to see more results</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Filter Slide-out Panel -->
    <div v-if="showFilters" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black bg-opacity-50" @click="showFilters = false"></div>
      <div class="absolute right-0 top-0 h-full w-96 bg-white transform transition-transform">
        <div class="p-6 h-full overflow-y-auto">
          <!-- Header -->
          <div class="flex justify-between items-center mb-6">
            <h2 class="text-xl font-semibold text-gray-900">Filter Recipes</h2>
            <button @click="showFilters = false" class="p-2 hover:bg-gray-100 rounded-lg">
              <Icon v-if="isIconReady" icon="mdi:close" class="h-5 w-5" />
              <span v-else>X</span>
            </button>
          </div>

          <!-- All Filters in One Card -->
          <div class="bg-gray-50 rounded-xl p-5 space-y-6">
            <!-- Recipe Type -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Recipe Type</h3>
              <div class="space-y-2">
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="all"
                    v-model="filters.recipeType"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">All Recipes</span>
                </label>
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="premium"
                    v-model="filters.recipeType"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Premium Only</span>
                </label>
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="free"
                    v-model="filters.recipeType"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Free Only</span>
                </label>
              </div>
            </div>

            <!-- Category -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Category</h3>
              <select 
                v-model="filters.category"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              >
                <option value="">All Categories</option>
                <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                  {{ cat.name }}
                </option>
              </select>
            </div>

            <!-- Clear Filters -->
            <button
              @click="clearFilters"
              class="w-full py-3 text-gray-600 hover:text-gray-800 border border-gray-300 rounded-lg hover:border-gray-400 transition-colors font-medium"
            >
              Clear All Filters
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
      <div class="bg-white rounded-2xl p-6 max-w-md mx-4">
        <div class="text-center">
          <Icon v-if="isIconReady" icon="mdi:alert-circle-outline" class="h-16 w-16 text-red-500 mx-auto mb-4" />
          <span v-else class="text-4xl mb-4 block">⚠️</span>
          <h3 class="text-xl font-semibold text-gray-900 mb-2">Delete Recipe</h3>
          <p class="text-gray-600 mb-6">Are you sure you want to delete "{{ recipeToDelete?.title }}"? This action cannot be undone.</p>
          
          <div class="flex gap-3">
            <button
              @click="showDeleteModal = false"
              class="flex-1 py-3 text-gray-600 border border-gray-300 rounded-lg hover:border-gray-400 transition-colors"
            >
              Cancel
            </button>
            <button
              @click="confirmDelete"
              class="flex-1 py-3 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors font-semibold"
              :disabled="deleting"
            >
              <span v-if="deleting" class="flex items-center justify-center">
                <Icon v-if="isIconReady" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
                <span v-else class="mr-2">⏳</span>
                Deleting...
              </span>
              <span v-else>Delete Recipe</span>
            </button>
          </div>
        </div>
      </div>
    </div>

    <Footer />
  </div>
</template>

<script setup>
// Composables
const { getMyRecipes, deleteRecipe } = useMyRecipes()
const { getCategories } = useRecipeApi()
const { user, isAuthenticated } = useAuth()
const { getRecipesStats } = useRecipeStats()

// Data
const loading = ref(false)
const showFilters = ref(false)
const showSortDropdown = ref(false)
const showDeleteModal = ref(false)
const deleting = ref(false)
const userRecipes = ref([])
const categories = ref([])
const recipeToDelete = ref(null)
const errorMessage = ref('')
const isIconReady = ref(false)

// Filters
const filters = reactive({
  recipeType: 'all',
  category: ''
})

// Sort
const selectedSort = ref('newest')
const sortOptions = [
  { label: 'Newest First', value: 'newest' },
  { label: 'Oldest First', value: 'oldest' },
  { label: 'Most Popular', value: 'popular' },
  { label: 'A-Z', value: 'alphabetical' }
]

// Computed
const stats = computed(() => {
  const total = userRecipes.value.length
  const premium = userRecipes.value.filter(r => r.is_premium).length
  const likes = userRecipes.value.reduce((sum, r) => sum + (r.likesCount || 0), 0)
  
  return { total, premium, likes }
})

const filteredRecipes = computed(() => {
  return userRecipes.value.filter(recipe => {
    if (filters.recipeType === 'premium' && !recipe.is_premium) return false
    if (filters.recipeType === 'free' && recipe.is_premium) return false
    if (filters.category && recipe.category?.id !== filters.category) return false
    return true
  })
})

const sortedRecipes = computed(() => {
  const recipes = [...filteredRecipes.value]
  
  switch (selectedSort.value) {
    case 'oldest':
      return recipes.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
    case 'popular':
      return recipes.sort((a, b) => (b.likesCount || 0) - (a.likesCount || 0))
    case 'alphabetical':
      return recipes.sort((a, b) => a.title.localeCompare(b.title))
    case 'newest':
    default:
      return recipes.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  }
})

// Methods
const selectSort = (sortValue) => {
  selectedSort.value = sortValue
  showSortDropdown.value = false
}

const clearFilters = () => {
  filters.recipeType = 'all'
  filters.category = ''
}

const navigateToUpload = () => {
  navigateTo('/recipes/upload')
}

const navigateToLogin = () => {
  navigateTo('/login')
}

const handleEdit = (recipe) => {
  navigateTo(`/recipes/${recipe.id}/edit`)
}

const handleDelete = (recipe) => {
  recipeToDelete.value = recipe
  showDeleteModal.value = true
}

const confirmDelete = async () => {
  if (!recipeToDelete.value) return
  
  deleting.value = true
  try {
    await deleteRecipe(recipeToDelete.value.id)
    userRecipes.value = userRecipes.value.filter(r => r.id !== recipeToDelete.value.id)
    showDeleteModal.value = false
    recipeToDelete.value = null
    console.log('Recipe deleted successfully')
  } catch (error) {
    console.error('Failed to delete recipe:', error)
  } finally {
    deleting.value = false
  }
}

// In pages/my-recipes.vue - update the loadData function
const loadData = async () => {
  if (!isAuthenticated.value) {
    console.log('User not authenticated, skipping data load')
    userRecipes.value = []
    categories.value = []
    return
  }

  loading.value = true
  errorMessage.value = ''
  try {
    console.log('Loading user recipes for user:', user.value)
    const recipesData = await getMyRecipes()
    console.log('Raw recipes data:', recipesData)
    
    // Create new recipe objects with stats instead of modifying existing ones
    if (recipesData && recipesData.length > 0) {
      const recipeIds = recipesData.map(recipe => recipe.id)
      console.log('Recipe IDs for stats:', recipeIds)
      
      const stats = await getRecipesStats(recipeIds)
      console.log('Stats data:', stats)
      
      // Create new recipe objects with stats included
      userRecipes.value = recipesData.map(recipe => {
        const recipeStats = stats[recipe.id] || { likesCount: 0, averageRating: 0, ratingsCount: 0 }
        
        // Create a new object with all recipe properties and stats
        return {
          ...recipe, // Spread all original recipe properties
          likesCount: recipeStats.likesCount,
          averageRating: recipeStats.averageRating,
          ratingsCount: recipeStats.ratingsCount,
          // Ensure recipe_images is properly handled
          recipe_images: recipe.recipe_images || []
        }
      })
      
      console.log('Processed recipes with stats:', userRecipes.value)
    } else {
      userRecipes.value = recipesData || []
      console.log('No recipes found or empty array')
    }
    
    // Load categories
    categories.value = await getCategories()
    console.log('Categories loaded:', categories.value)
  } catch (error) {
    console.error('Error loading data:', error)
    errorMessage.value = 'Failed to load recipes. Please try again.'
    userRecipes.value = []
    categories.value = []
  } finally {
    loading.value = false
  }
}
// Close dropdowns when clicking outside
const closeDropdowns = (event) => {
  if (!event.target.closest('.relative')) {
    showSortDropdown.value = false
  }
}

onMounted(() => {
  console.log('MyRecipes page mounted, auth state:', { 
    isAuthenticated: isAuthenticated.value, 
    user: user.value 
  })
  
  setTimeout(() => {
    isIconReady.value = true
  }, 150)
  
  loadData()
  document.addEventListener('click', closeDropdowns)
})

onUnmounted(() => {
  document.removeEventListener('click', closeDropdowns)
})
</script>