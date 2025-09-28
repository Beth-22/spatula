// composables/useCreators.js
import { gql } from "graphql-tag";

export const useCreators = () => {
  const { $apollo } = useNuxtApp();
  const loading = ref(false);
  const error = ref(null);

  // Get all creators (users who have published recipes)
  const getCreators = async () => {
    loading.value = true;
    error.value = null;

    try {
      const query = gql`
        query GetCreators {
          users(where: { recipes: { id: { _is_null: false } } }) {
            id
            username
            created_at
            recipes {
              id
              recipe_images {
                image_url
              }
            }
          }
        }
      `;

      const result = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });

      // Manually compute recipe counts since aggregate is not available
      const creators = result.data?.users || [];

      return creators.map((creator) => ({
        ...creator,
        // Calculate recipe count manually
        recipe_count: creator.recipes?.length || 0,
        // Use first recipe's image for display
        display_image: creator.recipes?.[0]?.recipe_images?.[0]?.image_url,
      }));
    } catch (err) {
      console.error("Error getting creators:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // Get creator by ID
  const getCreatorById = async (creatorId) => {
    loading.value = true;
    error.value = null;

    try {
      const query = gql`
        query GetCreator($creatorId: uuid!) {
          users_by_pk(id: $creatorId) {
            id
            username
            created_at
            recipes {
              id
              title
              description
              prep_time
              is_premium
              price
              created_at
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
        variables: { creatorId },
        fetchPolicy: "network-only",
      });

      const creator = result.data?.users_by_pk;
      if (creator) {
        // Create a new object instead of modifying the existing one
        return {
          ...creator,
          recipe_count: creator.recipes?.length || 0,
        };
      }

      return creator;
    } catch (err) {
      console.error("Error getting creator:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // Get creator by username (for slug routes)
  const getCreatorByUsername = async (username) => {
    loading.value = true;
    error.value = null;

    try {
      const query = gql`
        query GetCreatorByUsername($username: String!) {
          users(where: { username: { _eq: $username } }) {
            id
            username
            created_at
            recipes {
              id
              title
              description
              prep_time
              is_premium
              price
              created_at
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
        variables: { username },
        fetchPolicy: "network-only",
      });

      const creator = result.data?.users?.[0] || null;
      if (creator) {
        // Create a new object instead of modifying the existing one
        return {
          ...creator,
          recipe_count: creator.recipes?.length || 0,
        };
      }

      return creator;
    } catch (err) {
      console.error("Error getting creator by username:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  return {
    loading,
    error,
    getCreators,
    getCreatorById,
    getCreatorByUsername,
  };
};
