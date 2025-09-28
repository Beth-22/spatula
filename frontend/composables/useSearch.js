// composables/useSearch.js
import { gql } from "graphql-tag";

export const useSearch = () => {
  const { $apollo } = useNuxtApp();
  const loading = ref(false);
  const error = ref(null);

  // Search recipes by query
  const searchRecipes = async (query, filters = {}) => {
    loading.value = true;
    error.value = null;

    try {
      // Build the where clause dynamically to avoid null issues
      const whereConditions = [];

      // Search term condition
      if (query) {
        whereConditions.push(`
          _or: [
            { title: { _ilike: "%${query}%" } }
            { description: { _ilike: "%${query}%" } }
            { ingredients: { name: { _ilike: "%${query}%" } } }
            { user: { username: { _ilike: "%${query}%" } } }
          ]
        `);
      }

      // Category filter
      if (filters.category) {
        whereConditions.push(`category_id: { _eq: "${filters.category}" }`);
      }

      // Creator filter
      if (filters.creator) {
        whereConditions.push(`user_id: { _eq: "${filters.creator}" }`);
      }

      // Time range filter
      if (filters.timeRange) {
        const timeFilter = getTimeFilter(filters.timeRange);
        if (timeFilter) {
          whereConditions.push(timeFilter);
        }
      }

      // Ingredients filter
      if (filters.ingredients && filters.ingredients.length > 0) {
        const ingredientConditions = filters.ingredients
          .map(
            (ingredient) =>
              `{ ingredients: { name: { _ilike: "%${ingredient}%" } } }`
          )
          .join(", ");
        whereConditions.push(`_or: [${ingredientConditions}]`);
      }

      // Build the final where clause
      const whereClause =
        whereConditions.length > 0
          ? `where: { _and: [${whereConditions
              .map((cond) => `{ ${cond} }`)
              .join(", ")}] }`
          : "";

      const searchQuery = gql`
        query SearchRecipes {
          recipes(
            ${whereClause}
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
            ingredients {
              id
              name
              quantity
              unit
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
      `;

      console.log("Executing search query with filters:", filters);

      const result = await $apollo.query({
        query: searchQuery,
        fetchPolicy: "network-only",
      });

      return result.data?.recipes || [];
    } catch (err) {
      console.error("Error searching recipes:", err);
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  // Get all categories for filter dropdown (including empty ones)
  const getCategories = async () => {
    try {
      const query = gql`
        query GetCategories {
          categories {
            id
            name
          }
        }
      `;

      const result = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });

      return result.data?.categories || [];
    } catch (err) {
      console.error("Error getting categories:", err);
      return [];
    }
  };

  // Get only users who have created recipes
  const getCreators = async () => {
    try {
      const query = gql`
        query GetCreators {
          users(where: { recipes: { id: { _is_null: false } } }) {
            id
            username
          }
        }
      `;

      const result = await $apollo.query({
        query,
        fetchPolicy: "network-only",
      });

      return result.data?.users || [];
    } catch (err) {
      console.error("Error getting creators:", err);
      return [];
    }
  };

  // Get popular ingredients for suggestions
  const getPopularIngredients = () => {
    // Return the specified popular ingredients
    return [
      "Chicken",
      "Tomato",
      "Meat",
      "Eggs",
      "Salad",
      "Potato",
      "Flour",
      "Avocados",
    ];
  };

  // Helper function for time range conversion
  const getTimeFilter = (timeRange) => {
    switch (timeRange) {
      case "quick":
        return "prep_time: { _lte: 30 }";
      case "medium":
        return "prep_time: { _gte: 30, _lte: 60 }";
      case "long":
        return "prep_time: { _gte: 60, _lte: 120 }";
      case "allDay":
        return "prep_time: { _gte: 120 }";
      default:
        return null;
    }
  };

  return {
    loading,
    error,
    searchRecipes,
    getCategories,
    getCreators,
    getPopularIngredients,
  };
};
