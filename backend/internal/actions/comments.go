package actions

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Beth-22/spatula/internal/auth"
	"github.com/Beth-22/spatula/internal/models"
	"github.com/Beth-22/spatula/internal/repository"
	"github.com/go-chi/chi/v5"
)

// CreateCommentHandler creates a new comment for a recipe
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
    // Get user ID from context using your existing helper
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var req models.CreateCommentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate request
    if req.RecipeID == "" {
        http.Error(w, "Recipe ID is required", http.StatusBadRequest)
        return
    }
    if req.Content == "" || len(req.Content) > 1000 {
        http.Error(w, "Comment content must be between 1 and 1000 characters", http.StatusBadRequest)
        return
    }

    // Create comment via repository
    commentID, err := repository.CreateComment(userID, req.RecipeID, req.Content)
    if err != nil {
        log.Printf("Error creating comment: %v", err)
        http.Error(w, "Failed to create comment", http.StatusInternalServerError)
        return
    }

    // Return success response
    response := map[string]interface{}{
        "success": true,
        "message": "Comment created successfully",
        "comment_id": commentID,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// GetRecipeCommentsHandler retrieves comments for a recipe
func GetRecipeCommentsHandler(w http.ResponseWriter, r *http.Request) {
    recipeID := chi.URLParam(r, "recipeId")
    if recipeID == "" {
        http.Error(w, "Recipe ID is required", http.StatusBadRequest)
        return
    }

    // Get pagination parameters
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }
    
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if limit < 1 || limit > 50 {
        limit = 20
    }

    offset := (page - 1) * limit

    // Get comments from repository
    comments, total, err := repository.GetRecipeComments(recipeID, limit, offset)
    if err != nil {
        log.Printf("Error fetching comments: %v", err)
        http.Error(w, "Failed to fetch comments", http.StatusInternalServerError)
        return
    }

    // Prepare response
    response := models.CommentsListResponse{
        Comments: comments,
        Total:    total,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// DeleteCommentHandler deletes a comment (only by owner)
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    commentID := chi.URLParam(r, "commentId")
    if commentID == "" {
        http.Error(w, "Comment ID is required", http.StatusBadRequest)
        return
    }

    // Delete comment via repository (will verify ownership)
    err = repository.DeleteComment(commentID, userID)
    if err != nil {
        log.Printf("Error deleting comment: %v", err)
        http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "success": true,
        "message": "Comment deleted successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}