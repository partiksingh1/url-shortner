package main

import (
	"log"
	"net/http"
	"os"
	"url-shortener/config"
	"url-shortener/models"
	"url-shortener/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file (if exists)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	db, err := initDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	// Perform automatic migration
	db.AutoMigrate(&models.URL{})

	// Initialize the router
	r := mux.NewRouter()

	// Setup application routes
	routes.SetupRoutes(r, db)

	// Configure CORS settings
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // React app frontend (can be adjusted)
			"http://localhost:5173", // Vite or other frontends
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache pre-flight request for 5 minutes
	})

	// Apply CORS middleware to the router
	handler := corsOptions.Handler(r)

	// Get the server port from the environment variable (default to 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// initDB initializes the database connection using the configuration.
func initDB() (*gorm.DB, error) {
	// Get the database configuration (DSN URL from environment variable)
	dbConfig := config.GetDBConfig()

	// Use the PostgreSQL connection URL (DSN)
	db, err := gorm.Open(postgres.Open(dbConfig.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Return the database connection object
	return db, nil
}
