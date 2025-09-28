package repository

import (
	"context"
	"fmt"
	"log"
	//"strings"

	"github.com/Beth-22/spatula/internal/graphql"
	"github.com/Beth-22/spatula/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// User operations (existing - keep these)
func CreateUser(email, username, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	query := `
	mutation CreateUser($email: String!, $username: String!, $password_hash: String!) {
		insert_users_one(object: {email: $email, username: $username, password_hash: $password_hash}) {
			id
		}
	}`

	variables := map[string]interface{}{
		"email":         email,
		"username":      username,
		"password_hash": string(hashedPassword),
	}

	resp, err := graphql.Mutate(context.Background(), query, variables)
	if err != nil {
		return "", fmt.Errorf("failed to create user via GraphQL: %w", err)
	}

	// Extract the response data safely
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response format from GraphQL")
	}

	insertData, ok := data["insert_users_one"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to extract user data from response")
	}

	userID, ok := insertData["id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get user ID from response")
	}

	log.Printf("Successfully created user: %s", userID)
	return userID, nil
}

func VerifyUser(email, password string) (string, error) {
	query := `
	query GetUser($email: String!) {
		users(where: {email: {_eq: $email}}) {
			id
			password_hash
		}
	}`

	variables := map[string]interface{}{"email": email}
	resp, err := graphql.Query(context.Background(), query, variables)
	if err != nil {
		return "", fmt.Errorf("failed to query user: %w", err)
	}

	// Extract the response data safely
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response format from GraphQL")
	}

	usersInterface, ok := data["users"].([]interface{})
	if !ok {
		return "", fmt.Errorf("failed to extract users array from response")
	}

	if len(usersInterface) == 0 {
		return "", fmt.Errorf("user not found with email: %s", email)
	}

	user, ok := usersInterface[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to extract user object from response")
	}

	userID, ok := user["id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get user ID from response")
	}

	storedHash, ok := user["password_hash"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get password hash from response")
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	log.Printf("Successfully authenticated user: %s", userID)
	return userID, nil
}

// RECIPE OPERATIONS

func CreateRecipe(userID string, req *models.CreateRecipeRequest) (string, error) {
	// 1. First create the recipe
	recipeQuery := `
	mutation CreateRecipe($object: recipes_insert_input!) {
		insert_recipes_one(object: $object) {
			id
			title
		}
	}`

	recipeVariables := map[string]interface{}{
		"object": map[string]interface{}{
			"title":          req.Title,
			"description":    req.Description,
			"prep_time":      req.PrepTime,
			"featured_image": req.FeaturedImage,
			"is_premium":     req.IsPremium,
			"price":          req.Price,
			"user_id":        userID,
			"category_id":    req.CategoryID,
		},
	}

	resp, err := graphql.Mutate(context.Background(), recipeQuery, recipeVariables)
	if err != nil {
		return "", fmt.Errorf("failed to create recipe: %w", err)
	}

	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response format")
	}

	recipeData, ok := data["insert_recipes_one"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to extract recipe data")
	}

	recipeID, ok := recipeData["id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get recipe ID")
	}

	// 2. Add ingredients if provided
	if len(req.Ingredients) > 0 {
		ingredientObjects := make([]map[string]interface{}, len(req.Ingredients))
		for i, ing := range req.Ingredients {
			ingredientObjects[i] = map[string]interface{}{
				"recipe_id": recipeID,
				"name":      ing.Name,
				"quantity":  ing.Quantity,
				"unit":      ing.Unit,
			}
		}

		ingredientQuery := `
		mutation AddIngredients($objects: [ingredients_insert_input!]!) {
			insert_ingredients(objects: $objects) {
				affected_rows
			}
		}`

		_, err = graphql.Mutate(context.Background(), ingredientQuery, map[string]interface{}{
			"objects": ingredientObjects,
		})
		if err != nil {
			return "", fmt.Errorf("failed to add ingredients: %w", err)
		}
	}

	// 3. Add steps if provided
	if len(req.Steps) > 0 {
		stepObjects := make([]map[string]interface{}, len(req.Steps))
		for i, step := range req.Steps {
			stepObjects[i] = map[string]interface{}{
				"recipe_id":   recipeID,
				"step_number": step.StepNumber,
				"instruction": step.Instruction,
			}
		}

		stepQuery := `
		mutation AddSteps($objects: [steps_insert_input!]!) {
			insert_steps(objects: $objects) {
				affected_rows
			}
		}`

		_, err = graphql.Mutate(context.Background(), stepQuery, map[string]interface{}{
			"objects": stepObjects,
		})
		if err != nil {
			return "", fmt.Errorf("failed to add steps: %w", err)
		}
	}

	log.Printf("Successfully created recipe: %s", recipeID)
	return recipeID, nil
}
// Add this to your backend/internal/repository/hasura_client.go file

