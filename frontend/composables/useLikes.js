// composables/useLikes.js
import { gql } from "graphql-tag";

export const useLikes = () => {
  const { $apollo } = useNuxtApp();
  const { user } = useAuth();
  const loading = ref(false);
  const error = ref(null);

  // Toggle like for a recipe - Pure GraphQL with manual counting
  const toggleLike = async (recipeId) => {
    if (!user.value?.id) {
      throw new Error("Please log in to like recipes");
    }

    loading.value = true;
    error.value = null;

    try {
      // First check if like exists
      const checkQuery = gql`
        query CheckLikeExists($recipe_id: uuid!, $user_id: uuid!) {
          likes(
            where: {
              recipe_id: { _eq: $recipe_id }
              user_id: { _eq: $user_id }
            }
          ) {
            id
          }
        }
      `;

      const checkResult = await $apollo.query({
        query: checkQuery,
        variables: {
          recipe_id: recipeId,
          user_id: user.value.id,
        },
        fetchPolicy: "network-only",
      });

      const alreadyLiked = checkResult.data?.likes?.length > 0;

      if (alreadyLiked) {
        // Unlike the recipe
        const unlikeMutation = gql`
          mutation UnlikeRecipe($recipe_id: uuid!, $user_id: uuid!) {
            delete_likes(
              where: {
                recipe_id: { _eq: $recipe_id }
                user_id: { _eq: $user_id }
              }
            ) {
              affected_rows
            }
          }
        `;

        await $apollo.mutate({
          mutation: unlikeMutation,
          variables: {
            recipe_id: recipeId,
            user_id: user.value.id,
          },
        });
      } else {
        // Like the recipe
        const likeMutation = gql`
          mutation LikeRecipe($object: likes_insert_input!) {
            insert_likes_one(object: $object) {
              id
            }
          }
        `;

        await $apollo.mutate({
          mutation: likeMutation,
          variables: {
            object: {
              recipe_id: recipeId,
              user_id: user.value.id,
            },
          },
        });
      }

      // Get updated like count manually (since aggregate doesn't work)
      const countQuery = gql`
        query GetLikesCount($recipe_id: uuid!) {
          likes(where: { recipe_id: { _eq: $recipe_id } }) {
            id
          }
        }
      `;

      const countResult = await $apollo.query({
        query: countQuery,
        variables: { recipe_id: recipeId },
        fetchPolicy: "network-only",
      });

      const likeCount = countResult.data?.likes?.length || 0;

      return {
        liked: !alreadyLiked,
        likeCount: likeCount,
      };
    } catch (err) {
      console.error("Error toggling like:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // Get like status for a recipe
  const getLikeStatus = async (recipeId) => {
    if (!user.value?.id) {
      // For non-authenticated users, just get the count
      try {
        const count = await getLikesCount(recipeId);
        return { liked: false, likeCount: count };
      } catch (error) {
        return { liked: false, likeCount: 0 };
      }
    }

    try {
      const query = gql`
        query GetLikeStatus($recipe_id: uuid!, $user_id: uuid!) {
          # Get all likes for this recipe to count manually
          all_likes: likes(where: { recipe_id: { _eq: $recipe_id } }) {
            id
            user_id
          }
          # Check if current user liked this recipe
          user_likes: likes(
            where: {
              recipe_id: { _eq: $recipe_id }
              user_id: { _eq: $user_id }
            }
          ) {
            id
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: {
          recipe_id: recipeId,
          user_id: user.value.id,
        },
        fetchPolicy: "network-only",
      });

      // Manual count since aggregate doesn't work
      const likeCount = result.data?.all_likes?.length || 0;
      const liked = result.data?.user_likes?.length > 0;

      return {
        liked: liked,
        likeCount: likeCount,
      };
    } catch (err) {
      console.error("Error getting like status:", err);
      return { liked: false, likeCount: 0 };
    }
  };

  // Get likes count for a recipe (public - no auth required)
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

      // Manual count - just return the array length
      return result.data?.likes?.length || 0;
    } catch (err) {
      console.error("Error getting likes count:", err);
      return 0;
    }
  };

  return {
    loading,
    error,
    toggleLike,
    getLikeStatus,
    getLikesCount,
  };
};
