<!-- pages/categories/[slug].vue -->
<template>
  <div class="min-h-screen bg-white">
    <Navigation />
    
    <!-- Category Header -->
    <div class="relative h-64 bg-gradient-to-r from-orange-500 to-red-500">
      <div class="absolute inset-0 bg-black/40"></div>
      <div class="relative max-w-7xl mx-auto px-4 h-full flex items-center justify-center">
        <div class="text-center text-white">
          <h1 class="text-5xl font-serif font-bold mb-4">{{ categoryName }}</h1>
          <p class="text-xl">{{ recipes.length }} Recipes Available</p>
        </div>
      </div>
    </div>

    <!-- Recipes Grid -->
    <div class="max-w-7xl mx-auto px-4 py-12">
      <div v-if="recipesLoading" class="flex justify-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500"></div>
      </div>
      
      <div v-else-if="recipes.length === 0" class="text-center py-12">
        <p class="text-gray-600 text-lg">No recipes found in this category.</p>
        <button @click="$router.push('/recipes')" class="mt-4 bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600">
          Browse All Recipes
        </button>
      </div>
      
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8 justify-center">
        <RecipeCard 
          v-for="recipe in recipes" 
          :key="recipe.id"
          :recipe="recipe"
        />
      </div>
    </div>

    <Footer />
  </div>
</template>

<script setup>
const route = useRoute()
const router = useRouter()

const { getRecipesByCategory } = useRecipeApi();

const recipes = ref([]);
const recipesLoading = ref(true);
const categoryName = ref('');

// Fetch recipes for this category
onMounted(async () => {
  try {
    // Convert slug to category name
    const slug = route.params.slug;
    categoryName.value = slug.split('-').map(word => 
      word.charAt(0).toUpperCase() + word.slice(1)
    ).join(' ');

    recipes.value = await getRecipesByCategory(categoryName.value);
  } catch (error) {
    console.error('Error loading category recipes:', error);
  } finally {
    recipesLoading.value = false;
  }
});

useHead({
  title: `${categoryName.value} Recipes - Spatula`
})
</script>