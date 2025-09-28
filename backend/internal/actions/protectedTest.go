package actions

import (
	"encoding/json"
	"net/http"
	
	"github.com/Beth-22/spatula/internal/auth"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"message": "Welcome to protected route!",
		"user_id": userID,
		"status":  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}