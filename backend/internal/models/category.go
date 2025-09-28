package models

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}