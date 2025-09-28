package actions

import (
    "encoding/json"
    "net/http"
    "github.com/Beth-22/spatula/internal/auth"  
    "github.com/Beth-22/spatula/internal/repository"  
    "log"
)

type SignupInput struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Token string `json:"token"`
}

// SignupHandler handles Hasura signup action
func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var input SignupInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        log.Printf("Signup error - JSON decode failed: %v", err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    log.Printf("Signup attempt for email: %s, username: %s", input.Email, input.Username)

    userID, err := repository.CreateUser(input.Email, input.Username, input.Password)
    if err != nil {
        log.Printf("Signup error - CreateUser failed: %v", err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    token, err := auth.GenerateJWT(userID)
    if err != nil {
        log.Printf("Signup error - GenerateJWT failed: %v", err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate token"})
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(AuthResponse{Token: token})
}
// LoginHandler handles Hasura login action
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var input LoginInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    userID, err := repository.VerifyUser(input.Email, input.Password)
    if err != nil {
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := auth.GenerateJWT(userID)  // Removed "user" role parameter
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(AuthResponse{Token: token})
}