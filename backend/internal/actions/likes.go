package actions

import (
    "encoding/json"
    "net/http"

    "github.com/Beth-22/spatula/internal/auth"
    "github.com/Beth-22/spatula/internal/models"
    "github.com/Beth-22/spatula/internal/repository"
)

func ToggleLikeHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Authentication required", http.StatusUnauthorized)
        return
    }

    var req models.LikeRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    liked, likeCount, err := repository.ToggleLike(userID, req.RecipeID)
    if err != nil {
        http.Error(w, "Failed to toggle like: "+err.Error(), http.StatusInternalServerError)
        return
    }

    response := models.LikeResponse{
        Liked:     liked,
        LikeCount: likeCount,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func GetUserLikeHandler(w http.ResponseWriter, r *http.Request) {
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

    liked, err := repository.GetUserLike(userID, recipeID)
    if err != nil {
        http.Error(w, "Failed to get like status: "+err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "liked": liked,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}