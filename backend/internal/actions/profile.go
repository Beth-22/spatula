package actions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Beth-22/spatula/internal/auth"
	"github.com/Beth-22/spatula/internal/models"
	"github.com/Beth-22/spatula/internal/repository"
)

type ProfileUpdateRequest struct {
    Username        string `json:"username,omitempty"`
    Email           string `json:"email,omitempty"`
    CurrentPassword string `json:"currentPassword,omitempty"`
    NewPassword     string `json:"newPassword,omitempty"`
}

type ProfileResponse struct {
    User *models.User `json:"user"`
}

// GetProfileHandler returns the current user's profile
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    user, err := repository.GetUserByID(userID)
    if err != nil {
        log.Printf("GetProfile error: %v", err)
        http.Error(w, "Failed to get user profile", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ProfileResponse{User: user})
}

// UpdateProfileHandler updates user profile
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var req ProfileUpdateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Build update object
    updates := make(map[string]interface{})
    if req.Username != "" {
        updates["username"] = req.Username
    }
    if req.Email != "" {
        updates["email"] = req.Email
    }

    // Update profile if there are changes
    if len(updates) > 0 {
        err = repository.UpdateUser(userID, updates)
        if err != nil {
            log.Printf("UpdateProfile error: %v", err)
            http.Error(w, "Failed to update profile", http.StatusInternalServerError)
            return
        }
    }

    // Update password if provided
    if req.NewPassword != "" {
        if req.CurrentPassword == "" {
            http.Error(w, "Current password is required to change password", http.StatusBadRequest)
            return
        }

        err = repository.UpdatePassword(userID, req.CurrentPassword, req.NewPassword)
        if err != nil {
            log.Printf("UpdatePassword error: %v", err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
    }

    // Return updated user
    user, err := repository.GetUserByID(userID)
    if err != nil {
        log.Printf("GetUser after update error: %v", err)
        http.Error(w, "Profile updated but failed to fetch updated data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ProfileResponse{User: user})
}

// DeleteAccountHandler deletes user account
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var req struct {
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if req.Password == "" {
        http.Error(w, "Password is required", http.StatusBadRequest)
        return
    }

    err = repository.DeleteUser(userID, req.Password)
    if err != nil {
        log.Printf("DeleteAccount error: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Account deleted successfully"})
}