package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds every setting the service needs to run.
type Config struct {
	Port string
	Env  string
}

// Load reads configuration from the environment and returns an error
// if any value is invalid.
func Load() (Config, error) {
	// A .env file is a development convenience. Its absence is not an error:
	// in staging and production the environment comes from the platform.
	_ = godotenv.Load()

	cfg := Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("APP_ENV", "development"),
	}

	switch cfg.Env {
	case "development", "staging", "production":
		// valid
	default:
		return Config{}, fmt.Errorf("invalid APP_ENV %q: must be development, staging or production", cfg.Env)
	}
	return cfg, nil
}

// getEnv returns the value of the environment variable named by key,
// or fallback if the variable is unset or empty.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
