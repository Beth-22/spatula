<!-- pages/bookmarks.vue -->
<template>
  <div class="min-h-screen bg-gray-50">
    <Navigation />
    
    <!-- Header -->
    <div class="bg-white border-b">
      <div class="max-w-7xl mx-auto px-4 py-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-4xl font-serif font-bold text-gray-900 mb-2">My Bookmarks</h1>
            <p class="text-gray-600">Your saved recipes for later</p>
          </div>
          
          <!-- Filter & Sort Buttons -->
          <div class="flex gap-2">
            <!-- Sort Button -->
            <div class="relative">
              <button
                @click="showSortDropdown = !showSortDropdown"
                class="flex items-center gap-2 px-4 py-3 bg-white border border-gray-200 rounded-xl hover:border-gray-300 transition-colors"
              >
                <UIcon name="i-heroicons-arrows-up-down" class="h-5 w-5 text-gray-600" />
                <span class="text-gray-700">Sort</span>
              </button>
              
              <!-- Sort Dropdown -->
              <div v-if="showSortDropdown" class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-xl shadow-lg z-10">
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
              class="flex items-center gap-2 px-4 py-3 bg-orange-500 text-white rounded-xl hover:bg-orange-600 transition-colors"
            >
              <UIcon name="i-heroicons-funnel" class="h-5 w-5" />
              <span>Filters</span>
            </button>
          </div>
        </div>
        
        <!-- Quick Stats -->
        <div class="mt-6">
          <p class="text-gray-600" v-if="!loading">
            {{ filteredBookmarks.length }} bookmarked recipes
          </p>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Empty State -->
      <div v-if="!loading && bookmarks.length === 0" class="text-center py-16">
        <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <UIcon name="i-heroicons-bookmark" class="h-12 w-12 text-gray-400" />
        </div>
        <h3 class="text-2xl font-semibold text-gray-900 mb-2">No bookmarks yet</h3>
        <p class="text-gray-600 mb-6">Start saving your favorite recipes to see them here</p>
        <button 
          @click="$router.push('/recipes')"
          class="bg-orange-500 text-white px-6 py-3 rounded-lg hover:bg-orange-600 transition-colors font-semibold"
        >
          Browse Recipes
        </button>
      </div>

      <!-- Results Grid -->
      <div v-else>
        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <RecipeCardSkeleton v-for="n in 6" :key="n" />
        </div>
        
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <RecipeCard 
            v-for="bookmark in sortedBookmarks"
            :key="bookmark.recipe.id"
            :recipe="bookmark.recipe"
            :show-bookmark="true"
            :initial-bookmarked="true"
            @bookmark-removed="removeBookmark"
          />
        </div>

        <!-- No Results State -->
        <div v-if="!loading && bookmarks.length > 0 && filteredBookmarks.length === 0" class="text-center py-12">
          <UIcon name="i-heroicons-funnel-x-mark" class="h-16 w-16 text-gray-300 mx-auto mb-4" />
          <h3 class="text-xl font-semibold text-gray-900 mb-2">No bookmarks match your filters</h3>
          <p class="text-gray-600">Try adjusting your filters to see more results</p>
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
            <h2 class="text-xl font-semibold text-gray-900">Filter Bookmarks</h2>
            <button @click="showFilters = false" class="p-2 hover:bg-gray-100 rounded-lg">
              <UIcon name="i-heroicons-x-mark" class="h-5 w-5" />
            </button>
          </div>

          <!-- All Filters in One Card -->
          <div class="bg-gray-50 rounded-xl p-5 space-y-6">
            <!-- Categories -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Categories</h3>
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

            <!-- Preparation Time -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Preparation Time</h3>
              <div class="space-y-2">
                <label
                  v-for="range in timeRanges"
                  :key="range.value"
                  class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors"
                >
                  <input
                    type="radio"
                    :value="range.value"
                    v-model="filters.timeRange"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">{{ range.label }}</span>
                </label>
              </div>
            </div>

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
                    value="free"
                    v-model="filters.recipeType"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Free Recipes</span>
                </label>
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="premium"
                    v-model="filters.recipeType"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Premium Recipes</span>
                </label>
              </div>
            </div>

            <!-- Bookmarked Date -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Bookmarked</h3>
              <div class="space-y-2">
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="all"
                    v-model="filters.bookmarkDate"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Any Time</span>
                </label>
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="week"
                    v-model="filters.bookmarkDate"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Past Week</span>
                </label>
                <label class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors">
                  <input
                    type="radio"
                    value="month"
                    v-model="filters.bookmarkDate"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">Past Month</span>
                </label>
              </div>
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

    <Footer />
  </div>
</template>

