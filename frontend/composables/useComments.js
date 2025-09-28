import { gql } from "graphql-tag";

export const useComments = () => {
  const { $apollo } = useNuxtApp();
  const { user } = useAuth();
  const loading = ref(false);
  const error = ref(null);

  // Get comments for a recipe - REMOVE AGGREGATE
  const getRecipeComments = async (recipeId, page = 1, limit = 20) => {
    loading.value = true;
    error.value = null;

    try {
      const offset = (page - 1) * limit;

      const query = gql`
        query GetRecipeComments(
          $recipe_id: uuid!
          $limit: Int!
          $offset: Int!
        ) {
          comments(
            where: { recipe_id: { _eq: $recipe_id } }
            order_by: { created_at: desc }
            limit: $limit
            offset: $offset
          ) {
            id
            content
            created_at
            user {
              id
              username
            }
          }
        }
      `;

      const variables = {
        recipe_id: recipeId,
        limit: limit,
        offset: offset,
      };

      console.log("📡 Fetching comments with variables:", variables);

      const result = await $apollo.query({
        query,
        variables,
        fetchPolicy: "network-only",
      });

      const comments = result.data?.comments || [];
      console.log("✅ Comments fetched:", comments);

      return {
        comments: comments,
        total: comments.length, // Simple count instead of aggregate
      };
    } catch (err) {
      console.error("❌ Error fetching comments:", err);
      error.value = err.message;
      return { comments: [], total: 0 };
    } finally {
      loading.value = false;
    }
  };

  // Create a new comment
  const createComment = async (recipeId, content) => {
    loading.value = true;
    error.value = null;

    try {
      console.log("💬 Creating comment for recipe:", recipeId);
      console.log("👤 Current user:", user.value);

      if (!user.value?.id) {
        throw new Error("Please log in to comment");
      }

      const mutation = gql`
        mutation CreateComment($object: comments_insert_input!) {
          insert_comments_one(object: $object) {
            id
            content
            created_at
            user {
              id
              username
            }
          }
        }
      `;

      const variables = {
        object: {
          recipe_id: recipeId,
          content: content,
          user_id: user.value.id,
        },
      };

      console.log("📡 Sending comment mutation:", variables);

      const result = await $apollo.mutate({
        mutation,
        variables,
      });

      console.log("✅ Comment created:", result.data?.insert_comments_one);

      return result.data?.insert_comments_one;
    } catch (err) {
      console.error("❌ Error creating comment:", err);
      error.value = err.message;
      return null;
    } finally {
      loading.value = false;
    }
  };

  return {
    loading,
    error,
    getRecipeComments,
    createComment,
  };
};
