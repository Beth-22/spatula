// composables/useRecipeApi.js
import { gql } from "graphql-tag";

export const useRecipeApi = () => {
  const { $apollo } = useNuxtApp();

  // Get single recipe by ID
  const getRecipeById = async (recipeId) => {
    const query = gql`
      query GetRecipe($id: uuid!) {
        recipes_by_pk(id: $id) {
          id
          title
          description
          prep_time
          is_premium
          price
          created_at
          updated_at
          user {
            id
            username
          }
          category {
            id
            name
          }
          ingredients {
            id
            name
            quantity
            unit
          }
          steps(order_by: { step_number: asc }) {
            id
            step_number
            instruction
          }
          recipe_images(order_by: { is_featured: desc, created_at: asc }) {
            id
            image_url
            is_featured
            created_at
          }
        }
      }
    `;

    try {
      const { data } = await $apollo.query({
        query,
        variables: { id: recipeId },
        fetchPolicy: "network-only",
      });

      console.log("Recipe data loaded:", data.recipes_by_pk);
      return data.recipes_by_pk;
    } catch (error) {
      console.error("Error fetching recipe:", error);
      return null;
    }
  };

  const getCategories = async () => {
    const query = gql`
      query GetCategories {
        categories {
          id
          name
          image
        }
      }
    `;

    try {
      const { data } = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });
      return data.categories || [];
    } catch (error) {
      console.error("Error fetching categories:", error);
      return [];
    }
  };

  // Get recipes by category name
  const getRecipesByCategory = async (categoryName) => {
    try {
      const query = gql`
        query GetRecipesByCategory($categoryName: String!) {
          recipes(where: { category: { name: { _eq: $categoryName } } }) {
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
            recipe_images(order_by: { is_featured: desc, created_at: asc }) {
              id
              image_url
              is_featured
              created_at
            }
          }
        }
      `;

      const { data } = await $apollo.query({
        query,
        variables: { categoryName },
        fetchPolicy: "network-only",
      });

      console.log("Recipes by category loaded:", data.recipes);
      return data.recipes || [];
    } catch (error) {
      console.error("Error fetching recipes by category:", error);
      return [];
    }
  };

  // Get all recipes
  const getAllRecipes = async () => {
    try {
      const query = gql`
        query GetAllRecipes {
          recipes(order_by: { created_at: desc }) {
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
            recipe_images(order_by: { is_featured: desc, created_at: asc }) {
              id
              image_url
              is_featured
              created_at
            }
          }
        }
      `;

      const { data } = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });

      console.log("All recipes loaded:", data.recipes);
      return data.recipes || [];
    } catch (error) {
      console.error("Error fetching all recipes:", error);
      return [];
    }
  };

  // Search recipes by keyword
  const searchRecipes = async (searchTerm, filters = {}) => {
    try {
      console.log(
        "Searching recipes with term:",
        searchTerm,
        "filters:",
        filters
      );

      let whereConditions = [];
      let variables = {};

      // Base search condition
      if (searchTerm && searchTerm.trim()) {
        whereConditions.push(`{
          _or: [
            {title: {_ilike: $searchTerm}},
            {description: {_ilike: $searchTerm}},
            {ingredients: {name: {_ilike: $searchTerm}}}
          ]
        }`);
        variables.searchTerm = `%${searchTerm}%`;
      }

      // Category filter
      if (filters.category) {
        whereConditions.push(`{category_id: {_eq: $category}}`);
        variables.category = filters.category;
      }

      // Premium filter
      if (filters.premium !== undefined) {
        whereConditions.push(`{is_premium: {_eq: $premium}}`);
        variables.premium = filters.premium;
      }

      // Time filter
      if (filters.timeRange) {
        switch (filters.timeRange) {
          case "quick":
            whereConditions.push(`{prep_time: {_lt: 30}}`);
            break;
          case "medium":
            whereConditions.push(`{prep_time: {_gte: 30, _lte: 60}}`);
            break;
          case "long":
            whereConditions.push(`{prep_time: {_gte: 60, _lte: 120}}`);
            break;
          case "allDay":
            whereConditions.push(`{prep_time: {_gt: 120}}`);
            break;
        }
      }

      // Build the where clause
      let whereClause = "";
      if (whereConditions.length > 0) {
        whereClause = `where: {_and: [${whereConditions.join(", ")}]}`;
      }

      const query = gql`
        query SearchRecipes($searchTerm: String, $category: uuid, $premium: Boolean) {
          recipes(order_by: {created_at: desc}, ${whereClause}) {
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
            recipe_images(order_by: {is_featured: desc, created_at: asc}) {
              id
              image_url
              is_featured
              created_at
            }
          }
        }
      `;

      const { data } = await $apollo.query({
        query,
        variables,
        fetchPolicy: "network-only",
      });

      console.log("Search results:", data.recipes);
      return data.recipes || [];
    } catch (error) {
      console.error("Error searching recipes:", error);
      return [];
    }
  };

  // Get popular ingredients for search suggestions
  const getPopularIngredients = async () => {
    try {
      const query = gql`
        query GetPopularIngredients {
          ingredients(
            limit: 20
            order_by: { recipes_aggregate: { count: desc } }
          ) {
            name
          }
        }
      `;

      const { data } = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });

      return data.ingredients.map((ing) => ing.name) || [];
    } catch (error) {
      console.error("Error fetching popular ingredients:", error);
      return [
        "chicken",
        "beef",
        "pasta",
        "rice",
        "tomato",
        "cheese",
        "garlic",
        "onion",
      ];
    }
  };

  const getFeaturedRecipes = async () => {
    try {
      const query = gql`
        query GetFeaturedRecipes {
          recipes(
            where: { is_premium: { _eq: true } }
            limit: 6
            order_by: { created_at: desc }
          ) {
            id
            title
            description
            prep_time
            is_premium
            price
            created_at
            user {
              username
            }
            category {
              name
            }
            recipe_images(order_by: { is_featured: desc, created_at: asc }) {
              image_url
            }
          }
        }
      `;

      const { data } = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });

      return data.recipes || [];
    } catch (error) {
      console.error("Error fetching featured recipes:", error);
      return [];
    }
  };

  return {
    getRecipeById,
    getCategories,
    getRecipesByCategory,
    getAllRecipes,
    searchRecipes,
    getPopularIngredients,
    getFeaturedRecipes,
  };
};
