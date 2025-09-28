<!-- pages/search.vue -->
<template>
  <div class="min-h-screen bg-gray-50">
    <Navigation />
    
    <!-- Search Header -->
    <div class="bg-white border-b">
      <div class="max-w-7xl mx-auto px-4 py-6">
        <div class="flex items-center justify-between gap-4">
          <!-- Search Bar -->
          <div class="relative flex-1 max-w-2xl">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <Icon icon="mdi:magnify" class="h-5 w-5 text-gray-400" />
            </div>
            <input
              v-model="searchQuery"
              type="text"
              class="block w-full pl-10 pr-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent bg-white text-lg"
              placeholder="Search recipes, ingredients, creators..."
              @input="handleSearch"
              @keyup.enter="performSearch"
            >
          </div>
          
          <!-- Filter & Sort Buttons -->
          <div class="flex gap-2">
            <!-- Sort Button -->
            <div class="relative">
              <button
                @click="showSortDropdown = !showSortDropdown"
                class="flex items-center gap-2 px-4 py-3 bg-white border border-gray-200 rounded-xl hover:border-gray-300 transition-colors"
              >
                <Icon icon="mdi:sort" class="h-5 w-5 text-gray-600" />
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
              <Icon icon="mdi:filter" class="h-5 w-5" />
              <span>Filters</span>
            </button>
          </div>
        </div>
        
        <!-- Quick Stats -->
        <div class="text-center mt-4">
          <p class="text-gray-600" v-if="!loading && searchQuery">
            Found {{ filteredRecipes.length }} recipes for "{{ searchQuery }}"
          </p>
          <p v-else-if="loading" class="text-gray-600">Searching...</p>
          <p v-else class="text-gray-600">Enter a search term to find recipes</p>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Results Grid -->
      <div>
        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <RecipeCardSkeleton v-for="n in 6" :key="n" />
        </div>
        
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <RecipeCard 
            v-for="recipe in sortedRecipes"
            :key="recipe.id"
            :recipe="recipe"
            :show-bookmark="true"
          />
        </div>

        <!-- Empty State -->
        <div v-if="!loading && searchQuery && filteredRecipes.length === 0" class="text-center py-12">
          <Icon icon="mdi:chef-hat" class="h-16 w-16 text-gray-300 mx-auto mb-4" />
          <h3 class="text-xl font-semibold text-gray-900 mb-2">No recipes found</h3>
          <p class="text-gray-600">Try adjusting your search or filters</p>
          <button 
            @click="clearSearch"
            class="mt-4 bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors"
          >
            Clear Search
          </button>
        </div>

        <!-- Initial State -->
        <div v-if="!loading && !searchQuery" class="text-center py-12">
          <Icon icon="mdi:magnify" class="h-16 w-16 text-gray-300 mx-auto mb-4" />
          <h3 class="text-xl font-semibold text-gray-900 mb-2">Search Recipes</h3>
          <p class="text-gray-600">Enter a search term to find delicious recipes</p>
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
            <h2 class="text-xl font-semibold text-gray-900">Filters</h2>
            <button @click="showFilters = false" class="p-2 hover:bg-gray-100 rounded-lg">
              <Icon icon="mdi:close" class="h-5 w-5" />
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
                @change="applyFilters"
              >
                <option value="">All Categories</option>
                <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                  {{ cat.name }}
                </option>
              </select>
            </div>

            <!-- Creators -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Creators</h3>
              <select 
                v-model="filters.creator"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                @change="applyFilters"
              >
                <option value="">All Creators</option>
                <option v-for="creator in creators" :key="creator.id" :value="creator.id">
                  {{ creator.username }}
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
                    @change="applyFilters"
                    class="text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-3 text-gray-700">{{ range.label }}</span>
                </label>
              </div>
            </div>

            <!-- Popular Ingredients -->
            <div>
              <h3 class="font-semibold text-gray-900 mb-3">Popular Ingredients</h3>
              <div class="grid grid-cols-2 gap-2">
                <label
                  v-for="ingredient in popularIngredients"
                  :key="ingredient"
                  class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-white transition-colors"
                >
                  <input
                    type="checkbox"
                    :value="ingredient"
                    v-model="filters.ingredients"
                    @change="applyFilters"
                    class="rounded text-orange-500 focus:ring-orange-500"
                  />
                  <span class="ml-2 text-sm text-gray-700">{{ ingredient }}</span>
                </label>
              </div>
              
              <!-- Selected Ingredients -->
              <div v-if="filters.ingredients.length > 0" class="flex flex-wrap gap-2 mt-3">
                <span
                  v-for="ingredient in filters.ingredients"
                  :key="ingredient"
                  class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-orange-100 text-orange-800"
                >
                  {{ ingredient }}
                  <button @click="removeIngredient(ingredient)" class="ml-1 hover:text-orange-900">
                    <Icon icon="mdi:close" class="h-3 w-3" />
                  </button>
                </span>
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
const { searchRecipes, getCategories, getCreators, getPopularIngredients } = useSearch()

