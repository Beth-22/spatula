// backend/internal/actions/recipes.go
package actions

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Beth-22/spatula/internal/auth"
	"github.com/Beth-22/spatula/internal/models"
	"github.com/Beth-22/spatula/internal/repository"
	"github.com/Beth-22/spatula/internal/utils"
)

// CreateRecipeHandler - Creates a new recipe with ingredients, steps, and ALL images
func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Authentication required")
		return
	}

	var req models.CreateRecipeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if strings.TrimSpace(req.Title) == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Recipe title is required")
		return
	}
	if strings.TrimSpace(req.Description) == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Recipe description is required")
		return
	}
	if req.CategoryID == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Category is required")
		return
	}
	if req.PrepTime <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Prep time must be greater than 0")
		return
	}

	// Create recipe in database
	recipeID, err := repository.CreateRecipe(userID, &req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create recipe: "+err.Error())
		return
	}

	// NEW: Save ALL images to recipe_images table
	if len(req.AllImages) > 0 {
		successfulImageCount := 0
		
		for _, imageUrl := range req.AllImages {
			if strings.TrimSpace(imageUrl) == "" {
				continue // Skip empty URLs
			}
			
			// Determine if this is the featured image
			isFeatured := (imageUrl == req.FeaturedImage)
			
			err := repository.CreateRecipeImage(recipeID, imageUrl, isFeatured)
			if err != nil {
				// Log the error but continue with other images
				utils.LogError("Failed to create recipe image for URL " + imageUrl + ": " + err.Error())
			} else {
				successfulImageCount++
			}
		}
		
		utils.LogInfo("Successfully saved " + string(successfulImageCount) + " images for recipe " + recipeID)
		
		// If no images were saved successfully, log a warning but don't fail the recipe creation
		if successfulImageCount == 0 {
			utils.LogWarning("No images were saved for recipe " + recipeID)
		}
	} else {
		utils.LogWarning("No images provided for recipe " + recipeID)
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message":   "Recipe created successfully",
		"recipe_id": recipeID,
		"status":    "success",
	})
}

// GetRecipeHandler - Get single recipe by ID
func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.Query().Get("id")
	if recipeID == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Recipe ID is required")
		return
	}

	recipe, err := repository.GetRecipeByID(recipeID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Recipe not found: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, recipe)
}

// GetMyRecipesHandler - Get authenticated user's recipes
func GetMyRecipesHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Authentication required")
		return
	}

	recipes, err := repository.GetUserRecipes(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch recipes: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, recipes)
}