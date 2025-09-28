package models

import (
	"time"
)

type Purchase struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	RecipeID           string    `json:"recipe_id"`
	Amount             float64   `json:"amount"`
	Status             string    `json:"status"` // pending, completed, failed
	ChapaTransactionID *string   `json:"chapa_transaction_id,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type PurchaseRequest struct {
	RecipeID    string  `json:"recipe_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	PhoneNumber string  `json:"phone_number"`
}

type PurchaseResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}