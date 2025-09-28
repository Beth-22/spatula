<!-- components/CreatorsSection.vue -->
<template>
  <section class="bg-[#F5F5F7] py-16">
    <div class="max-w-7xl mx-auto px-4">
      <!-- components/CreatorsSection.vue - Update the heading section -->
<div class="text-center mb-12">
  <h2 class="text-4xl font-serif font-bold text-gray-900 mb-4">
    {{ activeTab === 'categories' ? 'Browse by Category' : 'Browse by Creator' }}
  </h2>
  <p class="text-lg text-gray-600">
    {{ activeTab === 'categories' ? 'Discover recipes by category' : 'Discover recipes from talented chefs' }}
  </p>
</div>

      <!-- Toggle between Categories and Creators -->
      <div class="flex justify-center mb-8">
        <div class="bg-gray-100 rounded-lg p-1 inline-flex">
          <button
            @click="activeTab = 'categories'"
            class="px-6 py-2 rounded-md transition-all duration-200 font-medium"
            :class="activeTab === 'categories' ? 'bg-white text-orange-600 shadow-sm' : 'text-gray-600 hover:text-gray-900'"
          >
            Categories
          </button>
          <button
            @click="activeTab = 'creators'"
            class="px-6 py-2 rounded-md transition-all duration-200 font-medium"
            :class="activeTab === 'creators' ? 'bg-white text-orange-600 shadow-sm' : 'text-gray-600 hover:text-gray-900'"
          >
            Creators
          </button>
        </div>
      </div>

      <!-- Categories Content -->
      <div v-if="activeTab === 'categories'" class="fade-in">
        <div v-if="categoriesLoading" class="flex justify-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500"></div>
          <p class="ml-4 text-gray-600">Loading categories...</p>
        </div>

        <div v-else-if="categories.length > 0" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div 
            v-for="category in categories" 
            :key="category.id"
            class="relative rounded-xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer h-80 group hover:scale-105"
            @click="navigateToCategory(category)"
          >
            <img 
              :src="getCategoryImage(category)" 
              :alt="category.name"
              class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent"></div>
            <div class="absolute bottom-0 left-0 right-0 p-6 text-white">
              <h3 class="font-bold text-2xl mb-2">{{ category.name }}</h3>
            </div>
          </div>
        </div>
      </div>

      <!-- Creators Content -->
      <div v-else class="fade-in">
        <div v-if="creatorsLoading" class="flex justify-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500"></div>
          <p class="ml-4 text-gray-600">Loading creators...</p>
        </div>

        <div v-else-if="creators.length > 0" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          <div 
            v-for="creator in creators" 
            :key="creator.id"
            class="relative rounded-xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer h-80 group hover:scale-105"
            @click="navigateToCreator(creator)"
          >
            <!-- Creator Image -->
            <img 
              :src="getCreatorImage(creator)" 
              :alt="creator.username"
              class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
            />
            
            <!-- Gradient Overlay -->
            <div class="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent"></div>
            
            <!-- Creator Content -->
            <div class="absolute bottom-0 left-0 right-0 p-6 text-white">
              <h3 class="font-bold text-2xl mb-2"> {{ creator.username }}</h3>
              <p class="text-gray-200 text-sm">{{ creator.recipe_count }} recipes</p>
            </div>
          </div>
        </div>

        <!-- Error State -->
        <div v-else-if="creatorsError" class="text-center py-12">
          <div class="bg-red-50 border border-red-200 rounded-lg p-6 max-w-md mx-auto">
            <div class="text-red-600 mb-4">
              <UIcon name="i-heroicons-exclamation-triangle" class="w-12 h-12 mx-auto" />
            </div>
            <p class="text-red-800 font-medium mb-2">Failed to load creators</p>
            <p class="text-red-600 text-sm mb-4">{{ creatorsError }}</p>
            <button @click="loadCreators" class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors">
              Try Again
            </button>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-12">
          <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6 max-w-md mx-auto">
            <UIcon name="i-heroicons-user-group" class="w-12 h-12 text-yellow-600 mx-auto mb-4" />
            <p class="text-yellow-800 font-medium">No creators found</p>
            <p class="text-yellow-600 text-sm mt-2">Creators will appear here once they publish recipes.</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
// Composables
const { getCategories } = useCategories();
const { getCreators } = useCreators();

// Reactive data
const activeTab = ref('categories');
const categories = ref([]);
const creators = ref([]);
const categoriesLoading = ref(true);
const creatorsLoading = ref(true);
const creatorsError = ref('');

// Load categories
const loadCategories = async () => {
  try {
    categoriesLoading.value = true;
    const categoriesData = await getCategories();
    categories.value = categoriesData || [];
  } catch (err) {
    console.error('Error loading categories:', err);
  } finally {
    categoriesLoading.value = false;
  }
};

// Load creators
const loadCreators = async () => {
  try {
    creatorsLoading.value = true;
    creatorsError.value = '';
    const creatorsData = await getCreators();
    creators.value = creatorsData || [];
  } catch (err) {
    console.error('Error loading creators:', err);
    creatorsError.value = err.message || 'Failed to load creators. Please try again.';
  } finally {
    creatorsLoading.value = false;
  }
};

// Navigation
const navigateToCategory = (category) => {
  const slug = category.name.toLowerCase().replace(/\s+/g, '-');
  navigateTo(`/categories/${slug}`);
};

const navigateToCreator = (creator) => {
  navigateTo(`/creators/${creator.username}`);
};

// Get creator image with fallback
const getCreatorImage = (creator) => {
  // Use display_image if available from the composable
  if (creator.display_image) {
    return creator.display_image;
  }
  
  // Use recipe image if available
  if (creator.recipes?.[0]?.recipe_images?.[0]?.image_url) {
    return creator.recipes[0].recipe_images[0].image_url;
  }
  
  // Fallback to generic chef image
  return 'https://images.unsplash.com/photo-1600565193348-f74bd3c7ccdf?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80';
};

// Get category image (same as before)
const getCategoryImage = (category) => {
  if (category.image) return category.image;
  
  const fallbackImages = {
    'asian': 'https://images.unsplash.com/photo-1555939594-58d7cb561ad1?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80',
    'italian': 'https://images.unsplash.com/photo-1535930749574-1399327ce78f?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80',
    // ... other fallbacks
  };
  
  return fallbackImages[category.name.toLowerCase()] || 'https://images.unsplash.com/photo-1546069901-ba9599a7e63c?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80';
};

// Load data on mount
onMounted(() => {
  loadCategories();
  loadCreators();
});

// Watch tab changes to ensure data is loaded
watch(activeTab, (newTab) => {
  if (newTab === 'creators' && creators.value.length === 0 && !creatorsError.value) {
    loadCreators();
  }
});
</script>

<style scoped>
.fade-in {
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>