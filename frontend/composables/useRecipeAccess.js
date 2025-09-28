import { CHECK_RECIPE_ACCESS } from "~/graphql/queries";

export const useRecipeAccess = () => {
  const { $apollo } = useNuxtApp();
  const { user } = useAuth();

  const checkAccess = async (recipeId) => {
    if (!user.value) {
      return {
        hasAccess: false,
        isPremium: false,
        reason: "Not authenticated",
      };
    }

    try {
      const { data } = await $apollo.query({
        query: CHECK_RECIPE_ACCESS,
        variables: {
          recipeId: recipeId,
          userId: user.value.id,
        },
        fetchPolicy: "network-only",
      });

      const recipe = data.recipes_by_pk;
      if (!recipe) {
        return {
          hasAccess: false,
          isPremium: false,
          reason: "Recipe not found",
        };
      }

      // If recipe is not premium, allow access
      if (!recipe.is_premium) {
        return {
          hasAccess: true,
          isPremium: false,
          price: recipe.price,
        };
      }

      // Check if user has purchased this recipe
      const hasPurchased = data.purchases.length > 0;

      return {
        hasAccess: hasPurchased,
        isPremium: true,
        price: recipe.price,
        reason: hasPurchased ? "" : "Purchase required",
      };
    } catch (error) {
      console.error("Error checking recipe access:", error);
      return {
        hasAccess: false,
        isPremium: false,
        reason: "Error checking access",
      };
    }
  };

  return {
    checkAccess,
  };
};
