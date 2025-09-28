// composables/useBookmarks.js - UPDATED QUERY
import { gql } from "graphql-tag";

export const useBookmarks = () => {
  const { $apollo } = useNuxtApp();
  const { user } = useAuth();
  const loading = ref(false);
  const error = ref(null);

  // Get user's bookmarks
  const getBookmarks = async () => {
    if (!user.value?.id) {
      throw new Error("Please log in to view bookmarks");
    }

    loading.value = true;
    error.value = null;

    try {
      const query = gql`
        query GetUserBookmarks($user_id: uuid!) {
          bookmarks(where: { user: { id: { _eq: $user_id } } }) {
            id
            created_at
            recipe {
              id
              title
              description
              prep_time
              is_premium
              price
              created_at
              user {
                id
                username
              }
              category {
                id
                name
              }
              recipe_images {
                id
                image_url
                is_featured
              }
              likes {
                id
              }
              ratings {
                id
                rating
              }
            }
          }
        }
      `;

      const result = await $apollo.query({
        query,
        variables: { user_id: user.value.id },
        fetchPolicy: "network-only",
      });

      console.log("Bookmarks data loaded:", result.data?.bookmarks);
      return result.data?.bookmarks || [];
    } catch (err) {
      console.error("Error getting bookmarks:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // ... rest of the functions remain the same
  const addBookmark = async (recipeId) => {
    if (!user.value?.id) {
      throw new Error("Please log in to bookmark recipes");
    }

    loading.value = true;
    error.value = null;

    try {
      const mutation = gql`
        mutation AddBookmark($recipe_id: uuid!, $user_id: uuid!) {
          insert_bookmarks_one(
            object: { recipe_id: $recipe_id, user_id: $user_id }
          ) {
            id
            created_at
            recipe {
              id
              title
            }
          }
        }
      `;

      const result = await $apollo.mutate({
        mutation,
        variables: {
          recipe_id: recipeId,
          user_id: user.value.id,
        },
      });

      return result.data?.insert_bookmarks_one;
    } catch (err) {
      console.error("Error adding bookmark:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // Remove bookmark
  const removeBookmark = async (recipeId) => {
    if (!user.value?.id) {
      throw new Error("Please log in to manage bookmarks");
    }

    loading.value = true;
    error.value = null;

    try {
      const mutation = gql`
        mutation RemoveBookmark($recipe_id: uuid!, $user_id: uuid!) {
          delete_bookmarks(
            where: {
              recipe_id: { _eq: $recipe_id }
              user: { id: { _eq: $user_id } }
            }
          ) {
            affected_rows
          }
        }
      `;

      const result = await $apollo.mutate({
        mutation,
        variables: {
          recipe_id: recipeId,
          user_id: user.value.id,
        },
      });

      return result.data?.delete_bookmarks;
    } catch (err) {
      console.error("Error removing bookmark:", err);
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // Toggle bookmark
  const toggleBookmark = async (recipeId) => {
    try {
      const checkQuery = gql`
        query CheckBookmark($recipe_id: uuid!, $user_id: uuid!) {
          bookmarks(
            where: {
              recipe_id: { _eq: $recipe_id }
              user: { id: { _eq: $user_id } }
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

      const alreadyBookmarked = checkResult.data?.bookmarks?.length > 0;

      if (alreadyBookmarked) {
        await removeBookmark(recipeId);
        return { bookmarked: false };
      } else {
        await addBookmark(recipeId);
        return { bookmarked: true };
      }
    } catch (err) {
      console.error("Error toggling bookmark:", err);
      throw err;
    }
  };

  // Check if recipe is bookmarked
  const getBookmarkStatus = async (recipeId) => {
    if (!user.value?.id) {
      return { bookmarked: false };
    }

    try {
      const query = gql`
        query GetBookmarkStatus($recipe_id: uuid!, $user_id: uuid!) {
          bookmarks(
            where: {
              recipe_id: { _eq: $recipe_id }
              user: { id: { _eq: $user_id } }
            }
          ) {
            id
            created_at
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

      return {
        bookmarked: result.data?.bookmarks?.length > 0,
        bookmark: result.data?.bookmarks?.[0],
      };
    } catch (err) {
      console.error("Error getting bookmark status:", err);
      return { bookmarked: false };
    }
  };

  return {
    loading,
    error,
    getBookmarks,
    addBookmark,
    removeBookmark,
    toggleBookmark,
    getBookmarkStatus,
  };
};
