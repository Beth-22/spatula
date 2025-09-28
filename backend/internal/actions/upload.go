package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	//"mime/multipart"
	"net/http"
	"net/url"
	//"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Beth-22/spatula/internal/auth"
)

type UploadResponse struct {
	URL     string `json:"url"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT for folder structure
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse multipart form (max 10MB file size)
	err = r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "File too large or invalid form", http.StatusBadRequest)
		return
	}

	// Get the file from form data
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create unique filename with user folder
	timestamp := time.Now().Unix()
	fileExt := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%s/%d%s", userID, timestamp, fileExt)
	
	// Read file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Upload to Supabase Storage using direct HTTP API
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	bucket := os.Getenv("SUPABASE_BUCKET")

	// Create the upload URL
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucket, fileName)

	// Create HTTP request
	req, err := http.NewRequest("POST", uploadURL, bytes.NewReader(fileBytes))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create request: %v", err), http.StatusInternalServerError)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", fileHeader.Header.Get("Content-Type"))
	req.Header.Set("Cache-Control", "max-age=3600")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Upload failed: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Supabase error: %s - %s", resp.Status, string(body)), http.StatusInternalServerError)
		return
	}

	// Get public URL - FIXED FORMAT
publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucket, url.PathEscape(fileName))

	// Return success response
	response := UploadResponse{
		URL:     publicURL,
		Success: true,
		Message: "File uploaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}