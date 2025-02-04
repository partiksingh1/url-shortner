package config

import (
	"os"
)

type DBConfig struct {
	DSN string
}

// GetDBConfig now returns a DBConfig with the full connection URL
func GetDBConfig() *DBConfig {
	// Fetch the PostgreSQL URL from environment variable
	// For local development, you can set it in your .env file
	url := getEnv("DATABASE_URL", "")

	// If the DATABASE_URL is empty, return an error (you can decide what to do here)
	if url == "" {
		panic("DATABASE_URL is not set!")
	}

	return &DBConfig{DSN: url}
}

// getEnv fetches the environment variable or returns the default value if not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
