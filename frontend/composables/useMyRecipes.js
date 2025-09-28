// frontend/composables/useMyRecipes.js
import { gql } from "graphql-tag";

export const useMyRecipes = () => {
  const { $apollo } = useNuxtApp();
  const { user } = useAuth();

  const getMyRecipes = async () => {
    if (!user.value?.id) {
      console.error("No user ID available");
      return [];
    }

    const query = gql`
      query GetMyRecipes($user_id: uuid!) {
        recipes(
          where: { user_id: { _eq: $user_id } }
          order_by: { created_at: desc }
        ) {
          id
          title
          description
          prep_time
          is_premium
          price
          created_at
          updated_at
          category {
            id
            name
          }
          ingredients {
            name
            quantity
            unit
          }
          steps {
            step_number
            instruction
          }
          recipe_images(where: { is_featured: { _eq: true } }, limit: 1) {
            image_url
            is_featured
          }
        }
      }
    `;

    try {
      const { data } = await $apollo.query({
        query,
        variables: { user_id: user.value.id },
        fetchPolicy: "network-only",
      });
      console.log("My recipes data with images:", data.recipes);

      // Debug: Check if images are coming through
      data.recipes.forEach((recipe) => {
        console.log(`Recipe: ${recipe.title}, Images:`, recipe.recipe_images);
      });

      return data.recipes || [];
    } catch (error) {
      console.error("Error fetching my recipes:", error);
      return [];
    }
  };

  const deleteRecipe = async (recipeId) => {
    const mutation = gql`
      mutation DeleteRecipe($id: uuid!) {
        delete_recipes_by_pk(id: $id) {
          id
        }
      }
    `;

    try {
      const { data } = await $apollo.mutate({
        mutation,
        variables: { id: recipeId },
      });
      return data.delete_recipes_by_pk;
    } catch (error) {
      console.error("Error deleting recipe:", error);
      throw error;
    }
  };

  return {
    getMyRecipes,
    deleteRecipe,
  };
};
