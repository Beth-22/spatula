<!-- frontend/pages/recipes/upload.vue -->
<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 to-amber-50 py-8">
    <Navigation />

    <div class="max-w-4xl mx-auto px-4">
      <!-- Header -->
      <div class="text-center mb-8">
        <h1 class="text-4xl font-serif font-bold text-gray-900 mb-4">
          Share Your Recipe
        </h1>
        <p class="text-lg text-gray-600">
          Create something delicious and share it with the community
        </p>
      </div>

      <!-- Progress Steps -->
      <div class="flex justify-center mb-8">
        <div class="flex items-center space-x-8">
          <div
            v-for="(step, index) in steps"
            :key="step.id"
            class="flex items-center"
          >
            <div
              class="w-10 h-10 rounded-full flex items-center justify-center transition-all duration-300"
              :class="
                currentStep >= index
                  ? 'bg-orange-500 text-white'
                  : 'bg-gray-200 text-gray-500'
              "
            >
              <Icon v-if="isIconReady" :icon="step.icon" class="h-5 w-5" />
              <span v-else>{{ index + 1 }}</span>
            </div>
            <span
              class="ml-3 font-medium"
              :class="
                currentStep >= index ? 'text-orange-600' : 'text-gray-500'
              "
            >
              {{ step.label }}
            </span>
            <div
              v-if="index < steps.length - 1"
              class="w-16 h-0.5 ml-8"
              :class="currentStep > index ? 'bg-orange-500' : 'bg-gray-200'"
            ></div>
          </div>
        </div>
      </div>

      <!-- Form Container -->
      <div
        class="bg-white rounded-2xl shadow-lg border border-orange-100 overflow-hidden"
      >
        <!-- Basic Information Step -->
        <div v-if="currentStep === 0" class="p-8">
          <h2 class="text-2xl font-semibold text-gray-900 mb-6">
            Basic Information
          </h2>

          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- Recipe Title -->
            <div class="lg:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2"
                >Recipe Title *</label
              >
              <input
                v-model="recipe.title"
                type="text"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all"
                placeholder="e.g., Classic Chocolate Chip Cookies"
                maxlength="200"
              />
              <div class="text-xs text-gray-500 mt-1 text-right">
                {{ recipe.title.length }}/200
              </div>
            </div>

            <!-- Description -->
            <div class="lg:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2"
                >Description *</label
              >
              <textarea
                v-model="recipe.description"
                rows="4"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all resize-none"
                placeholder="Describe your recipe... What makes it special?"
              ></textarea>
            </div>

            <!-- Category -->
            <div class="lg:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2"
                >Category *</label
              >
              <select
                v-model="recipe.category_id"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              >
                <option value="">Select a category</option>
                <option
                  v-for="category in categories"
                  :key="category.id"
                  :value="category.id"
                >
                  {{ category.name }}
                </option>
              </select>
            </div>

            <!-- Preparation Time -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2"
                >Preparation Time (minutes) *</label
              >
              <input
                v-model.number="recipe.prep_time"
                type="number"
                min="1"
                max="1440"
                class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all"
                placeholder="30"
              />
            </div>

            <!-- Premium Toggle -->
            <div>
              <label class="flex items-center space-x-3">
                <input
                  v-model="recipe.is_premium"
                  type="checkbox"
                  class="w-4 h-4 text-orange-500 focus:ring-orange-500 border-gray-300 rounded"
                />
                <span class="text-sm font-medium text-gray-700"
                  >Make this a premium recipe</span
                >
              </label>

              <!-- Price Input (shown if premium) -->
              <div v-if="recipe.is_premium" class="mt-4">
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Price ($)</label
                >
                <input
                  v-model.number="recipe.price"
                  type="number"
                  min="0.01"
                  step="0.01"
                  class="w-32 px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all"
                  placeholder="4.99"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Images Step -->
        <div v-if="currentStep === 1" class="p-8">
          <h2 class="text-2xl font-semibold text-gray-900 mb-6">
            Recipe Images
          </h2>

          <!-- Featured Image Upload -->
          <div class="mb-8">
            <label class="block text-sm font-medium text-gray-700 mb-4"
              >Featured Image *</label
            >
            <ImageUploader
              v-model:images="recipeImages"
              :max-images="10"
              :featured-index="featuredImageIndex"
              @featured-changed="setFeaturedImage"
              @images-updated="handleImagesUpdate"
            />
          </div>

          <div class="text-sm text-gray-600">
            <p>• First image will be used as featured image</p>
            <p>• Supported formats: JPEG, PNG, WebP, GIF</p>
            <p>• Max file size: 10MB per image</p>
          </div>
        </div>

        <!-- Ingredients Step -->
        <div v-if="currentStep === 2" class="p-8">
          <h2 class="text-2xl font-semibold text-gray-900 mb-6">Ingredients</h2>

          <div class="space-y-4">
            <div
              v-for="(ingredient, index) in ingredients"
              :key="index"
              class="flex items-center gap-4"
            >
              <input
                v-model="ingredient.name"
                type="text"
                class="flex-1 px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                placeholder="Ingredient name"
              />
              <input
                v-model="ingredient.quantity"
                type="text"
                class="w-24 px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                placeholder="2"
              />
              <input
                v-model="ingredient.unit"
                type="text"
                class="w-32 px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                placeholder="cups"
              />
              <button
                @click="removeIngredient(index)"
                class="p-3 text-red-500 hover:bg-red-50 rounded-xl transition-colors"
                :disabled="ingredients.length === 1"
              >
                <Icon
                  v-if="isIconReady"
                  icon="mdi:trash-can-outline"
                  class="h-5 w-5"
                />
                <span v-else>🗑️</span>
              </button>
            </div>
          </div>

          <button
            @click="addIngredient"
            class="mt-4 flex items-center space-x-2 text-orange-600 hover:text-orange-700 font-medium"
          >
            <Icon v-if="isIconReady" icon="mdi:plus" class="h-5 w-5" />
            <span v-else>➕</span>
            <span>Add Another Ingredient</span>
          </button>
        </div>

        <!-- Instructions Step -->
        <div v-if="currentStep === 3" class="p-8">
          <h2 class="text-2xl font-semibold text-gray-900 mb-6">
            Instructions
          </h2>

          <div class="space-y-6">
            <div
              v-for="(step, index) in stepsData"
              :key="index"
              class="flex gap-4"
            >
              <!-- Step Number -->
              <div class="flex-shrink-0">
                <div
                  class="w-10 h-10 bg-orange-500 text-white rounded-full flex items-center justify-center font-semibold"
                >
                  {{ index + 1 }}
                </div>
              </div>

              <!-- Step Content -->
              <div class="flex-1">
                <textarea
                  v-model="step.instruction"
                  rows="3"
                  class="w-full px-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-orange-500 focus:border-transparent resize-none"
                  placeholder="Describe this step in detail..."
                ></textarea>
              </div>

              <!-- Remove Step Button -->
              <button
                @click="removeStep(index)"
                class="flex-shrink-0 p-3 text-red-500 hover:bg-red-50 rounded-xl transition-colors self-start"
                :disabled="stepsData.length === 1"
              >
                <Icon
                  v-if="isIconReady"
                  icon="mdi:trash-can-outline"
                  class="h-5 w-5"
                />
                <span v-else>🗑️</span>
              </button>
            </div>
          </div>

          <button
            @click="addStep"
            class="mt-6 flex items-center space-x-2 text-orange-600 hover:text-orange-700 font-medium"
          >
            <Icon v-if="isIconReady" icon="mdi:plus" class="h-5 w-5" />
            <span v-else>➕</span>
            <span>Add Another Step</span>
          </button>
        </div>

        <!-- Navigation Buttons -->
        <div
          class="flex justify-between items-center p-8 border-t border-gray-200"
        >
          <button
            @click="previousStep"
            v-if="currentStep > 0"
            class="px-6 py-3 text-gray-600 hover:text-gray-800 font-medium transition-colors"
          >
            Previous
          </button>
          <div v-else></div>

          <button
            @click="nextStep"
            class="px-8 py-3 bg-gradient-to-r from-orange-500 to-red-500 text-white rounded-xl hover:shadow-lg transition-all font-semibold"
            :disabled="!canProceed || loading"
          >
            {{
              currentStep === steps.length - 1 ? "Publish Recipe" : "Continue"
            }}
          </button>
        </div>
      </div>

      <!-- Loading Overlay -->
      <div
        v-if="loading"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      >
        <div class="bg-white rounded-2xl p-8 text-center">
          <div
            class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto mb-4"
          ></div>
          <p class="text-gray-700">Publishing your recipe...</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Composables
