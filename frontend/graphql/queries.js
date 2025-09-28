import { gql } from "@apollo/client/core";

export const GET_RECIPE_BY_ID = gql`
  query GetRecipeById($id: uuid!) {
    recipes_by_pk(id: $id) {
      id
      title
      description
      prep_time
      is_premium
      price
      user_id
      category_id
      created_at
      updated_at
      user {
        id
        username
        # Remove email field since it doesn't exist
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
