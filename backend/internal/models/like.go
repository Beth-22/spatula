package models

import "time"

type Like struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    RecipeID  string    `json:"recipe_id"`
    CreatedAt time.Time `json:"created_at"`
}

type LikeRequest struct {
    RecipeID string `json:"recipe_id" validate:"required,uuid"`
}

type LikeResponse struct {
    ID        string    `json:"id"`
    Liked     bool      `json:"liked"`
    LikeCount int       `json:"like_count"`
    CreatedAt time.Time `json:"created_at,omitempty"`
}