// backend/internal/repository/hasura_client.go
func CreateRecipeImage(recipeID string, imageURL string, isFeatured bool) error {
    query := `
    mutation InsertRecipeImage($object: recipe_images_insert_input!) {
        insert_recipe_images_one(object: $object) {
            id
        }
    }`

    variables := map[string]interface{}{
        "object": map[string]interface{}{
            "recipe_id":   recipeID,
            "image_url":   imageURL,
            "is_featured": isFeatured,
        },
    }

    // Use the existing graphql.Mutate function
    _, err := graphql.Mutate(context.Background(), query, variables)
    if err != nil {
        return fmt.Errorf("failed to create recipe image: %w", err)
    }

    log.Printf("Successfully created recipe image for recipe: %s", recipeID)
    return nil
}


func GetRecipeByID(recipeID string) (*models.Recipe, error) {
    query := `
    query GetRecipe($id: uuid!) {
        recipes_by_pk(id: $id) {
            id
            title
            description
            prep_time
            featured_image
            is_premium
            price
            user_id
            category_id
            created_at
            updated_at
            user {
                id
                username
                email
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
            steps(order_by: {step_number: asc}) {
                id
                step_number
                instruction
            }
            likes_aggregate {
                aggregate {
                    count
                }
            }
            ratings_aggregate {
                aggregate {
                    avg {
                        rating
                    }
                    count
                }
            }
            recipe_images(order_by: {is_featured: desc, created_at: asc}) {
                id
                image_url
                is_featured
                created_at
            }
        }
    }`

    variables := map[string]interface{}{"id": recipeID}
    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return nil, fmt.Errorf("failed to query recipe: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid response format")
    }

    recipeData, ok := data["recipes_by_pk"].(map[string]interface{})
    if !ok || recipeData == nil {
        return nil, fmt.Errorf("recipe not found")
    }

    return mapToRecipe(recipeData), nil
}

func GetUserRecipes(userID string) ([]*models.Recipe, error) {
	query := `
	query GetUserRecipes($user_id: uuid!) {
		recipes(where: {user_id: {_eq: $user_id}}, order_by: {created_at: desc}) {
			id
			title
			description
			prep_time
			featured_image
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
			steps_aggregate {
				aggregate {
					count
				}
			}
			likes_aggregate {
				aggregate {
					count
				}
			}
			ratings_aggregate {
				aggregate {
					avg {
						rating
					}
					count
				}
			}
		}
	}`

	variables := map[string]interface{}{"user_id": userID}
	resp, err := graphql.Query(context.Background(), query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query user recipes: %w", err)
	}

	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	recipesInterface, ok := data["recipes"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to extract recipes array")
	}

	var recipes []*models.Recipe
	for _, recipeInterface := range recipesInterface {
		if recipeMap, ok := recipeInterface.(map[string]interface{}); ok {
			recipes = append(recipes, mapToRecipe(recipeMap))
		}
	}

	return recipes, nil
}