const { uploadImage } = useUpload();
const { getCategories } = useRecipeApi();

// Form data
const currentStep = ref(0);
const loading = ref(false);
const recipeImages = ref([]);
const featuredImageIndex = ref(0);
const isIconReady = ref(false);

// Recipe data
const recipe = reactive({
  title: "",
  description: "",
  category_id: "",
  prep_time: null,
  is_premium: false,
  price: null,
});

const ingredients = ref([{ name: "", quantity: "", unit: "" }]);
const stepsData = ref([{ instruction: "" }]);

// Categories
const categories = ref([]);

// Steps configuration
const steps = [
  { id: 0, label: "Basic Info", icon: "mdi:information" },
  { id: 1, label: "Images", icon: "mdi:image" },
  { id: 2, label: "Ingredients", icon: "mdi:food-apple" },
  { id: 3, label: "Instructions", icon: "mdi:book-open" },
];

// Computed
const canProceed = computed(() => {
  switch (currentStep.value) {
    case 0:
      return (
        recipe.title &&
        recipe.description &&
        recipe.category_id &&
        recipe.prep_time
      );
    case 1:
      return recipeImages.value.length > 0;
    case 2:
      return ingredients.value.every(
        (ing) => ing.name.trim() && ing.quantity.trim()
      );
    case 3:
      return stepsData.value.every((step) => step.instruction.trim());
    default:
      return false;
  }
});

