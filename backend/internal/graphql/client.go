package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// GraphQLRequest represents the payload we send to Hasura
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// GraphQLResponse represents the response from Hasura
type GraphQLResponse struct {
	Data   interface{}   `json:"data"`
	Errors []GraphQLError `json:"errors"`
}

type GraphQLError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

// Remove the init() function and make these functions
func getHasuraURL() string {
	url := os.Getenv("HASURA_GRAPHQL_URL")
	if url == "" {
		panic("HASURA_GRAPHQL_URL environment variable is required")
	}
	return url
}

func getAdminSecret() string {
	secret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if secret == "" {
		panic("HASURA_GRAPHQL_ADMIN_SECRET environment variable is required")
	}
	return secret
}

func Query(ctx context.Context, query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	return executeGraphQL(ctx, query, variables)
}

func Mutate(ctx context.Context, query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	return executeGraphQL(ctx, query, variables)
}

func executeGraphQL(ctx context.Context, query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	hasuraURL := getHasuraURL()
	adminSecret := getAdminSecret()

	payload := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", hasuraURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", adminSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hasura returned non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	var result GraphQLResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
	}

	if len(result.Errors) > 0 {
		errorMessages := make([]string, len(result.Errors))
		for i, err := range result.Errors {
			errorMessages[i] = err.Message
		}
		return nil, fmt.Errorf("graphql errors: %v", errorMessages)
	}

	return &result, nil
}