// Helper function to map GraphQL response to Recipe model
func mapToRecipe(data map[string]interface{}) *models.Recipe {
	recipe := &models.Recipe{
		ID:            getString(data, "id"),
		Title:         getString(data, "title"),
		Description:   getString(data, "description"),
		PrepTime:      getInt(data, "prep_time"),
		FeaturedImage: getString(data, "featured_image"),
		IsPremium:     getBool(data, "is_premium"),
		UserID:        getString(data, "user_id"),
		CategoryID:    getString(data, "category_id"),
	}

	if price, ok := data["price"].(float64); ok {
		recipe.Price = &price
	}

	// Map user if present
	if userData, ok := data["user"].(map[string]interface{}); ok {
		recipe.User = &models.User{
			ID:       getString(userData, "id"),
			Username: getString(userData, "username"),
			Email:    getString(userData, "email"),
		}
	}

	// Map category if present
	if categoryData, ok := data["category"].(map[string]interface{}); ok {
		recipe.Category = &models.Category{
			ID:   getString(categoryData, "id"),
			Name: getString(categoryData, "name"),
		}
	}
	// Map recipe images if present
if imagesInterface, ok := data["recipe_images"].([]interface{}); ok {
    for _, imageInterface := range imagesInterface {
        if imageMap, ok := imageInterface.(map[string]interface{}); ok {
            recipe.RecipeImages = append(recipe.RecipeImages, models.RecipeImage{
                ID:        getString(imageMap, "id"),
                RecipeID:  recipe.ID,
                ImageURL:  getString(imageMap, "image_url"),
                IsFeatured: getBool(imageMap, "is_featured"),
                CreatedAt: parseTime(getString(imageMap, "created_at")),
            })
        }
    }
}

	// Map ingredients if present
	if ingredientsInterface, ok := data["ingredients"].([]interface{}); ok {
		for _, ingInterface := range ingredientsInterface {
			if ingMap, ok := ingInterface.(map[string]interface{}); ok {
				recipe.Ingredients = append(recipe.Ingredients, models.Ingredient{
					ID:       getString(ingMap, "id"),
					RecipeID: recipe.ID,
					Name:     getString(ingMap, "name"),
					Quantity: getString(ingMap, "quantity"),
					Unit:     getString(ingMap, "unit"),
				})
			}
		}
	}

	// Map steps if present
	if stepsInterface, ok := data["steps"].([]interface{}); ok {
		for _, stepInterface := range stepsInterface {
			if stepMap, ok := stepInterface.(map[string]interface{}); ok {
				recipe.Steps = append(recipe.Steps, models.Step{
					ID:         getString(stepMap, "id"),
					RecipeID:   recipe.ID,
					StepNumber: getInt(stepMap, "step_number"),
					Instruction: getString(stepMap, "instruction"),
				})
			}
		}
	}

	return recipe
}

// Helper functions for safe type conversion
func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getInt(data map[string]interface{}, key string) int {
	if val, ok := data[key]; ok && val != nil {
		if num, ok := val.(float64); ok {
			return int(num)
		}
	}
	return 0
}

