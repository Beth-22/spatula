<template>
  <div class="space-y-6">
    <!-- Categories Filter Panel -->
    <div class="border border-gray-200 rounded-lg">
      <button
        @click="toggleSection('categories')"
        class="flex items-center justify-between w-full p-4 text-left font-semibold text-gray-900"
      >
        <span>Categories</span>
        <Icon :icon="openSections.categories ? 'mdi:chevron-up' : 'mdi:chevron-down'" class="h-5 w-5" />
      </button>
      
      <div v-if="openSections.categories" class="p-4 border-t border-gray-200 space-y-2">
        <button
          v-for="category in categories"
          :key="category.id"
          @click="toggleCategory(category.id)"
          class="flex items-center justify-between w-full p-2 rounded-lg transition-colors"
          :class="selectedCategory === category.id 
            ? 'bg-orange-50 text-orange-700' 
            : 'text-gray-600 hover:bg-gray-50'"
        >
          <span>{{ category.name }}</span>
          <span class="text-sm text-gray-400">({{ category.recipes_aggregate.aggregate.count }})</span>
        </button>
      </div>
    </div>

    <!-- Preparation Time Filter Panel -->
    <div class="border border-gray-200 rounded-lg">
      <button
        @click="toggleSection('time')"
        class="flex items-center justify-between w-full p-4 text-left font-semibold text-gray-900"
      >
        <span>Preparation Time</span>
        <Icon :icon="openSections.time ? 'mdi:chevron-up' : 'mdi:chevron-down'" class="h-5 w-5" />
      </button>
      
      <div v-if="openSections.time" class="p-4 border-t border-gray-200 space-y-2">
        <label
          v-for="range in timeRanges"
          :key="range.value"
          class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-gray-50"
        >
          <input
            type="radio"
            :value="range.value"
            v-model="localTimeRange"
            @change="updateTimeRange"
            class="text-orange-500 focus:ring-orange-500"
          />
          <span class="ml-3 text-gray-700">{{ range.label }}</span>
        </label>
      </div>
    </div>

    <!-- Ingredients Filter Panel -->
    <div class="border border-gray-200 rounded-lg">
      <button
        @click="toggleSection('ingredients')"
        class="flex items-center justify-between w-full p-4 text-left font-semibold text-gray-900"
      >
        <span>Ingredients</span>
        <Icon :icon="openSections.ingredients ? 'mdi:chevron-up' : 'mdi:chevron-down'" class="h-5 w-5" />
      </button>
      
      <div v-if="openSections.ingredients" class="p-4 border-t border-gray-200 space-y-3">
        <!-- Search Ingredients -->
        <div class="relative">
          <input
            v-model="ingredientQuery"
            type="text"
            class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent text-sm"
            placeholder="Search ingredients..."
            @input="handleIngredientSearch"
          >
          <div v-if="ingredientSuggestions.length > 0" class="absolute z-10 w-full mt-1 bg-white border border-gray-200 rounded-lg shadow-lg max-h-60 overflow-auto">
            <button
              v-for="ingredient in ingredientSuggestions"
              :key="ingredient"
              @click="addIngredient(ingredient)"
              class="w-full px-3 py-2 text-left hover:bg-gray-50 text-sm"
            >
              {{ ingredient }}
            </button>
          </div>
        </div>
        
        <!-- Popular Ingredients Checkboxes -->
        <div class="space-y-2">
          <p class="text-sm font-medium text-gray-700">Popular Ingredients</p>
          <label
            v-for="ingredient in popularIngredients.slice(0, 5)"
            :key="ingredient"
            class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-gray-50"
          >
            <input
              type="checkbox"
              :value="ingredient"
              v-model="localIngredients"
              @change="updateIngredients"
              class="rounded text-orange-500 focus:ring-orange-500"
            />
            <span class="ml-3 text-sm text-gray-700">{{ ingredient }}</span>
          </label>
        </div>
        
        <!-- Selected Ingredients -->
        <div v-if="localIngredients.length > 0" class="flex flex-wrap gap-2">
          <span
            v-for="ingredient in localIngredients"
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
    </div>

    <!-- Clear Filters -->
    <button
      @click="clearFilters"
      class="w-full py-3 text-sm text-gray-600 hover:text-gray-800 border border-gray-200 rounded-lg hover:border-gray-300 transition-colors font-medium"
    >
      Clear All Filters
    </button>
  </div>
</template>

<script setup>
const props = defineProps({
  categories: Array,
  selectedCategory: String,
  selectedTimeRange: String,
  selectedIngredients: Array,
  popularIngredients: Array
})

const emit = defineEmits(['update:filters', 'close'])

// Local state
const openSections = reactive({
  categories: true,
  time: true,
  ingredients: true
})

const ingredientQuery = ref('')
const ingredientSuggestions = ref([])
const localTimeRange = ref(props.selectedTimeRange)
const localIngredients = ref([...props.selectedIngredients])

// Time ranges
const timeRanges = [
  { label: "Quick (< 30min)", value: "quick", max: 30 },
  { label: "Medium (30-60min)", value: "medium", min: 30, max: 60 },
  { label: "Long (1-2 hours)", value: "long", min: 60, max: 120 },
  { label: "All Day (> 2 hours)", value: "allDay", min: 120 }
]

// Methods
const toggleSection = (section) => {
  openSections[section] = !openSections[section]
}

const toggleCategory = (categoryId) => {
  const newCategory = props.selectedCategory === categoryId ? null : categoryId
  emit('update:filters', { category: newCategory })
}

const updateTimeRange = () => {
  emit('update:filters', { timeRange: localTimeRange.value })
}

const updateIngredients = () => {
  emit('update:filters', { ingredients: [...localIngredients.value] })
}

const handleIngredientSearch = useDebounceFn(() => {
  if (!ingredientQuery.value.trim()) {
    ingredientSuggestions.value = props.popularIngredients.slice(0, 5)
    return
  }
  
  const query = ingredientQuery.value.toLowerCase()
  ingredientSuggestions.value = props.popularIngredients
    .filter(ingredient => ingredient.toLowerCase().includes(query))
    .slice(0, 5)
}, 200)

const addIngredient = (ingredient) => {
  if (!localIngredients.value.includes(ingredient)) {
    localIngredients.value.push(ingredient)
    updateIngredients()
  }
  ingredientQuery.value = ''
  ingredientSuggestions.value = []
}

const removeIngredient = (ingredient) => {
  localIngredients.value = localIngredients.value.filter(i => i !== ingredient)
  updateIngredients()
}

const clearFilters = () => {
  localTimeRange.value = null
  localIngredients.value = []
  emit('update:filters', { 
    category: null, 
    timeRange: null, 
    ingredients: [] 
  })
  ingredientQuery.value = ''
  ingredientSuggestions.value = []
}

// Watch for prop changes
watch(() => props.selectedTimeRange, (newVal) => {
  localTimeRange.value = newVal
})

watch(() => props.selectedIngredients, (newVal) => {
  localIngredients.value = [...newVal]
})

// Load popular ingredients on mount
onMounted(() => {
  ingredientSuggestions.value = props.popularIngredients.slice(0, 5)
})
</script>