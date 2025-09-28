package models

type Ingredient struct {
	ID       string `json:"id"`
	RecipeID string `json:"recipe_id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit,omitempty"`
}

type CreateIngredientRequest struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit,omitempty"`
}