func getBool(data map[string]interface{}, key string) bool {
	if val, ok := data[key]; ok && val != nil {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}

// New function to get recipe statistics for a user
func GetUserRecipeStats(userID string) (map[string]interface{}, error) {
	query := `
	query GetUserRecipeStats($user_id: uuid!) {
		recipes_aggregate(where: {user_id: {_eq: $user_id}}) {
			aggregate {
				count
			}
		}
		premium_recipes: recipes_aggregate(where: {user_id: {_eq: $user_id}, is_premium: {_eq: true}}) {
			aggregate {
				count
			}
		}
		total_likes: likes_aggregate(where: {recipe: {user_id: {_eq: $user_id}}}) {
			aggregate {
				count
			}
		}
	}`

	variables := map[string]interface{}{"user_id": userID}
	resp, err := graphql.Query(context.Background(), query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query user stats: %w", err)
	}

	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	stats := make(map[string]interface{})
	
	// Extract total recipes count
	if recipesAgg, ok := data["recipes_aggregate"].(map[string]interface{}); ok {
		if aggregate, ok := recipesAgg["aggregate"].(map[string]interface{}); ok {
			if count, ok := aggregate["count"].(float64); ok {
				stats["total"] = int(count)
			}
		}
	}
	
	// Extract premium recipes count
	if premiumAgg, ok := data["premium_recipes"].(map[string]interface{}); ok {
		if aggregate, ok := premiumAgg["aggregate"].(map[string]interface{}); ok {
			if count, ok := aggregate["count"].(float64); ok {
				stats["premium"] = int(count)
			}
		}
	}
	
	// Extract total likes count
	if likesAgg, ok := data["total_likes"].(map[string]interface{}); ok {
		if aggregate, ok := likesAgg["aggregate"].(map[string]interface{}); ok {
			if count, ok := aggregate["count"].(float64); ok {
				stats["likes"] = int(count)
			}
		}
	}

	return stats, nil
}

// Comment operations
func CreateComment(userID, recipeID, content string) (string, error) {
    query := `
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
    }`

    variables := map[string]interface{}{
        "object": map[string]interface{}{
            "user_id":   userID,
            "recipe_id": recipeID,
            "content":   content,
        },
    }

    resp, err := graphql.Mutate(context.Background(), query, variables)
    if err != nil {
        return "", fmt.Errorf("failed to create comment: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("invalid response format")
    }

    commentData, ok := data["insert_comments_one"].(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("failed to extract comment data")
    }

    commentID, ok := commentData["id"].(string)
    if !ok {
        return "", fmt.Errorf("failed to get comment ID")
    }

    log.Printf("Successfully created comment: %s", commentID)
    return commentID, nil
}

func GetRecipeComments(recipeID string, limit, offset int) ([]models.CommentResponse, int, error) {
    query := `
    query GetRecipeComments($recipe_id: uuid!, $limit: Int!, $offset: Int!) {
        comments(
            where: {recipe_id: {_eq: $recipe_id}}
            order_by: {created_at: desc}
            limit: $limit
            offset: $offset
        ) {
            id
            content
            created_at
            updated_at
            user {
                id
                username
                email
            }
        }
        comments_aggregate(where: {recipe_id: {_eq: $recipe_id}}) {
            aggregate {
                count
            }
        }
    }`

    variables := map[string]interface{}{
        "recipe_id": recipeID,
        "limit":     limit,
        "offset":    offset,
    }

    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to query comments: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return nil, 0, fmt.Errorf("invalid response format")
    }

    // Extract comments
    commentsInterface, ok := data["comments"].([]interface{})
    if !ok {
        return nil, 0, fmt.Errorf("failed to extract comments array")
    }

    var comments []models.CommentResponse
    for _, commentInterface := range commentsInterface {
        if commentMap, ok := commentInterface.(map[string]interface{}); ok {
            comments = append(comments, mapToCommentResponse(commentMap))
        }
    }

    // Extract total count
    total := 0
    if aggregateInterface, ok := data["comments_aggregate"].(map[string]interface{}); ok {
        if aggregate, ok := aggregateInterface["aggregate"].(map[string]interface{}); ok {
            if count, ok := aggregate["count"].(float64); ok {
                total = int(count)
            }
        }
    }

    return comments, total, nil
}

func DeleteComment(commentID, userID string) error {
    query := `
    mutation DeleteComment($comment_id: uuid!, $user_id: uuid!) {
        delete_comments(where: {id: {_eq: $comment_id}, user_id: {_eq: $user_id}}) {
            affected_rows
        }
    }`

    variables := map[string]interface{}{
        "comment_id": commentID,
        "user_id":    userID,
    }

    resp, err := graphql.Mutate(context.Background(), query, variables)
    if err != nil {
        return fmt.Errorf("failed to delete comment: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return fmt.Errorf("invalid response format")
    }

    deleteData, ok := data["delete_comments"].(map[string]interface{})
    if !ok {
        return fmt.Errorf("failed to extract delete response")
    }

    affectedRows, ok := deleteData["affected_rows"].(float64)
    if !ok || affectedRows == 0 {
        return fmt.Errorf("comment not found or not authorized")
    }

    log.Printf("Successfully deleted comment: %s", commentID)
    return nil
}

// Helper function to map GraphQL response to CommentResponse
func mapToCommentResponse(data map[string]interface{}) models.CommentResponse {
    comment := models.CommentResponse{
        ID:        getString(data, "id"),
        Content:   getString(data, "content"),
        CreatedAt: parseTime(getString(data, "created_at")),
        UpdatedAt: parseTime(getString(data, "updated_at")),
    }

    // Map user if present
    if userData, ok := data["user"].(map[string]interface{}); ok {
        comment.User = models.User{
            ID:       getString(userData, "id"),
            Username: getString(userData, "username"),
            Email:    getString(userData, "email"),
        }
    }

    return comment
}

// Like operations
func ToggleLike(userID, recipeID string) (bool, int, error) {
    // First check if user already liked the recipe
    checkQuery := `
    query CheckLike($user_id: uuid!, $recipe_id: uuid!) {
        likes(where: {user_id: {_eq: $user_id}, recipe_id: {_eq: $recipe_id}}) {
            id
        }
        likes_aggregate(where: {recipe_id: {_eq: $recipe_id}}) {
            aggregate {
                count
            }
        }
    }`

    checkVariables := map[string]interface{}{
        "user_id":   userID,
        "recipe_id": recipeID,
    }

    checkResp, err := graphql.Query(context.Background(), checkQuery, checkVariables)
    if err != nil {
        return false, 0, fmt.Errorf("failed to check like: %w", err)
    }

    checkData, ok := checkResp.Data.(map[string]interface{})
    if !ok {
        return false, 0, fmt.Errorf("invalid response format")
    }

    // Check if like exists
    likesInterface, ok := checkData["likes"].([]interface{})
    if !ok {
        return false, 0, fmt.Errorf("failed to extract likes array")
    }

    alreadyLiked := len(likesInterface) > 0

    if alreadyLiked {
        // Unlike the recipe
        unlikeQuery := `
        mutation UnlikeRecipe($user_id: uuid!, $recipe_id: uuid!) {
            delete_likes(where: {user_id: {_eq: $user_id}, recipe_id: {_eq: $recipe_id}}) {
                affected_rows
            }
        }`

        _, err := graphql.Mutate(context.Background(), unlikeQuery, checkVariables)
        if err != nil {
            return false, 0, fmt.Errorf("failed to unlike recipe: %w", err)
        }
    } else {
        // Like the recipe
        likeQuery := `
        mutation LikeRecipe($object: likes_insert_input!) {
            insert_likes_one(object: $object) {
                id
            }
        }`

        likeVariables := map[string]interface{}{
            "object": map[string]interface{}{
                "user_id":   userID,
                "recipe_id": recipeID,
            },
        }

        _, err := graphql.Mutate(context.Background(), likeQuery, likeVariables)
        if err != nil {
            return false, 0, fmt.Errorf("failed to like recipe: %w", err)
        }
    }

    // Get updated like count
    countAggregate, ok := checkData["likes_aggregate"].(map[string]interface{})
    if ok {
        if aggregate, ok := countAggregate["aggregate"].(map[string]interface{}); ok {
            if count, ok := aggregate["count"].(float64); ok {
                return !alreadyLiked, int(count), nil
            }
        }
    }

    return !alreadyLiked, 0, nil
}

func GetUserLike(userID, recipeID string) (bool, error) {
    query := `
    query GetUserLike($user_id: uuid!, $recipe_id: uuid!) {
        likes(where: {user_id: {_eq: $user_id}, recipe_id: {_eq: $recipe_id}}) {
            id
        }
    }`

    variables := map[string]interface{}{
        "user_id":   userID,
        "recipe_id": recipeID,
    }

    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return false, fmt.Errorf("failed to get user like: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return false, fmt.Errorf("invalid response format")
    }

    likesInterface, ok := data["likes"].([]interface{})
    if !ok {
        return false, fmt.Errorf("failed to extract likes array")
    }

    return len(likesInterface) > 0, nil
}

// Rating operations
func SetRating(userID, recipeID string, rating int) (float64, int, error) {
    // First check if user already rated the recipe
    checkQuery := `
    query CheckRating($user_id: uuid!, $recipe_id: uuid!) {
        ratings(where: {user_id: {_eq: $user_id}, recipe_id: {_eq: $recipe_id}}) {
            id
            rating
        }
        ratings_aggregate(where: {recipe_id: {_eq: $recipe_id}}) {
            aggregate {
                avg {
                    rating
                }
                count
            }
        }
    }`

    checkVariables := map[string]interface{}{
        "user_id":   userID,
        "recipe_id": recipeID,
    }

    checkResp, err := graphql.Query(context.Background(), checkQuery, checkVariables)
    if err != nil {
        return 0, 0, fmt.Errorf("failed to check rating: %w", err)
    }

    checkData, ok := checkResp.Data.(map[string]interface{})
    if !ok {
        return 0, 0, fmt.Errorf("invalid response format")
    }

    // Check if rating exists
    ratingsInterface, ok := checkData["ratings"].([]interface{})
    if !ok {
        return 0, 0, fmt.Errorf("failed to extract ratings array")
    }

    if len(ratingsInterface) > 0 {
        // Update existing rating
        ratingData := ratingsInterface[0].(map[string]interface{})
        ratingID := ratingData["id"].(string)

        updateQuery := `
        mutation UpdateRating($rating_id: uuid!, $rating: numeric!) {
            update_ratings_by_pk(pk_columns: {id: $rating_id}, _set: {rating: $rating}) {
                id
            }
        }`

        updateVariables := map[string]interface{}{
            "rating_id": ratingID,
            "rating":    rating,
        }

        _, err := graphql.Mutate(context.Background(), updateQuery, updateVariables)
        if err != nil {
            return 0, 0, fmt.Errorf("failed to update rating: %w", err)
        }
    } else {
        // Create new rating
        insertQuery := `
        mutation InsertRating($object: ratings_insert_input!) {
            insert_ratings_one(object: $object) {
                id
            }
        }`

        insertVariables := map[string]interface{}{
            "object": map[string]interface{}{
                "user_id":   userID,
                "recipe_id": recipeID,
                "rating":    rating,
            },
        }

        _, err := graphql.Mutate(context.Background(), insertQuery, insertVariables)
        if err != nil {
            return 0, 0, fmt.Errorf("failed to insert rating: %w", err)
        }
    }

    // Get updated average and count
    aggregate, ok := checkData["ratings_aggregate"].(map[string]interface{})
    if ok {
        if agg, ok := aggregate["aggregate"].(map[string]interface{}); ok {
            avgRating := 0.0
            ratingCount := 0

            if avgData, ok := agg["avg"].(map[string]interface{}); ok {
                if avg, ok := avgData["rating"].(float64); ok {
                    avgRating = avg
                }
            }

            if count, ok := agg["count"].(float64); ok {
                ratingCount = int(count)
            }

            return avgRating, ratingCount, nil
        }
    }

    return 0, 0, nil
}

func GetUserRating(userID, recipeID string) (int, error) {
    query := `
    query GetUserRating($user_id: uuid!, $recipe_id: uuid!) {
        ratings(where: {user_id: {_eq: $user_id}, recipe_id: {_eq: $recipe_id}}) {
            rating
        }
    }`

    variables := map[string]interface{}{
        "user_id":   userID,
        "recipe_id": recipeID,
    }

    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return 0, fmt.Errorf("failed to get user rating: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return 0, fmt.Errorf("invalid response format")
    }

    ratingsInterface, ok := data["ratings"].([]interface{})
    if !ok {
        return 0, fmt.Errorf("failed to extract ratings array")
    }

    if len(ratingsInterface) > 0 {
        ratingData := ratingsInterface[0].(map[string]interface{})
        if rating, ok := ratingData["rating"].(float64); ok {
            return int(rating), nil
        }
    }

    return 0, nil
}
// Purchase operations
func CreatePurchase(purchase *models.Purchase) (string, error) {
    query := `
    mutation CreatePurchase($object: purchases_insert_input!) {
        insert_purchases_one(object: $object) {
            id
            status
            created_at
        }
    }`

    variables := map[string]interface{}{
        "object": map[string]interface{}{
            "user_id":     purchase.UserID,
            "recipe_id":   purchase.RecipeID,
            "amount":      purchase.Amount,
            "status":      purchase.Status,
            "chapa_transaction_id": purchase.ChapaTransactionID,
        },
    }

    resp, err := graphql.Mutate(context.Background(), query, variables)
    if err != nil {
        return "", fmt.Errorf("failed to create purchase: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("invalid response format")
    }

    purchaseData, ok := data["insert_purchases_one"].(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("failed to extract purchase data")
    }

    purchaseID, ok := purchaseData["id"].(string)
    if !ok {
        return "", fmt.Errorf("failed to get purchase ID")
    }

    log.Printf("Successfully created purchase: %s", purchaseID)
    return purchaseID, nil
}

func UpdatePurchaseStatus(purchaseID, status string, chapaTransactionID *string) error {
    query := `
    mutation UpdatePurchase($purchase_id: uuid!, $status: String!, $chapa_transaction_id: String) {
        update_purchases_by_pk(
            pk_columns: {id: $purchase_id}, 
            _set: {status: $status, chapa_transaction_id: $chapa_transaction_id}
        ) {
            id
            status
        }
    }`

    variables := map[string]interface{}{
        "purchase_id": purchaseID,
        "status": status,
    }
    
    if chapaTransactionID != nil {
        variables["chapa_transaction_id"] = *chapaTransactionID
    }

    _, err := graphql.Mutate(context.Background(), query, variables)
    if err != nil {
        return fmt.Errorf("failed to update purchase status: %w", err)
    }

    log.Printf("Successfully updated purchase %s to status: %s", purchaseID, status)
    return nil
}

func HasUserPurchasedRecipe(userID, recipeID string) (bool, error) {
    query := `
    query CheckPurchase($user_id: uuid!, $recipe_id: uuid!) {
        purchases(
            where: {
                user_id: {_eq: $user_id}, 
                recipe_id: {_eq: $recipe_id},
                status: {_eq: "completed"}
            },
            limit: 1
        ) {
            id
        }
    }`

    variables := map[string]interface{}{
        "user_id": userID,
        "recipe_id": recipeID,
    }

    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return false, fmt.Errorf("failed to check purchase: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return false, fmt.Errorf("invalid response format")
    }

    purchasesInterface, ok := data["purchases"].([]interface{})
    if !ok {
        return false, fmt.Errorf("failed to extract purchases array")
    }

    return len(purchasesInterface) > 0, nil
}

func GetPurchaseByID(purchaseID string) (*models.Purchase, error) {
    query := `
    query GetPurchase($id: uuid!) {
        purchases_by_pk(id: $id) {
            id
            user_id
            recipe_id
            amount
            status
            chapa_transaction_id
            created_at
            updated_at
        }
    }`

    variables := map[string]interface{}{"id": purchaseID}
    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return nil, fmt.Errorf("failed to query purchase: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid response format")
    }

    purchaseData, ok := data["purchases_by_pk"].(map[string]interface{})
    if !ok || purchaseData == nil {
        return nil, fmt.Errorf("purchase not found")
    }

    purchase := &models.Purchase{
        ID:        getString(purchaseData, "id"),
        UserID:    getString(purchaseData, "user_id"),
        RecipeID:  getString(purchaseData, "recipe_id"),
        Status:    getString(purchaseData, "status"),
        CreatedAt: parseTime(getString(purchaseData, "created_at")),
        UpdatedAt: parseTime(getString(purchaseData, "updated_at")),
    }

    if amount, ok := purchaseData["amount"].(float64); ok {
        purchase.Amount = amount
    }

    if chapaID, ok := purchaseData["chapa_transaction_id"].(string); ok && chapaID != "" {
        purchase.ChapaTransactionID = &chapaID
    }

    return purchase, nil
}