// Data
const searchQuery = ref('')
const loading = ref(false)
const showFilters = ref(false)
const showSortDropdown = ref(false)
const recipes = ref([])
const categories = ref([])
const creators = ref([])
const popularIngredients = ref([])

// Filters
const filters = reactive({
  category: '',
  creator: '',
  timeRange: '',
  ingredients: []
})

// Sort
const selectedSort = ref('newest')
const sortOptions = [
  { label: 'Newest First', value: 'newest' },
  { label: 'Highest Rating', value: 'ratings' },
  { label: 'Quickest First', value: 'time' }
]

// Computed
const filteredRecipes = computed(() => {
  return recipes.value
})

const sortedRecipes = computed(() => {
  const recipesToSort = [...filteredRecipes.value]
  
  switch (selectedSort.value) {
    case 'ratings':
      return recipesToSort.sort((a, b) => {
        const aRatings = a.ratings || []
        const bRatings = b.ratings || []
        const aAvg = aRatings.length > 0 ? aRatings.reduce((sum, r) => sum + r.rating, 0) / aRatings.length : 0
        const bAvg = bRatings.length > 0 ? bRatings.reduce((sum, r) => sum + r.rating, 0) / bRatings.length : 0
        return bAvg - aAvg
      })
    case 'time':
      return recipesToSort.sort((a, b) => a.prep_time - b.prep_time)
    case 'newest':
    default:
      return recipesToSort.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  }
})

// Methods
const performSearch = async () => {
  if (!searchQuery.value.trim() && Object.keys(filters).every(key => 
    !filters[key] || (Array.isArray(filters[key]) && filters[key].length === 0)
  )) {
    recipes.value = []
    return
  }

  loading.value = true
  try {
    const results = await searchRecipes(searchQuery.value, filters)
    recipes.value = results
    console.log('Search results:', results)
  } catch (error) {
    console.error('Search error:', error)
    recipes.value = []
  } finally {
    loading.value = false
  }
}

const handleSearch = useDebounceFn(() => {
  performSearch()
}, 500)

const applyFilters = () => {
  performSearch()
}

const selectSort = (sortValue) => {
  selectedSort.value = sortValue
  showSortDropdown.value = false
}

const removeIngredient = (ingredient) => {
  filters.ingredients = filters.ingredients.filter(i => i !== ingredient)
  performSearch()
}

const clearFilters = () => {
  filters.category = ''
  filters.creator = ''
  filters.timeRange = ''
  filters.ingredients = []
  if (searchQuery.value || recipes.value.length > 0) {
    performSearch()
  }
}

const clearSearch = () => {
  searchQuery.value = ''
  recipes.value = []
  clearFilters()
}

// Close dropdowns when clicking outside
const closeDropdowns = (event) => {
  if (!event.target.closest('.relative')) {
    showSortDropdown.value = false
  }
}

// Load initial data
const loadInitialData = async () => {
  try {
    const [categoriesData, creatorsData] = await Promise.all([
      getCategories(),
      getCreators()
    ])
    
    categories.value = categoriesData
    creators.value = creatorsData
    popularIngredients.value = getPopularIngredients()
  } catch (error) {
    console.error('Error loading initial data:', error)
  }
}

onMounted(() => {
  loadInitialData()
  document.addEventListener('click', closeDropdowns)
  
  // Check if there's a search query in the URL
  const route = useRoute()
  if (route.query.q) {
    searchQuery.value = route.query.q
    performSearch()
  }
})

onUnmounted(() => {
  document.removeEventListener('click', closeDropdowns)
})

const timeRanges = [
  { label: "Any Time", value: "" },
  { label: "Quick (< 30min)", value: "quick" },
  { label: "Medium (30-60min)", value: "medium" },
  { label: "Long (1-2 hours)", value: "long" },
  { label: "All Day (> 2 hours)", value: "allDay" }
]
</script>