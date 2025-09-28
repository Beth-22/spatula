<template>
  <div class="min-h-screen bg-white">
    <!-- Loading State -->
    <div v-if="loading" class="min-h-screen flex items-center justify-center">
      <div class="text-center">
        <div
          class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto mb-4"
        ></div>
        <p class="text-gray-600">Loading recipe...</p>
      </div>
    </div>

    <!-- Error State -->
    <div
      v-else-if="error"
      class="min-h-screen flex items-center justify-center"
    >
      <div class="text-center">
        <div
          class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4"
        >
          <span class="text-2xl">❌</span>
        </div>
        <h2 class="text-xl font-semibold text-gray-900 mb-2">
          Recipe Not Found
        </h2>
        <p class="text-gray-600 mb-4">{{ error }}</p>
        <button
          @click="goBack"
          class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600"
        >
          Go Back
        </button>
      </div>
    </div>

    <!-- Recipe Content - ChefSteps Style with Blur Effect -->
    <div v-else-if="recipe" class="min-h-screen bg-white">
      <!-- Hero Section -->
      <div class="relative h-96 bg-gray-900">
        <!-- Background Image -->
        <div class="absolute inset-0 overflow-hidden">
          <img
            :src="getRecipeImage(recipe)"
            :alt="recipe.title"
            class="w-full h-full object-cover opacity-90"
          />
          <div class="absolute inset-0 bg-black/40"></div>
        </div>

        <!-- Navigation -->
        <nav class="relative z-10">
          <div class="max-w-6xl mx-auto px-6 py-6">
            <div class="flex justify-between items-center">
              <NuxtLink
  to="/my-recipes"
  class="flex items-center text-white hover:text-gray-200 transition-colors text-lg font-medium"
>
  <Icon icon="mdi:arrow-left" class="h-5 w-5 mr-2" />
  Back to Recipes
