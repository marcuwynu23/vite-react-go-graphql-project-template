package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DatabaseURL    string
	CORSOrigins    []string
	Environment    string
	SeedEmail      string
	SeedPassword   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	origins := os.Getenv("CORS_ORIGINS")
	if strings.TrimSpace(origins) == "" {
		origins = "http://localhost:5173"
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: strings.TrimSpace(os.Getenv("DB_URL")),
		CORSOrigins: splitComma(origins),
		Environment: getEnv("ENV", "development"),
		SeedEmail:   getEnv("SEED_EMAIL", "admin@example.com"),
		SeedPassword: getEnv("SEED_PASSWORD", "admin123"),
	}, nil
}

func getEnv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func splitComma(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	if len(out) == 0 {
		return []string{"http://localhost:5173"}
	}
	return out
}
