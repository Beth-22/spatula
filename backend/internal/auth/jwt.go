package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte

type HasuraClaims struct {
	XHasuraAllowedRoles []string `json:"x-hasura-allowed-roles"`
	XHasuraDefaultRole  string   `json:"x-hasura-default-role"`
	XHasuraUserID       string   `json:"x-hasura-user-id"`
}

func GenerateJWT(userID string) (string, error) {
	hasuraClaims := HasuraClaims{
		XHasuraAllowedRoles: []string{"user"},
		XHasuraDefaultRole:  "user",
		XHasuraUserID:       userID,
	}

	claims := jwt.MapClaims{
		"sub":                         userID,
		"iss":                         "spatula-auth",
		"exp":                         time.Now().Add(24 * time.Hour).Unix(),
		"iat":                         time.Now().Unix(),
		"https://hasura.io/jwt/claims": hasuraClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	
	if err != nil {
		return nil, err // Return the actual error
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, fmt.Errorf("invalid token") // Use a simple error message
}
