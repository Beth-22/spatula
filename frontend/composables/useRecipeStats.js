// composables/useRecipeStats.js
import { gql } from "graphql-tag";

export const useRecipeStats = () => {
  const { $apollo } = useNuxtApp();

  // Get likes count for a recipe
  const getLikesCount = async (recipeId) => {
    try {
      const query = gql`
        query GetLikesCount($recipe_id: uuid!) {
          likes(where: { recipe_id: { _eq: $recipe_id } }) {
            id
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: { recipe_id: recipeId },
        fetchPolicy: "network-only",
      });

      return result.data?.likes?.length || 0;
    } catch (error) {
      console.error("Error getting likes count:", error);
      return 0;
    }
  };

  // Get ratings data for a recipe
  const getRatingsData = async (recipeId) => {
    try {
      const query = gql`
        query GetRatingsData($recipe_id: uuid!) {
          ratings(where: { recipe_id: { _eq: $recipe_id } }) {
            rating
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: { recipe_id: recipeId },
        fetchPolicy: "network-only",
      });

      const ratings = result.data?.ratings || [];
      const count = ratings.length;

      if (count === 0) {
        return { averageRating: 0, ratingsCount: 0 };
      }

      const totalRating = ratings.reduce(
        (sum, r) => sum + parseFloat(r.rating),
        0
      );
      const averageRating = totalRating / count;

      return {
        averageRating: parseFloat(averageRating.toFixed(1)),
        ratingsCount: count,
      };
    } catch (error) {
      console.error("Error getting ratings data:", error);
      return { averageRating: 0, ratingsCount: 0 };
    }
  };

  // Get stats for multiple recipes (optimized for lists)
  const getRecipesStats = async (recipeIds) => {
    try {
      const query = gql`
        query GetRecipesStats($recipe_ids: [uuid!]!) {
          likes(where: { recipe_id: { _in: $recipe_ids } }) {
            recipe_id
          }
          ratings(where: { recipe_id: { _in: $recipe_ids } }) {
            recipe_id
            rating
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: { recipe_ids: recipeIds },
        fetchPolicy: "network-only",
      });

      const likes = result.data?.likes || [];
      const ratings = result.data?.ratings || [];

      const stats = {};

      // Calculate likes count for each recipe
      recipeIds.forEach((recipeId) => {
        const recipeLikes = likes.filter((like) => like.recipe_id === recipeId);
        const recipeRatings = ratings.filter(
          (rating) => rating.recipe_id === recipeId
        );

        const likesCount = recipeLikes.length;
        const ratingsCount = recipeRatings.length;

        let averageRating = 0;
        if (ratingsCount > 0) {
          const totalRating = recipeRatings.reduce(
            (sum, r) => sum + parseFloat(r.rating),
            0
          );
          averageRating = parseFloat((totalRating / ratingsCount).toFixed(1));
        }

        stats[recipeId] = {
          likesCount,
          averageRating,
          ratingsCount,
        };
      });

      return stats;
    } catch (error) {
      console.error("Error getting recipes stats:", error);
      return {};
    }
  };

  return {
    getLikesCount,
    getRatingsData,
    getRecipesStats,
  };
};
