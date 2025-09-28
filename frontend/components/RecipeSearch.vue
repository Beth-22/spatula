<!-- components/RecipeSearch.vue -->
<template>
  <div class="space-y-6">
    <!-- Search Bar -->
    <div class="relative">
      <div class="relative">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search recipes, ingredients..."
          class="w-full pl-10 pr-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all"
          @input="handleSearch"
        />
        <div class="absolute left-3 top-1/2 transform -translate-y-1/2">
          <Icon icon="mdi:magnify" class="h-5 w-5 text-gray-400" />
        </div>
        <button
          v-if="searchQuery"
          @click="clearSearch"
          class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
        >
          <Icon icon="mdi:close" class="h-5 w-5" />
        </button>
      </div>
    </div>

    <!-- Quick Filters -->
    <div class="flex flex-wrap gap-2">
      <button
        v-for="filter in quickFilters"
        :key="filter.value"
        @click="toggleQuickFilter(filter.value)"
        class="px-3 py-2 rounded-lg text-sm font-medium transition-colors"
        :class="activeQuickFilters.includes(filter.value) 
          ? 'bg-orange-500 text-white' 
          : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
      >
        {{ filter.label }}
      </button>
    </div>

    <!-- Use your existing SearchFilters component -->
    <SearchFilters
      :categories="categories"
      :selected-category="filters.category"
      :selected-time-range="filters.timeRange"
      :selected-ingredients="filters.ingredients"
      :popular-ingredients="popularIngredients"
      @update:filters="handleFilterUpdate"
    />
  </div>
</template>

<script setup>
const { getCategories, getPopularIngredients } = useRecipeApi();

const emit = defineEmits(['search', 'filter']);

const searchQuery = ref('');
const categories = ref([]);
const popularIngredients = ref([]);
const activeQuickFilters = ref([]);
const filters = ref({
  category: '',
  timeRange: '',
  ingredients: []
});

const quickFilters = [
  { label: 'Quick Meals', value: 'quick' },
  { label: 'Premium', value: 'premium' },
];

// Load categories and popular ingredients
onMounted(async () => {
  categories.value = await getCategories();
  popularIngredients.value = await getPopularIngredients();
});

// Debounced search
const handleSearch = useDebounceFn(() => {
  emit('search', searchQuery.value.trim());
}, 300);

const clearSearch = () => {
  searchQuery.value = '';
  emit('search', '');
};

const toggleQuickFilter = (filter) => {
  const index = activeQuickFilters.value.indexOf(filter);
  if (index > -1) {
    activeQuickFilters.value.splice(index, 1);
  } else {
    activeQuickFilters.value.push(filter);
  }
  
  // Update filters based on quick filters
  if (filter === 'quick') {
    filters.value.timeRange = activeQuickFilters.value.includes('quick') ? 'quick' : '';
  }
  if (filter === 'premium') {
    filters.value.premium = activeQuickFilters.value.includes('premium');
  }
  
  updateFilters();
};

const handleFilterUpdate = (updatedFilters) => {
  filters.value = { ...filters.value, ...updatedFilters };
  updateFilters();
};

const updateFilters = () => {
  const allFilters = {
    ...filters.value,
    quickFilters: [...activeQuickFilters.value]
  };
  emit('filter', allFilters);
};
</script>