// Methods
const nextStep = async () => {
  if (currentStep.value === steps.length - 1) {
    await publishRecipe();
  } else {
    currentStep.value++;
  }
};

const previousStep = () => {
  if (currentStep.value > 0) {
    currentStep.value--;
  }
};

const addIngredient = () => {
  ingredients.value.push({ name: "", quantity: "", unit: "" });
};

const removeIngredient = (index) => {
  if (ingredients.value.length > 1) {
    ingredients.value.splice(index, 1);
  }
};

const addStep = () => {
  stepsData.value.push({ instruction: "" });
};

const removeStep = (index) => {
  if (stepsData.value.length > 1) {
    stepsData.value.splice(index, 1);
  }
};

const setFeaturedImage = (index) => {
  featuredImageIndex.value = index;
};

const handleImagesUpdate = (images) => {
  recipeImages.value = images;
};

// Upload images to Supabase
const uploadImages = async () => {
  const uploadedUrls = [];

  for (const imageFile of recipeImages.value) {
    if (imageFile.file) {
      try {
        console.log('Uploading image:', imageFile.file.name);
        const response = await uploadImage(imageFile.file);
        
        if (response.url) {
          console.log('Image uploaded successfully:', response.url);
          uploadedUrls.push(response.url);
        } else {
          throw new Error('No URL returned from upload');
        }
      } catch (error) {
        console.error('Image upload failed:', error);
        throw new Error(`Failed to upload ${imageFile.file.name}: ${error.message}`);
      }
    } else if (imageFile.url) {
      // Already uploaded
      uploadedUrls.push(imageFile.url);
    }
  }

  return uploadedUrls;
};

const publishRecipe = async () => {
  if (!canProceed.value) return;

  loading.value = true;

  try {
    console.log('Starting recipe publication...');
    
    // 1. Upload images
    console.log('Uploading images...');
    const imageUrls = await uploadImages();
    console.log('Images uploaded:', imageUrls);

    if (imageUrls.length === 0) {
      throw new Error('No images were uploaded successfully');
    }

    const featuredImage = imageUrls[featuredImageIndex.value] || imageUrls[0];
    console.log('Featured image:', featuredImage);

    // 2. Create recipe using your backend endpoint
    const recipeData = {
      title: recipe.title,
      description: recipe.description,
      prep_time: recipe.prep_time,
      featured_image: featuredImage,
      all_images: imageUrls, // NEW: Send ALL image URLs
      is_premium: recipe.is_premium,
      price: recipe.is_premium ? recipe.price : null,
      category_id: recipe.category_id,
      ingredients: ingredients.value
        .filter(ing => ing.name.trim() && ing.quantity.trim())
        .map((ing) => ({
          name: ing.name.trim(),
          quantity: ing.quantity.trim(),
          unit: ing.unit.trim() || null,
        })),
      steps: stepsData.value
        .filter(step => step.instruction.trim())
        .map((step, index) => ({
          step_number: index + 1,
          instruction: step.instruction.trim(),
        })),
    };

    console.log("Sending recipe data:", recipeData);

    const authToken = useCookie('auth-token');
    
    const response = await $fetch("http://localhost:8081/actions/recipes/create", {
      method: "POST",
      body: JSON.stringify(recipeData),
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${authToken.value}`
      },
    });

    console.log("Recipe creation response:", response);

    if (response.recipe_id) {
      // Redirect to the new recipe
      await navigateTo(`/recipes/${response.recipe_id}`);
    } else {
      throw new Error("No recipe ID returned from server");
    }
  } catch (error) {
    console.error("Recipe creation failed:", error);
    alert("Failed to create recipe: " + (error.message || "Unknown error"));
  } finally {
    loading.value = false;
  }
};

// Load categories
onMounted(async () => {
  setTimeout(() => {
    isIconReady.value = true;
  }, 100);

  try {
    console.log("Loading categories...");
    const categoriesData = await getCategories();
    console.log("Categories loaded:", categoriesData);
    categories.value = categoriesData;
  } catch (error) {
    console.error("Failed to load categories:", error);
    categories.value = [];
  }
});
</script>