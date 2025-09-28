// composables/useRatings.js - CORRECTED VERSION
import { gql } from "graphql-tag";

export const useRatings = () => {
  const { $apollo } = useNuxtApp();
  const { user } = useAuth();
  const loading = ref(false);
  const error = ref(null);

  const setRating = async (recipeId, rating) => {
    if (!user.value?.id) {
      throw new Error("Please log in to rate recipes");
    }

    loading.value = true;
    error.value = null;

    try {
      // For INSERT: use user_id (as confirmed by introspection)
      const mutation = gql`
        mutation SetRating($object: ratings_insert_input!) {
          insert_ratings_one(
            object: $object
            on_conflict: {
              constraint: ratings_user_id_recipe_id_key
              update_columns: [rating, updated_at]
            }
          ) {
            id
            rating
            recipe_id
            # For the response, we don't need user_id since it's not in the ratings type
          }
        }
      `;

      const variables = {
        object: {
          rating: parseFloat(rating),
          recipe_id: recipeId,
          user_id: user.value.id, // Use user_id for INSERT
        },
      };

      console.log("Setting rating with variables:", variables);

      const result = await $apollo.mutate({
        mutation,
        variables,
      });

      console.log("Rating mutation successful:", result);

      const stats = await getRatingStats(recipeId);

      return {
        userRating: rating,
        averageRating: stats.averageRating,
        ratingCount: stats.ratingCount,
      };
    } catch (err) {
      console.error("Error setting rating:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const deleteRating = async (recipeId) => {
    if (!user.value?.id) {
      throw new Error("Please log in to manage ratings");
    }

    loading.value = true;
    error.value = null;

    try {
      // For DELETE: use the relationship field 'user' since 'user_id' doesn't exist in ratings type
      const mutation = gql`
        mutation DeleteRating($recipe_id: uuid!) {
          delete_ratings(
            where: {
              recipe_id: { _eq: $recipe_id }
              user: { id: { _eq: "${user.value.id}" } }
            }
          ) {
            affected_rows
          }
        }
      `;

      const variables = {
        recipe_id: recipeId,
      };

      console.log("Deleting rating with variables:", variables);

      const result = await $apollo.mutate({
        mutation,
        variables,
      });

      console.log("Rating deletion successful:", result);

      const stats = await getRatingStats(recipeId);

      return {
        userRating: 0,
        averageRating: stats.averageRating,
        ratingCount: stats.ratingCount,
      };
    } catch (err) {
      console.error("Error deleting rating:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const getRatingStats = async (recipeId) => {
    try {
      // For QUERY: use the relationship field 'user'
      const query = gql`
        query GetRatingStats($recipe_id: uuid!) {
          ratings(where: { recipe_id: { _eq: $recipe_id } }) {
            id
            rating
            recipe_id
            created_at
            user {
              id
              username
            }
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: { recipe_id: recipeId },
        fetchPolicy: "network-only",
      });

      const ratings = result.data?.ratings || [];

      // Manual calculation of average and count
      let totalRating = 0;
      let ratingCount = ratings.length;

      if (ratingCount > 0) {
        ratings.forEach((rating) => {
          totalRating += parseFloat(rating.rating);
        });
      }

      const averageRating = ratingCount > 0 ? totalRating / ratingCount : 0;

      return {
        averageRating: parseFloat(averageRating.toFixed(1)),
        ratingCount: ratingCount,
      };
    } catch (err) {
      console.error("Error getting rating stats:", err);
      return { averageRating: 0, ratingCount: 0 };
    }
  };

  const getRatingStatus = async (recipeId) => {
    if (!user.value?.id) {
      return await getRatingStats(recipeId);
    }

    try {
      // For QUERY: use the relationship field 'user' to filter
      const query = gql`
        query GetRatingStatus($recipe_id: uuid!) {
          ratings(where: { recipe_id: { _eq: $recipe_id } }) {
            id
            rating
            recipe_id
            user {
              id
            }
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: { recipe_id: recipeId },
        fetchPolicy: "network-only",
      });

      const allRatings = result.data?.ratings || [];

      // Calculate average manually from all ratings
      let totalRating = 0;
      let ratingCount = allRatings.length;

      if (ratingCount > 0) {
        allRatings.forEach((rating) => {
          totalRating += parseFloat(rating.rating);
        });
      }

      const averageRating = ratingCount > 0 ? totalRating / ratingCount : 0;

      // Find user's rating by matching through the user relationship
      const userRatingObj = allRatings.find(
        (rating) => rating.user?.id === user.value.id
      );
      const userRating = userRatingObj?.rating || 0;

      return {
        userRating: userRating,
        averageRating: parseFloat(averageRating.toFixed(1)),
        ratingCount: ratingCount,
      };
    } catch (err) {
      console.error("Error getting rating status:", err);
      return { userRating: 0, averageRating: 0, ratingCount: 0 };
    }
  };

  return {
    loading,
    error,
    setRating,
    deleteRating,
    getRatingStatus,
  };
};
