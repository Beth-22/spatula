package config

import (
    "os"
    "github.com/Beth-22/spatula/internal/auth"
)

type Config struct {
    Port                     string
    JWTSecret                string
    HasuraGraphQLURL         string
    HasuraGraphQLAdminSecret string
    SupabaseURL              string
    SupabaseKey              string
    SupabaseBucket           string
    ChapaSecretKey           string
    ChapaPublicKey           string
    ChapaWebhookSecret       string
    ChapaBaseURL             string
}

func Load() *Config {
    cfg := &Config{
        Port:                     getEnv("PORT", "8081"),
        JWTSecret:                getEnv("JWT_SECRET", "my_super_secret_key_that_is_at_least_32_chars"),
        HasuraGraphQLURL:         getEnv("HASURA_GRAPHQL_URL", "http://localhost:8080/v1/graphql"),
        HasuraGraphQLAdminSecret: getEnv("HASURA_GRAPHQL_ADMIN_SECRET", "myadminsecretkey"),
        SupabaseURL:              getEnv("SUPABASE_URL", ""),
        SupabaseKey:              getEnv("SUPABASE_KEY", ""),
        SupabaseBucket:           getEnv("SUPABASE_BUCKET", "recipe-images"),
        ChapaSecretKey:           getEnv("CHAPA_SECRET_KEY", ""),
        ChapaPublicKey:           getEnv("CHAPA_PUBLIC_KEY", ""),
        ChapaWebhookSecret:       getEnv("CHAPA_WEBHOOK_SECRET", ""),
        ChapaBaseURL:             getEnv("CHAPA_BASE_URL", "https://api.chapa.co"),
    }
    
    // Initialize JWT secret in auth package
    auth.JWTSecret = []byte(cfg.JWTSecret)
    
    return cfg
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}