package actions

import (
    "encoding/json"
    "net/http"

    "github.com/Beth-22/spatula/internal/auth"
    "github.com/Beth-22/spatula/internal/models"
    "github.com/Beth-22/spatula/internal/repository"
)

func SetRatingHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Authentication required", http.StatusUnauthorized)
        return
    }

    var req models.RatingRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if req.Rating < 1 || req.Rating > 5 {
        http.Error(w, "Rating must be between 1 and 5", http.StatusBadRequest)
        return
    }

    avgRating, ratingCount, err := repository.SetRating(userID, req.RecipeID, req.Rating)
    if err != nil {
        http.Error(w, "Failed to set rating: "+err.Error(), http.StatusInternalServerError)
        return
    }

    response := models.RatingResponse{
        Rating:       req.Rating,
        AverageRating: avgRating,
        RatingCount:  ratingCount,
        UserRating:   req.Rating,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func GetUserRatingHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Authentication required", http.StatusUnauthorized)
        return
    }

    recipeID := r.URL.Query().Get("recipe_id")
    if recipeID == "" {
        http.Error(w, "Recipe ID is required", http.StatusBadRequest)
        return
    }

    rating, err := repository.GetUserRating(userID, recipeID)
    if err != nil {
        http.Error(w, "Failed to get user rating: "+err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "user_rating": rating,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}