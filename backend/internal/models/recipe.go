package models

import (
    "time"
)

type Recipe struct {
    ID            string    `json:"id"`
    Title         string    `json:"title"`
    Description   string    `json:"description"`
    PrepTime      int       `json:"prep_time"`  // minutes only
    FeaturedImage string    `json:"featured_image,omitempty"`
    IsPremium     bool      `json:"is_premium"`
    Price         *float64  `json:"price,omitempty"`
    UserID        string    `json:"user_id"`
    CategoryID    string    `json:"category_id"`
    Status        string    `json:"status"` // draft, published
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    
    // Relationships
    User         *User         `json:"user,omitempty"`
    Category     *Category     `json:"category,omitempty"`
    Ingredients  []Ingredient  `json:"ingredients,omitempty"`
    Steps        []Step        `json:"steps,omitempty"`
    RecipeImages []RecipeImage `json:"recipe_images,omitempty"`
}

type CreateRecipeRequest struct {
    Title         string                    `json:"title"`
    Description   string                    `json:"description"`
    PrepTime      int                       `json:"prep_time"` // minutes only
    FeaturedImage string                    `json:"featured_image,omitempty"`
    AllImages     []string                  `json:"all_images,omitempty"` // NEW: All uploaded images
    IsPremium     bool                      `json:"is_premium"`
    Price         *float64                  `json:"price,omitempty"`
    CategoryID    string                    `json:"category_id"`
    Ingredients   []CreateIngredientRequest `json:"ingredients"`
    Steps         []CreateStepRequest       `json:"steps"`
}

type RecipeImage struct {
    ID        string    `json:"id"`
    RecipeID  string    `json:"recipe_id"`
    ImageURL  string    `json:"image_url"`
    IsFeatured bool     `json:"is_featured"`
    CreatedAt time.Time `json:"created_at"`
}