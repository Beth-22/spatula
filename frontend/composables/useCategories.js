// composables/useCategories.js
import { gql } from "graphql-tag";

export const useCategories = () => {
  const getCategories = async () => {
    try {
      console.log("🔄 Fetching categories...");

      const { $apollo } = useNuxtApp();

      // Simple query without aggregates
      const GET_CATEGORIES = gql`
        query GetCategories {
          categories {
            id
            name
            image
          }
        }
      `;

      const { data } = await $apollo.query({
        query: GET_CATEGORIES,
        fetchPolicy: "network-only",
      });

      console.log("📦 Categories data:", data);

      if (!data || !data.categories) {
        console.log("No categories data found");
        return [];
      }

      // For now, we'll set recipeCount to 0 and fetch counts separately if needed
      const categories = data.categories.map((category) => ({
        id: category.id,
        name: category.name,
        image: category.image,
        recipeCount: 0, // Temporary - we'll fix this next
      }));

      console.log("📊 Processed categories:", categories);
      return categories;
    } catch (error) {
      console.error("❌ Error fetching categories:", error);
      return [];
    }
  };

  return { getCategories };
};
