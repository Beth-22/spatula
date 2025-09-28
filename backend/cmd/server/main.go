package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Beth-22/spatula/internal/actions"
	"github.com/Beth-22/spatula/internal/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from current directory
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Initialize JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}
	auth.JWTSecret = []byte(jwtSecret)

	// Check if Hasura environment variables are set
	hasuraURL := os.Getenv("HASURA_GRAPHQL_URL")
	hasuraSecret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if hasuraURL == "" || hasuraSecret == "" {
		log.Fatal("Hasura environment variables not set. Check your .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Create chi router
	r := chi.NewRouter()
	
	// Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Public routes (no authentication required)
	r.Post("/actions/signup", actions.SignupHandler)
	r.Post("/actions/login", actions.LoginHandler)
	r.Post("/actions/payment/webhook", actions.PaymentWebhookHandler) // Webhook should be public
	r.Get("/actions/recipes", actions.GetRecipeHandler)
	
	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Protected routes (require authentication)
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		
		// File upload
		r.Post("/actions/upload", actions.UploadHandler) 
		
		// Recipe management
		r.Post("/actions/recipes/create", actions.CreateRecipeHandler)
		r.Get("/actions/recipes/my-recipes", actions.GetMyRecipesHandler)
		
		// Profile management
		r.Get("/actions/profile", actions.GetProfileHandler)
		r.Put("/actions/profile", actions.UpdateProfileHandler)
		r.Delete("/actions/profile", actions.DeleteAccountHandler)
		
		// Comment routes
		r.Post("/actions/comments", actions.CreateCommentHandler)
		r.Get("/actions/recipes/{recipeId}/comments", actions.GetRecipeCommentsHandler)
		r.Delete("/actions/comments/{commentId}", actions.DeleteCommentHandler)
		
		// Like and rating routes
		r.Post("/actions/recipes/like", actions.ToggleLikeHandler)
		r.Get("/actions/recipes/like", actions.GetUserLikeHandler)
		r.Post("/actions/recipes/rate", actions.SetRatingHandler)
		r.Get("/actions/recipes/rating", actions.GetUserRatingHandler)
		
		// Payment routes
		r.Post("/actions/payment/initiate", actions.InitiatePaymentHandler)
		r.Post("/actions/payment/verify", actions.VerifyPaymentHandler)
		r.Get("/actions/recipes/access", actions.CheckRecipeAccessHandler) // Fixed route path
		
		// Test route
		r.Post("/actions/protectedTest", actions.ProtectedHandler)
	})

	// Start the HTTP server!
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}