<script setup>
// Composables
const { getBookmarks, removeBookmark: removeBookmarkApi } = useBookmarks()

// Data
const loading = ref(false)
const showFilters = ref(false)
const showSortDropdown = ref(false)
const bookmarks = ref([])
const categories = ref([])

// Filters
const filters = reactive({
  category: '',
  timeRange: '',
  recipeType: 'all',
  bookmarkDate: 'all'
})

// Sort
const selectedSort = ref('recent')
const sortOptions = [
  { label: 'Recently Added', value: 'recent' },
  { label: 'Highest Rating', value: 'ratings' },
  { label: 'Preparation Time', value: 'time' },
  { label: 'Alphabetical', value: 'alphabetical' }
]

// Computed
const filteredBookmarks = computed(() => {
  return bookmarks.value.filter(bookmark => {
    const recipe = bookmark.recipe
    
    // Category filter
    if (filters.category && recipe.category?.id !== filters.category) {
      return false
    }
    
    // Time range filter
    if (filters.timeRange) {
      const range = timeRanges.find(r => r.value === filters.timeRange)
      if (range.min && recipe.prep_time < range.min) return false
      if (range.max && recipe.prep_time > range.max) return false
    }
    
    // Recipe type filter
    if (filters.recipeType === 'free' && recipe.is_premium) return false
    if (filters.recipeType === 'premium' && !recipe.is_premium) return false
    
    // Bookmark date filter
    if (filters.bookmarkDate !== 'all') {
      const bookmarkDate = new Date(bookmark.created_at)
      const now = new Date()
      let cutoffDate = new Date()
      
      switch (filters.bookmarkDate) {
        case 'week':
          cutoffDate.setDate(now.getDate() - 7)
          break
        case 'month':
          cutoffDate.setMonth(now.getMonth() - 1)
          break
      }
      
      if (bookmarkDate < cutoffDate) return false
    }
    
    return true
  })
})

const sortedBookmarks = computed(() => {
  const bookmarksToSort = [...filteredBookmarks.value]
  
  switch (selectedSort.value) {
    case 'recent':
      return bookmarksToSort.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
    case 'ratings':
      return bookmarksToSort.sort((a, b) => {
        const aRating = calculateAverageRating(a.recipe.ratings || [])
        const bRating = calculateAverageRating(b.recipe.ratings || [])
        return bRating - aRating
      })
    case 'time':
      return bookmarksToSort.sort((a, b) => a.recipe.prep_time - b.recipe.prep_time)
    case 'alphabetical':
      return bookmarksToSort.sort((a, b) => a.recipe.title.localeCompare(b.recipe.title))
    default:
      return bookmarksToSort
  }
})

// Helper function to calculate average rating
const calculateAverageRating = (ratings) => {
  if (!ratings.length) return 0
  const total = ratings.reduce((sum, rating) => sum + rating.rating, 0)
  return total / ratings.length
}

// Methods
const selectSort = (sortValue) => {
  selectedSort.value = sortValue
  showSortDropdown.value = false
}

const removeBookmark = async (recipeId) => {
  try {
    // Remove from API
    await removeBookmarkApi(recipeId)
    
    // Remove from local state
    bookmarks.value = bookmarks.value.filter(b => b.recipe.id !== recipeId)
  } catch (error) {
    console.error('Error removing bookmark:', error)
  }
}

const clearFilters = () => {
  filters.category = ''
  filters.timeRange = ''
  filters.recipeType = 'all'
  filters.bookmarkDate = 'all'
}

// Close dropdowns when clicking outside
const closeDropdowns = (event) => {
  if (!event.target.closest('.relative')) {
    showSortDropdown.value = false
  }
}

// Fetch bookmarks
const loadBookmarks = async () => {
  loading.value = true
  try {
    const bookmarksData = await getBookmarks()
    bookmarks.value = bookmarksData
    
    // Extract unique categories from bookmarks
    const uniqueCategories = new Map()
    bookmarks.value.forEach(bookmark => {
      const cat = bookmark.recipe.category
      if (cat && !uniqueCategories.has(cat.id)) {
        uniqueCategories.set(cat.id, { id: cat.id, name: cat.name })
      }
    })
    categories.value = Array.from(uniqueCategories.values())
  } catch (error) {
    console.error('Error loading bookmarks:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadBookmarks()
  document.addEventListener('click', closeDropdowns)
})

onUnmounted(() => {
  document.removeEventListener('click', closeDropdowns)
})

const timeRanges = [
  { label: "Quick (< 30min)", value: "quick", max: 30 },
  { label: "Medium (30-60min)", value: "medium", min: 30, max: 60 },
  { label: "Long (1-2 hours)", value: "long", min: 60, max: 120 }
]
</script>