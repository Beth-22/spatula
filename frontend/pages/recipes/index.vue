<!-- pages/recipes/index.vue -->
<template>
  <div class="min-h-screen bg-gray-50">
    <Navigation />
    
    <!-- Page Header -->
    <div class="bg-white border-b">
      <div class="max-w-7xl mx-auto px-4 py-8">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
          <div>
            <h1 class="text-4xl font-serif font-bold text-gray-900 mb-2">All Recipes</h1>
            <p class="text-gray-600">{{ filteredRecipes.length }} recipes found</p>
          </div>
          
          <!-- Sort and Filter Controls -->
          <div class="flex gap-3">
            <!-- Sort Button -->
            <div class="relative">
              <button
                @click="showSortDropdown = !showSortDropdown"
                class="flex items-center gap-2 px-4 py-2 bg-white border border-gray-200 rounded-lg hover:border-gray-300 transition-colors"
              >
                <div v-if="isIconReady">
                  <Icon icon="mdi:sort" class="h-5 w-5 text-gray-600" />
                </div>
                <span v-else class="text-gray-600">↕️</span>
                <span class="text-gray-700">{{ getSortLabel(selectedSort) }}</span>
                <div v-if="isIconReady">
                  <Icon :icon="showSortDropdown ? 'mdi:chevron-up' : 'mdi:chevron-down'" class="h-4 w-4 text-gray-400" />
                </div>
                <span v-else class="text-gray-400">{{ showSortDropdown ? '▲' : '▼' }}</span>
              </button>
              
              <div v-if="showSortDropdown" class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-lg shadow-lg z-20">
                <button
                  v-for="option in sortOptions"
                  :key="option.value"
                  @click="selectSort(option.value)"
                  class="w-full px-4 py-2 text-left hover:bg-gray-50 transition-colors first:rounded-t-lg last:rounded-b-lg"
                  :class="selectedSort === option.value ? 'bg-orange-50 text-orange-700' : 'text-gray-700'"
                >
                  {{ option.label }}
                </button>
              </div>
            </div>
            
            <!-- Filter Toggle Button -->
            <button
              @click="showFilterPanel = true"
              class="flex items-center gap-2 px-4 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 transition-colors"
            >
              <div v-if="isIconReady">
                <Icon icon="mdi:filter" class="h-5 w-5" />
              </div>
              <span v-else>🔧</span>
              <span>Filters</span>
              <span v-if="hasActiveFilters" class="bg-white text-orange-500 text-xs rounded-full h-5 w-5 flex items-center justify-center">
                {{ activeFilterCount }}
              </span>
            </button>
          </div>
        </div>

        <!-- Active Filters -->
        <div v-if="hasActiveFilters" class="flex flex-wrap gap-2 mt-4">
          <span
            v-for="filter in activeFilters"
            :key="filter.key"
            class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-orange-100 text-orange-800"
          >
            {{ filter.label }}
            <button @click="removeFilter(filter.key)" class="ml-2 hover:text-orange-900">
              <div v-if="isIconReady">
                <Icon icon="mdi:close" class="h-4 w-4" />
              </div>
              <span v-else>✕</span>
            </button>
          </span>
          <button 
            @click="clearAllFilters"
            class="text-sm text-gray-600 hover:text-gray-800 underline"
          >
            Clear all
          </button>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500"></div>
      </div>

      <!-- Recipes Grid -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 justify-center">
        <RecipeCard 
          v-for="recipe in sortedRecipes" 
          :key="recipe.id"
          :recipe="recipe"
        />
      </div>

      <!-- Empty State -->
      <div v-if="!loading && sortedRecipes.length === 0" class="text-center py-16">
        <div class="max-w-md mx-auto">
          <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
            <div v-if="isIconReady">
              <Icon icon="mdi:chef-hat" class="h-12 w-12 text-gray-400" />
            </div>
            <span v-else class="text-2xl">👨‍🍳</span>
          </div>
          <h3 class="text-xl font-semibold text-gray-900 mb-2">No recipes found</h3>
          <p class="text-gray-600 mb-6">Try adjusting your filters to see more results</p>
          <button 
            @click="clearAllFilters"
            class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors"
          >
            Clear Filters
          </button>
        </div>
      </div>
    </div>

    <!-- Filter Panel -->
    <FilterPanel
      :show="showFilterPanel"
      :filters="filters"
      :categories="categories"
      @update:filters="handleFilterUpdate"
      @close="showFilterPanel = false"
    />

    <Footer />
  </div>
</template>

