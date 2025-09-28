package models

import (
    "time"
)

type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Username  string    `json:"username"`
    Password  string    `json:"-"` // Hidden from JSON
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type SignupRequest struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Token string `json:"token"`
    User  *User  `json:"user,omitempty"`
}