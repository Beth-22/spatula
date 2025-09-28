<template>
  <div v-if="show" class="fixed inset-0 z-50">
    <div
      class="absolute inset-0 bg-black bg-opacity-50"
      @click="$emit('close')"
    ></div>
    <div
      class="absolute right-0 top-0 h-full w-96 bg-white transform transition-transform"
    >
      <div class="p-6 h-full overflow-y-auto">
        <!-- Header -->
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-xl font-semibold text-gray-900">Filter Recipes</h2>
          <button
            @click="$emit('close')"
            class="p-2 hover:bg-gray-100 rounded-lg"
          >
            <Icon icon="mdi:close" class="h-5 w-5" />
          </button>
        </div>
        <!-- Filters -->
        <div class="space-y-6">
          <!-- Categories -->
          <div>
            <h3 class="font-semibold text-gray-900 mb-3">Categories</h3>
            <select
              v-model="localFilters.category"
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
            <select
              v-model="localFilters.timeRange"
              class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
            >
              <option value="">Any Time</option>
              <option value="quick">Quick (< 30min)</option>
              <option value="medium">Medium (30-60min)</option>
              <option value="long">Long (1-2 hours)</option>
              <option value="allDay">All Day (> 2 hours)</option>
            </select>
          </div>

          <!-- Recipe Type -->
          <div>
            <h3 class="font-semibold text-gray-900 mb-3">Recipe Type</h3>
            <select
              v-model="localFilters.premium"
              class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
            >
              <option value="">All Recipes</option>
              <option value="true">Premium Only</option>
              <option value="false">Free Only</option>
            </select>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-3 pt-6 border-t border-gray-200">
            <button
              @click="clearFilters"
              class="flex-1 py-3 text-gray-600 hover:text-gray-800 border border-gray-300 rounded-lg hover:border-gray-400 transition-colors font-medium"
            >
              Clear All
            </button>
            <button
              @click="applyFilters"
              class="flex-1 py-3 bg-gradient-to-r from-orange-500 to-red-500 text-white rounded-lg hover:from-orange-600 hover:to-red-600 transition-colors font-medium"
            >
              Apply Filters
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Define props
const props = defineProps({
  show: Boolean,
  filters: Object,
  categories: Array,
});

const emit = defineEmits(["update:filters", "close"]);

// Initialize localFilters with props.filters
const localFilters = ref({ ...props.filters });

// Watch for prop changes and update local state
watch(
  () => props.filters,
  (newVal) => {
    localFilters.value = { ...newVal };
  },
  { immediate: true, deep: true }
);

const applyFilters = () => {
  emit("update:filters", { ...localFilters.value });
  emit("close");
};

const clearFilters = () => {
  localFilters.value = {
    category: "",
    timeRange: "",
    premium: "",
  };
  emit("update:filters", { ...localFilters.value });
};
</script>
