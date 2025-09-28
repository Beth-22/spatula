<!-- components/CategoriesSection.vue -->
<template>
  <section class="bg-[#F5F5F7] py-16">
    <div class="max-w-7xl mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-4xl font-serif font-bold text-gray-900 mb-4">
          Browse Categories
        </h2>
        <p class="text-lg text-gray-600">Discover recipes by category</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center">
        <div
          class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500"
        ></div>
        <p class="ml-4 text-gray-600">Loading categories...</p>
      </div>

      <!-- Categories Grid -->
      <div
        v-else-if="categories.length > 0"
        class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6"
      >
        <div
          v-for="category in categories"
          :key="category.id"
          class="relative rounded-xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer h-80 group hover:scale-105"
          @click="navigateToCategory(category)"
        >
          <!-- Category Image -->
          <img
            :src="getCategoryImage(category)"
            :alt="category.name"
            class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
          />

          <!-- Gradient Overlay -->
          <div
            class="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent"
          ></div>

          <!-- Category Content -->
          <div class="absolute bottom-0 left-0 right-0 p-6 text-white">
            <h3 class="font-bold text-2xl mb-2">{{ category.name }}</h3>
          </div>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <div
          class="bg-red-50 border border-red-200 rounded-lg p-6 max-w-md mx-auto"
        >
          <div class="text-red-600 mb-4">
            <svg
              class="w-12 h-12 mx-auto"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.35 16.5c-.77.833.192 2.5 1.732 2.5z"
              ></path>
            </svg>
          </div>
          <p class="text-red-800 font-medium mb-2">Failed to load categories</p>
          <p class="text-red-600 text-sm mb-4">{{ error }}</p>
          <button
            @click="loadCategories"
            class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-12">
        <div
          class="bg-yellow-50 border border-yellow-200 rounded-lg p-6 max-w-md mx-auto"
        >
          <div class="text-yellow-600 mb-4">
            <svg
              class="w-12 h-12 mx-auto"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.35 16.5c-.77.833.192 2.5 1.732 2.5z"
              ></path>
            </svg>
          </div>
          <p class="text-yellow-800 font-medium">No categories found</p>
          <p class="text-yellow-600 text-sm mt-2">
            Categories will appear here once they are created in the database.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
// Use your existing composable
const { getCategories } = useCategories();

const categories = ref([]);
const loading = ref(true);
const error = ref("");

// Load categories on component mount
const loadCategories = async () => {
  try {
    loading.value = true;
    error.value = "";
    console.log("Loading categories...");

    const categoriesData = await getCategories();
    console.log("Categories data received:", categoriesData);

    categories.value = categoriesData || [];
  } catch (err) {
    console.error("Error loading categories:", err);
    error.value = err.message || "Failed to load categories. Please try again.";
  } finally {
    loading.value = false;
  }
};

// Navigate to category page
const navigateToCategory = (category) => {
  const slug = category.name.toLowerCase().replace(/\s+/g, "-");
  navigateTo(`/categories/${slug}`);
};

// Get category image with fallback
const getCategoryImage = (category) => {
  // If category has an image in database, use it
  if (category.image) {
    return category.image;
  }

  // Fallback images based on category name
  const fallbackImages = {
    asian:
      "https://images.unsplash.com/photo-1555939594-58d7cb561ad1?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    italian:
      "https://images.unsplash.com/photo-1535930749574-1399327ce78f?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    dessert:
      "https://images.unsplash.com/photo-1563729784474-d77dbb933a9e?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    vegetarian:
      "https://images.unsplash.com/photo-1512621776951-a57141f2eefd?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    mexican:
      "https://images.unsplash.com/photo-1565299624946-b28f40a0ca4b?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    mediterranean:
      "https://images.unsplash.com/photo-1476224203421-9ac39bcb3327?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    american:
      "https://images.unsplash.com/photo-1546833999-b9f581a1996d?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    indian:
      "https://images.unsplash.com/photo-1585937421612-70a008356fbe?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    chinese:
      "https://images.unsplash.com/photo-1563245372-f21724e3856d?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    japanese:
      "https://images.unsplash.com/photo-1580822184713-fc5400e7fe10?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    thai: "https://images.unsplash.com/photo-1559314809-0f1556aae4ac?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
    french:
      "https://images.unsplash.com/photo-1504674900247-0877df9cc836?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80",
  };

  const lowerCaseName = category.name.toLowerCase();
  return (
    fallbackImages[lowerCaseName] ||
    "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=80"
  );
};

// Load categories when component is mounted
onMounted(() => {
  loadCategories();
});
</script>