</NuxtLink>

              <div class="flex items-center space-x-4">
                <!-- Purchase Status -->
                <div
                  v-if="hasPurchased"
                  class="bg-green-500 text-white px-3 py-1 rounded-full text-sm font-medium"
                >
                  PURCHASED
                </div>
              </div>
            </div>
          </div>
        </nav>

        <!-- Recipe Title Centered -->
        <div
          class="relative z-10 max-w-4xl mx-auto px-6 h-full flex flex-col items-center justify-center text-center"
        >
          <h1
            class="text-5xl lg:text-6xl font-light text-white mb-4 leading-tight"
          >
            {{ recipe.title }}
          </h1>
          <div class="flex items-center space-x-4 text-gray-200">
            <Icon icon="mdi:chef-hat" class="h-6 w-6" />
            <span class="text-xl">{{
              recipe.user?.username || "Unknown Chef"
            }}</span>
          </div>

          <!-- Price Badge -->
          <div
            class="mt-4 bg-gradient-to-r from-orange-500 to-red-500 text-white text-lg px-6 py-2 rounded-full font-semibold shadow-lg"
          >
            ${{ recipe.price }}
          </div>
        </div>
      </div>

      <!-- Main Content with Blur Effect -->
      <div class="relative">
        <!-- Visible Content (Above the blur) -->
        <div class="max-w-4xl mx-auto px-6 py-12">
          <!-- Description Section - Always Visible -->
          <div class="prose prose-lg max-w-none mb-8">
            <h2 class="text-3xl font-light text-gray-900 mb-6">
              About This Recipe
            </h2>
            <p class="text-gray-700 text-lg leading-relaxed mb-8">
              {{ recipe.description }}
            </p>

            <!-- QUICK TEST BUTTON - Added right below description -->
            <div
              class="mt-8 p-6 bg-gradient-to-r from-orange-50 to-red-50 rounded-xl border border-orange-200"
            >
              <h3 class="text-xl font-semibold text-orange-800 mb-3">
                Quick Purchase Test
              </h3>
              <p class="text-orange-700 mb-4">
                Test the Chapa payment integration immediately:
              </p>
              <button
                @click="handlePurchase"
                class="bg-gradient-to-r from-orange-500 to-red-500 text-white px-8 py-3 rounded-lg font-semibold hover:from-orange-600 hover:to-red-600 transition-all transform hover:scale-105 shadow-lg"
              >
                🚀 Buy Recipe Now - ${{ recipe.price }}
              </button>
              <p class="text-orange-600 text-sm mt-3">
                This will open the payment modal
              </p>
            </div>
          </div>

          <!-- Recipe Details Card -->
          <div class="bg-gray-50 rounded-lg p-6 border border-gray-200 mb-12">
            <h3 class="text-xl font-medium text-gray-900 mb-6">
              Recipe Details
            </h3>
            <div class="grid grid-cols-2 gap-4">
              <div class="flex items-center justify-between">
                <span class="text-gray-600">Prep time</span>
                <span class="font-medium text-gray-900">{{
                  formatTime(recipe.prep_time)
                }}</span>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-600">Rating</span>
                <div class="flex items-center space-x-1">
                  <span class="font-medium text-gray-900">4.5</span>
                  <Icon icon="mdi:star" class="h-4 w-4 text-yellow-400" />
                </div>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-600">Likes</span>
                <div class="flex items-center space-x-1">
                  <span class="font-medium text-gray-900">24</span>
                  <Icon icon="mdi:heart" class="h-4 w-4 text-red-400" />
                </div>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-600">Category</span>
                <span class="font-medium text-gray-900 capitalize">{{
                  recipe.category?.name || "General"
                }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Blurred Content Section (Everything below description) -->
        <div
          class="relative"
          :class="hasPurchased ? '' : 'premium-blur-section'"
        >
          <!-- Ingredients Section -->
          <div class="max-w-4xl mx-auto px-6 py-8">
            <h2 class="text-3xl font-light text-gray-900 mb-8">Ingredients</h2>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div
                v-for="ingredient in recipe.ingredients"
                :key="ingredient.id"
                class="flex items-center justify-between py-3 border-b border-gray-100"
              >
                <span class="text-gray-700 text-lg">{{ ingredient.name }}</span>
                <span class="text-gray-900 font-medium text-lg"
                  >{{ ingredient.quantity }} {{ ingredient.unit || "" }}</span
                >
              </div>
            </div>
          </div>

          <!-- Steps Section -->
          <div class="max-w-4xl mx-auto px-6 py-8">
            <h2 class="text-3xl font-light text-gray-900 mb-8">Steps</h2>

            <div class="space-y-8">
              <div
                v-for="step in recipe.steps"
                :key="step.id"
                class="flex space-x-6"
              >
                <!-- Step Number -->
                <div class="flex-shrink-0">
                  <div
                    class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center"
                  >
                    <span class="text-gray-600 font-bold text-xl">{{
                      step.step_number
                    }}</span>
                  </div>
                </div>

                <!-- Step Content -->
                <div class="flex-1">
                  <p class="text-gray-700 text-lg leading-relaxed">
                    {{ step.instruction }}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- Comments Section -->
          <div class="max-w-4xl mx-auto px-6 py-8">
            <CommentsSection :recipe-id="recipe.id" />
          </div>

          <!-- Purchase Overlay (Shown when not purchased) -->
          <div
            v-if="!hasPurchased"
            class="absolute inset-0 bg-white/80 backdrop-blur-sm flex items-center justify-center z-10"
          >
            <div
              class="text-center p-8 bg-white rounded-2xl shadow-2xl border border-gray-200 max-w-md mx-4"
            >
              <div
                class="w-16 h-16 bg-gradient-to-r from-orange-500 to-red-500 rounded-full flex items-center justify-center mx-auto mb-4"
              >
                <Icon icon="mdi:lock" class="h-8 w-8 text-white" />
              </div>

              <h3 class="text-2xl font-bold text-gray-900 mb-2">
                Premium Recipe
              </h3>
              <p class="text-gray-600 mb-6">
                Purchase this recipe to unlock all ingredients, steps, and
                community features
              </p>

              <div class="space-y-4">
                <div class="text-3xl font-bold text-gray-900">
                  ${{ recipe.price }}
                </div>

                <button
                  @click="handlePurchase"
                  class="w-full bg-gradient-to-r from-orange-500 to-red-500 text-white py-3 rounded-lg font-semibold hover:from-orange-600 hover:to-red-600 transition-all transform hover:scale-105"
                >
                  Buy Recipe Now
                </button>

                <button
                  @click="goBack"
                  class="w-full bg-gray-200 text-gray-700 py-3 rounded-lg font-semibold hover:bg-gray-300 transition-colors"
                >
                  Browse Other Recipes
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Payment Modal -->
    <div
      v-if="showPaymentModal"
      class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
    >
      <div class="bg-white rounded-lg p-6 w-96 max-w-md mx-4">
        <h3 class="text-lg font-bold mb-4">Purchase Recipe</h3>
        <p class="mb-4">
          You are about to purchase "{{ recipe?.title }}" for ${{
            recipe?.price
          }}
        </p>

        <div class="mb-4">
          <label class="block text-sm font-medium mb-2">Phone Number</label>
          <input
            v-model="phoneNumber"
            type="tel"
            placeholder="Enter your phone number"
            class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
          />
        </div>

        <div class="flex space-x-3">
          <button
            @click="proceedToPayment"
            class="flex-1 bg-green-500 text-white py-3 rounded-lg font-semibold hover:bg-green-600 transition-colors"
          >
            Pay Now
          </button>
          <button
            @click="showPaymentModal = false"
            class="flex-1 bg-gray-500 text-white py-3 rounded-lg font-semibold hover:bg-gray-600 transition-colors"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { GET_RECIPE_BY_ID } from "~/graphql/queries";

const route = useRoute();
const router = useRouter();
const { id } = route.params;
const { isAuthenticated } = useAuth();

// Get Apollo client from Nuxt app
const { $apollo } = useNuxtApp();
// Get runtime config
const config = useRuntimeConfig();

// Reactive data
const loading = ref(true);
const error = ref("");
const recipe = ref(null);
const hasPurchased = ref(false);
const showPaymentModal = ref(false);
const phoneNumber = ref("");

// Get recipe image
const getRecipeImage = (recipe) => {
  if (recipe.recipe_images && recipe.recipe_images.length > 0) {
    return recipe.recipe_images[0].image_url;
  }
  return "/images/recipe-placeholder.jpg";
};

// Load recipe data using Apollo Client
const loadRecipeData = async () => {
  try {
    loading.value = true;
    error.value = "";

    const { data } = await $apollo.query({
      query: GET_RECIPE_BY_ID,
      variables: { id },
      fetchPolicy: "network-only",
    });

    recipe.value = data?.recipes_by_pk;

    if (!recipe.value) {
      error.value = "Recipe not found";
      return;
    }

    console.log("Premium recipe data loaded:", recipe.value);
    hasPurchased.value = false;
  } catch (err) {
    console.error("Error loading recipe:", err);
    error.value = "Failed to load recipe. Please try again.";
  } finally {
    loading.value = false;
  }
};

// Handle purchase
const handlePurchase = () => {
  if (!isAuthenticated.value) {
    navigateTo("/login");
    return;
  }
  showPaymentModal.value = true;
};

// Proceed to payment - UPDATED WITH BETTER AUTH HANDLING
const proceedToPayment = async () => {
  if (!phoneNumber.value) {
    alert("Please enter your phone number");
    return;
  }

  try {
    // Debug: Check if we're actually authenticated and have the token
    console.log("Auth status:", isAuthenticated.value);

    // Get the token from cookie manually to verify it exists
    const token = useCookie("auth-token").value;
    console.log("JWT Token from cookie:", token ? "Exists" : "Missing");

    if (!token) {
      alert("No authentication token found. Please log in again.");
      navigateTo("/login");
      return;
    }

    console.log(
      "Initiating payment to:",
      `${config.public.authEndpoint}/payment/initiate`
    );

    // Use fetch instead of $fetch for better control over credentials
    const response = await fetch(
      `${config.public.authEndpoint}/payment/initiate`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`, // Manually add Authorization header
        },
        body: JSON.stringify({
          recipe_id: recipe.value.id,
          amount: recipe.value.price,
          currency: "ETB",
          phone_number: phoneNumber.value,
        }),
        credentials: "include", // Still include cookies
      }
    );

    if (!response.ok) {
      if (response.status === 401) {
        throw new Error("UNAUTHORIZED");
      }
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();

    if (data.success) {
      console.log("Payment initiated successfully:", data);
      window.location.href = data.checkout_url;
    } else {
      alert("Failed to initiate payment: " + data.message);
    }
  } catch (error) {
    console.error("Payment initiation failed:", error);

    if (error.message === "UNAUTHORIZED") {
      // Clear the invalid token and redirect to login
      useCookie("auth-token").value = null;
      alert("Your session has expired. Please log in again.");
      navigateTo("/login");
    } else if (error.statusCode === 404) {
      alert("Payment endpoint not found. Please check backend server.");
    } else {
      alert(
        "Payment initiation failed. Please try again. Check console for details."
      );
    }
  } finally {
    showPaymentModal.value = false;
  }
};

const goBack = () => {
  router.back();
};

const formatTime = (minutes) => {
  if (!minutes) return "0m";
  const hours = Math.floor(minutes / 60);
  const mins = minutes % 60;
  return hours > 0 ? `${hours}h ${mins}m` : `${mins}m`;
};

onMounted(() => {
  loadRecipeData();
});
</script>
<style scoped>
.premium-blur-section {
  filter: blur(8px);
  pointer-events: none;
  user-select: none;
}

.premium-blur-section > * {
  filter: blur(8px);
}

/* Ensure the purchase overlay is not blurred */
.premium-blur-section > div:last-child {
  filter: none !important;
  pointer-events: all !important;
}

.transition-all {
  transition: all 0.3s ease;
}
</style>
