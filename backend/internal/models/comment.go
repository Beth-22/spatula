package models

import (
	"time"
)

type Comment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	RecipeID  string    `json:"recipe_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	User *User `json:"user,omitempty"`
}

type CreateCommentRequest struct {
	RecipeID string `json:"recipe_id" validate:"required,uuid"`
	Content  string `json:"content" validate:"required,min=1,max=1000"`
}

type CommentResponse struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type CommentsListResponse struct {
	Comments []CommentResponse `json:"comments"`
	Total    int               `json:"total"`
}