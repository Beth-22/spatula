package auth

import (
	"context"
	"errors"
	//"fmt"
	"net/http"
	"strings"
)

// Define custom types for context keys to avoid collisions
type contextKey string

const (
	userIDKey   contextKey = "user_id"
	userRoleKey contextKey = "user_role"
)

var (
	ErrUserIDNotFound   = errors.New("user ID not found in context")
	ErrUserRoleNotFound = errors.New("user role not found in context")
)

// Middleware validates JWT tokens and adds user info to context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for public endpoints
		if r.URL.Path == "/actions/signup" || r.URL.Path == "/actions/login" || r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check if it's a Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorization header format must be: Bearer {token}", http.StatusUnauthorized)
			return
		}

		// Extract the token
		tokenString := parts[1]

		// Validate the token
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Extract user ID from claims
		hasuraClaims, ok := claims["https://hasura.io/jwt/claims"].(map[string]interface{})
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID, ok := hasuraClaims["x-hasura-user-id"].(string)
		if !ok || userID == "" {
			http.Error(w, "User ID not found in token", http.StatusUnauthorized)
			return
		}

		// Add user info to request context using type-safe keys
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		
		if role, ok := hasuraClaims["x-hasura-default-role"].(string); ok {
			ctx = context.WithValue(ctx, userRoleKey, role)
		}

		// Proceed with the request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper function to get user ID from context
func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok || userID == "" {
		return "", ErrUserIDNotFound
	}
	return userID, nil
}

// Helper function to get user role from context
func GetUserRoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value(userRoleKey).(string)
	if !ok || role == "" {
		return "", ErrUserRoleNotFound
	}
	return role, nil
}