package repository

import (
    "context"
    "fmt"
    "log"
	"time"

    "github.com/Beth-22/spatula/internal/graphql"
    "github.com/Beth-22/spatula/internal/models"
    "golang.org/x/crypto/bcrypt"
)

// GetUserByID retrieves a user by their ID
func GetUserByID(userID string) (*models.User, error) {
    query := `
    query GetUser($id: uuid!) {
        users_by_pk(id: $id) {
            id
            email
            username
            created_at
            updated_at
        }
    }`

    variables := map[string]interface{}{"id": userID}
    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return nil, fmt.Errorf("failed to query user: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid response format")
    }

    userData, ok := data["users_by_pk"].(map[string]interface{})
    if !ok || userData == nil {
        return nil, fmt.Errorf("user not found")
    }

    user := &models.User{
        ID:        getString(userData, "id"),
        Email:     getString(userData, "email"),
        Username:  getString(userData, "username"),
        CreatedAt: parseTime(userData["created_at"]),
        UpdatedAt: parseTime(userData["updated_at"]),
    }

    return user, nil
}

// UpdateUser updates user profile information
func UpdateUser(userID string, updates map[string]interface{}) error {
    query := `
    mutation UpdateUser($id: uuid!, $changes: users_set_input!) {
        update_users_by_pk(pk_columns: {id: $id}, _set: $changes) {
            id
            username
            email
        }
    }`

    variables := map[string]interface{}{
        "id":      userID,
        "changes": updates,
    }

    _, err := graphql.Mutate(context.Background(), query, variables)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }

    log.Printf("Successfully updated user: %s", userID)
    return nil
}

// UpdatePassword updates user password with verification
func UpdatePassword(userID, currentPassword, newPassword string) error {
    // First verify current password
    query := `
    query GetUserPassword($id: uuid!) {
        users_by_pk(id: $id) {
            password_hash
        }
    }`

    variables := map[string]interface{}{"id": userID}
    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return fmt.Errorf("failed to query user: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return fmt.Errorf("invalid response format")
    }

    userData, ok := data["users_by_pk"].(map[string]interface{})
    if !ok || userData == nil {
        return fmt.Errorf("user not found")
    }

    storedHash, ok := userData["password_hash"].(string)
    if !ok {
        return fmt.Errorf("failed to get password hash")
    }

    // Verify current password
    err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(currentPassword))
    if err != nil {
        return fmt.Errorf("current password is incorrect")
    }

    // Hash new password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("failed to hash new password: %w", err)
    }

    // Update password
    updateQuery := `
    mutation UpdatePassword($id: uuid!, $new_hash: String!) {
        update_users_by_pk(pk_columns: {id: $id}, _set: {password_hash: $new_hash}) {
            id
        }
    }`

    updateVariables := map[string]interface{}{
        "id":       userID,
        "new_hash": string(hashedPassword),
    }

    _, err = graphql.Mutate(context.Background(), updateQuery, updateVariables)
    if err != nil {
        return fmt.Errorf("failed to update password: %w", err)
    }

    log.Printf("Successfully updated password for user: %s", userID)
    return nil
}

// DeleteUser deletes a user account
func DeleteUser(userID, password string) error {
    // First verify password
    query := `
    query GetUserPassword($id: uuid!) {
        users_by_pk(id: $id) {
            password_hash
        }
    }`

    variables := map[string]interface{}{"id": userID}
    resp, err := graphql.Query(context.Background(), query, variables)
    if err != nil {
        return fmt.Errorf("failed to query user: %w", err)
    }

    data, ok := resp.Data.(map[string]interface{})
    if !ok {
        return fmt.Errorf("invalid response format")
    }

    userData, ok := data["users_by_pk"].(map[string]interface{})
    if !ok || userData == nil {
        return fmt.Errorf("user not found")
    }

    storedHash, ok := userData["password_hash"].(string)
    if !ok {
        return fmt.Errorf("failed to get password hash")
    }

    // Verify password
    err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
    if err != nil {
        return fmt.Errorf("password is incorrect")
    }

    // Delete user (this will cascade delete related data due to foreign keys)
    deleteQuery := `
    mutation DeleteUser($id: uuid!) {
        delete_users_by_pk(id: $id) {
            id
        }
    }`

    deleteVariables := map[string]interface{}{"id": userID}
    _, err = graphql.Mutate(context.Background(), deleteQuery, deleteVariables)
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }

    log.Printf("Successfully deleted user: %s", userID)
    return nil
}

// Helper function to parse time
func parseTime(timeInterface interface{}) time.Time {
    if timeStr, ok := timeInterface.(string); ok {
        if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
            return t
        }
    }
    return time.Time{}
}