<script setup>
const { getAllRecipes, getCategories } = useRecipeApi()

const recipes = ref([])
const categories = ref([])
const loading = ref(true)
const showFilterPanel = ref(false)
const showSortDropdown = ref(false)
const isIconReady = ref(false)

// Filters
const filters = reactive({
  category: '',
  timeRange: '',
  premium: '',
  difficulty: '',
  rating: ''
})

// Sort
const selectedSort = ref('newest')
const sortOptions = [
  { label: 'Newest First', value: 'newest' },
  { label: 'Oldest First', value: 'oldest' },
  { label: 'Preparation Time (Shortest)', value: 'time_asc' },
  { label: 'Preparation Time (Longest)', value: 'time_desc' },
  { label: 'A-Z', value: 'alphabetical' }
]

// Computed properties
const filteredRecipes = computed(() => {
  return recipes.value.filter(recipe => {
    // Category filter
    if (filters.category && recipe.category?.id !== filters.category) return false
    
    // Time range filter
    if (filters.timeRange) {
      const prepTime = recipe.prep_time || 0
      switch (filters.timeRange) {
        case 'quick': if (prepTime >= 30) return false; break
        case 'medium': if (prepTime < 30 || prepTime > 60) return false; break
        case 'long': if (prepTime < 60 || prepTime > 120) return false; break
        case 'allDay': if (prepTime <= 120) return false; break
      }
    }
    
    // Premium filter
    if (filters.premium !== '') {
      const isPremium = filters.premium === 'true'
      if (recipe.is_premium !== isPremium) return false
    }
    
    return true
  })
})

const sortedRecipes = computed(() => {
  const recipesToSort = [...filteredRecipes.value]
  
  switch (selectedSort.value) {
    case 'oldest':
      return recipesToSort.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
    case 'time_asc':
      return recipesToSort.sort((a, b) => (a.prep_time || 0) - (b.prep_time || 0))
    case 'time_desc':
      return recipesToSort.sort((a, b) => (b.prep_time || 0) - (a.prep_time || 0))
    case 'alphabetical':
      return recipesToSort.sort((a, b) => a.title.localeCompare(b.title))
    case 'newest':
    default:
      return recipesToSort.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  }
})

const hasActiveFilters = computed(() => {
  return Object.values(filters).some(value => value !== '')
})

const activeFilterCount = computed(() => {
  return Object.values(filters).filter(value => value !== '').length
})

const activeFilters = computed(() => {
  const active = []
  
  if (filters.category) {
    const category = categories.value.find(cat => cat.id === filters.category)
    if (category) {
      active.push({ key: 'category', label: `Category: ${category.name}` })
    }
  }
  
  if (filters.timeRange) {
    const timeLabels = {
      quick: 'Quick (<30min)',
      medium: 'Medium (30-60min)',
      long: 'Long (1-2 hours)',
      allDay: 'All Day (>2 hours)'
    }
    active.push({ key: 'timeRange', label: `Time: ${timeLabels[filters.timeRange]}` })
  }
  
  if (filters.premium !== '') {
    const type = filters.premium === 'true' ? 'Premium Only' : 'Free Only'
    active.push({ key: 'premium', label: type })
  }
  
  return active
})

// Methods
const getSortLabel = (sortValue) => {
  const option = sortOptions.find(opt => opt.value === sortValue)
  return option ? option.label : 'Sort'
}

const selectSort = (sortValue) => {
  selectedSort.value = sortValue
  showSortDropdown.value = false
}

const handleFilterUpdate = (newFilters) => {
  Object.assign(filters, newFilters)
  showFilterPanel.value = false
}

const removeFilter = (filterKey) => {
  filters[filterKey] = ''
}

const clearAllFilters = () => {
  Object.keys(filters).forEach(key => {
    filters[key] = ''
  })
}

// Load data
const loadRecipes = async () => {
  try {
    loading.value = true
    recipes.value = await getAllRecipes()
    categories.value = await getCategories()
  } catch (error) {
    console.error('Error loading recipes:', error)
  } finally {
    loading.value = false
  }
}

// Close dropdown when clicking outside
const closeDropdowns = (event) => {
  if (!event.target.closest('.relative')) {
    showSortDropdown.value = false
  }
}

// Initialize
onMounted(async () => {
  setTimeout(() => {
    isIconReady.value = true
  }, 150)
  
  await loadRecipes()
  document.addEventListener('click', closeDropdowns)
})

onUnmounted(() => {
  document.removeEventListener('click', closeDropdowns)
})

useHead({
  title: 'All Recipes - Spatula'
})
</script>