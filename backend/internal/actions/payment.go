package actions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"log"

	"github.com/Beth-22/spatula/internal/auth"
	"github.com/Beth-22/spatula/internal/models"
	"github.com/Beth-22/spatula/internal/repository"
	"github.com/go-resty/resty/v2"
)

type PaymentRequest struct {
	RecipeID    string  `json:"recipe_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	PhoneNumber string  `json:"phone_number"`
}

type ChapaPaymentResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

type PaymentResponse struct {
	CheckoutURL string `json:"checkout_url"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	PurchaseID  string `json:"purchase_id"`
}

type ChapaWebhookPayload struct {
	Event string `json:"event"`
	Data  struct {
		TxRef          string  `json:"tx_ref"`
		Amount         float64 `json:"amount"`
		Currency       string  `json:"currency"`
		Status         string  `json:"status"`
		TransactionID  string  `json:"id"`
		ChapaReference string  `json:"chapa_reference"`
		CreatedAt      string  `json:"created_at"`
	} `json:"data"`
}

type PaymentVerificationRequest struct {
	PurchaseID string `json:"purchase_id"`
}

type PaymentVerificationResponse struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RecipeAccessResponse struct {
	HasAccess bool     `json:"has_access"`
	IsPremium bool     `json:"is_premium"`
	Price     *float64 `json:"price,omitempty"`
	Reason    string   `json:"reason,omitempty"`
}

func InitiatePaymentHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("=== PAYMENT INITIATION STARTED ===")
    log.Printf("Request Method: %s", r.Method)
    log.Printf("Request URL: %s", r.URL.String())

    // Get user ID from JWT
    userID, err := auth.GetUserIDFromContext(r.Context())
    if err != nil {
        log.Printf("AUTH ERROR: %v", err)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    log.Printf("User authenticated: %s", userID)

    var paymentReq PaymentRequest
    if err := json.NewDecoder(r.Body).Decode(&paymentReq); err != nil {
        log.Printf("JSON DECODE ERROR: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    log.Printf("Payment Request: %+v", paymentReq)

    // Verify recipe exists and is premium
    recipe, err := repository.GetRecipeByID(paymentReq.RecipeID)
    if err != nil {
        log.Printf("RECIPE NOT FOUND ERROR: RecipeID=%s, Error=%v", paymentReq.RecipeID, err)
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }

    log.Printf("Recipe found: %s (Premium: %t)", recipe.Title, recipe.IsPremium)

    if !recipe.IsPremium {
        log.Printf("RECIPE NOT PREMIUM: %s", paymentReq.RecipeID)
        http.Error(w, "Recipe is not premium", http.StatusBadRequest)
        return
    }

    // Check if user already purchased this recipe
    hasPurchased, err := repository.HasUserPurchasedRecipe(userID, paymentReq.RecipeID)
    if err != nil {
        log.Printf("PURCHASE CHECK ERROR: %v", err)
        http.Error(w, "Error checking purchase status", http.StatusInternalServerError)
        return
    }

    if hasPurchased {
        log.Printf("ALREADY PURCHASED: user=%s, recipe=%s", userID, paymentReq.RecipeID)
        http.Error(w, "You have already purchased this recipe", http.StatusConflict)
        return
    }

    // Create a pending purchase record
    purchase := &models.Purchase{
        UserID:   userID,
        RecipeID: paymentReq.RecipeID,
        Amount:   paymentReq.Amount,
        Status:   "pending",
    }

    purchaseID, err := repository.CreatePurchase(purchase)
    if err != nil {
        log.Printf("PURCHASE CREATION ERROR: %v", err)
        http.Error(w, "Failed to create purchase record", http.StatusInternalServerError)
        return
    }

    log.Printf("Purchase record created: %s", purchaseID)

    // Initialize Chapa payment
    client := resty.New()
    chapaURL := os.Getenv("CHAPA_BASE_URL") + "/v1/transaction/initialize"
    
    log.Printf("CHAPA CONFIG:")
    log.Printf("  URL: %s", chapaURL)
    log.Printf("  Secret Key present: %t", os.Getenv("CHAPA_SECRET_KEY") != "")

    var chapaResp ChapaPaymentResponse
    
    // FIX: Shorter transaction reference
    txRef := fmt.Sprintf("rcp-%s", purchaseID)
    if len(txRef) > 50 {
        txRef = txRef[:50]
    }
    
    // FIX: Clean description - remove special characters
    description := "Premium Recipe Access"
    if len(recipe.Title) > 20 {
        description = "Recipe " + recipe.Title[:20]
    } else {
        description = "Recipe " + recipe.Title
    }
    // Remove any problematic characters
    description = strings.ReplaceAll(description, ":", "")
    description = strings.ReplaceAll(description, "/", "")
    description = strings.ReplaceAll(description, "\\", "")
    description = strings.ReplaceAll(description, "@", "")
    description = strings.ReplaceAll(description, "&", "and")

    log.Printf("Transaction reference: %s (length: %d)", txRef, len(txRef))

    // Prepare request body with validation fixes
    requestBody := map[string]interface{}{
        "amount":         paymentReq.Amount,
        "currency":       paymentReq.Currency,
        "email":          "test@gmail.com", // Simple valid email
        "first_name":     "Customer",
        "last_name":      "Name",
        "phone_number":   paymentReq.PhoneNumber,
        "tx_ref":         txRef,
        "callback_url":   os.Getenv("CHAPA_WEBHOOK_URL") + "/actions/payment/webhook",
        "return_url":     os.Getenv("FRONTEND_URL") + "/payment/success?purchase_id=" + purchaseID,
        "customization": map[string]interface{}{
            "title":       "Spatula",
            "description": description,
        },
    }

    log.Printf("Chapa Request Body: %+v", requestBody)

    // Make the API call
    log.Printf("Making Chapa API call...")
    
    resp, err := client.R().
        SetAuthToken(os.Getenv("CHAPA_SECRET_KEY")).
        SetHeaders(map[string]string{
            "Content-Type": "application/json",
        }).
        SetBody(requestBody).
        SetResult(&chapaResp).
        Post(chapaURL)

    if err != nil {
        log.Printf("CHAPA API CALL ERROR: %v", err)
        repository.UpdatePurchaseStatus(purchaseID, "failed", nil)
        http.Error(w, fmt.Sprintf("Payment initiation failed: %v", err), http.StatusInternalServerError)
        return
    }

    log.Printf("Chapa Response Status: %d", resp.StatusCode())
    log.Printf("Chapa Response Body: %s", resp.String())

    if resp.StatusCode() != http.StatusOK {
        log.Printf("CHAPA API RETURNED NON-200 STATUS: %d", resp.StatusCode())
        repository.UpdatePurchaseStatus(purchaseID, "failed", nil)
        http.Error(w, fmt.Sprintf("Chapa API error: HTTP %d - %s", resp.StatusCode(), resp.String()), http.StatusInternalServerError)
        return
    }

    if chapaResp.Status != "success" {
        log.Printf("CHAPA BUSINESS ERROR: Status=%s, Message=%s", chapaResp.Status, chapaResp.Message)
        repository.UpdatePurchaseStatus(purchaseID, "failed", nil)
        http.Error(w, fmt.Sprintf("Chapa error: %s", chapaResp.Message), http.StatusInternalServerError)
        return
    }

    log.Printf("CHAPA PAYMENT INITIATED SUCCESSFULLY")
    log.Printf("Checkout URL: %s", chapaResp.Data.CheckoutURL)

    // Return success response
    response := PaymentResponse{
        CheckoutURL: chapaResp.Data.CheckoutURL,
        Success:     true,
        Message:     "Payment initiated successfully",
        PurchaseID:  purchaseID,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    
    log.Printf("=== PAYMENT INITIATION COMPLETED ===")
}

func PaymentWebhookHandler(w http.ResponseWriter, r *http.Request) {
	// Verify webhook signature (basic implementation)
	signature := r.Header.Get("Chapa-Signature")
	if signature == "" {
		http.Error(w, "Missing signature", http.StatusBadRequest)
		return
	}

	var webhookPayload ChapaWebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&webhookPayload); err != nil {
		http.Error(w, "Invalid webhook payload", http.StatusBadRequest)
		return
	}

	// Extract purchase ID from tx_ref (format: "recipe-{recipeID}-{userID}-{purchaseID}")
	txRef := webhookPayload.Data.TxRef
	parts := strings.Split(txRef, "-")
	if len(parts) < 4 {
		http.Error(w, "Invalid transaction reference", http.StatusBadRequest)
		return
	}

	purchaseID := parts[3]

	// Update purchase status based on webhook event
	if webhookPayload.Event == "charge.complete" {
		if webhookPayload.Data.Status == "success" {
			err := repository.UpdatePurchaseStatus(purchaseID, "completed", &webhookPayload.Data.TransactionID)
			if err != nil {
				log.Printf("Failed to update purchase status: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			log.Printf("Payment completed for purchase: %s", purchaseID)
		} else {
			err := repository.UpdatePurchaseStatus(purchaseID, "failed", nil)
			if err != nil {
				log.Printf("Failed to update purchase status: %v", err)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func VerifyPaymentHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var verifyReq PaymentVerificationRequest
	if err := json.NewDecoder(r.Body).Decode(&verifyReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get purchase details
	purchase, err := repository.GetPurchaseByID(verifyReq.PurchaseID)
	if err != nil {
		http.Error(w, "Purchase not found", http.StatusNotFound)
		return
	}

	// Verify the purchase belongs to the user
	if purchase.UserID != userID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	response := PaymentVerificationResponse{
		Status:  purchase.Status,
		Success: purchase.Status == "completed",
		Message: fmt.Sprintf("Payment status: %s", purchase.Status),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CheckRecipeAccessHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	recipeID := r.URL.Query().Get("recipe_id")
	if recipeID == "" {
		http.Error(w, "Recipe ID required", http.StatusBadRequest)
		return
	}

	// Get recipe details
	recipe, err := repository.GetRecipeByID(recipeID)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	// If recipe is not premium, allow access
	if !recipe.IsPremium {
		response := RecipeAccessResponse{
			HasAccess: true,
			IsPremium: false,
			Reason:    "Recipe is free",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// For premium recipes, check if user has purchased it
	hasPurchased, err := repository.HasUserPurchasedRecipe(userID, recipeID)
	if err != nil {
		http.Error(w, "Error checking purchase status", http.StatusInternalServerError)
		return
	}

	response := RecipeAccessResponse{
		HasAccess: hasPurchased,
		IsPremium: true,
		Price:     recipe.Price,
	}

	if !hasPurchased {
		response.Reason = "Purchase required"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}