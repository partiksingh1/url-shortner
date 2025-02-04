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
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize database
	db, err := initDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	db.AutoMigrate(&models.URL{})

	// Initialize router
	r := mux.NewRouter()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:5173"}, // Add your frontend URLs
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
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Create a handler with CORS middleware
	handler := c.Handler(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func initDB() (*gorm.DB, error) {
	config := config.GetDBConfig()
	dsn := config.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
