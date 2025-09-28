package models

import "time"

type Rating struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    RecipeID  string    `json:"recipe_id"`
    Rating    int       `json:"rating" validate:"min=1,max=5"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type RatingRequest struct {
    RecipeID string `json:"recipe_id" validate:"required,uuid"`
    Rating   int    `json:"rating" validate:"required,min=1,max=5"`
}

type RatingResponse struct {
    ID           string    `json:"id,omitempty"`
    Rating       int       `json:"rating"`
    AverageRating float64  `json:"average_rating"`
    RatingCount  int       `json:"rating_count"`
    UserRating   int       `json:"user_rating"`
    CreatedAt    time.Time `json:"created_at,omitempty"`
}