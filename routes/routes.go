package routes

import (
	"log"
	"url-shortener/handlers"
	"url-shortener/repository"
	"url-shortener/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// SetupRoutes sets up the URL shortener related routes.
func SetupRoutes(r *mux.Router, db *gorm.DB) {
	// Initialize repository, service, and handler
	urlRepo, err := repository.NewURLRepository(db) // Capture the error
	if err != nil {
		log.Fatalf("Error initializing URL repository: %v", err) // Log the error and stop execution if the repository initialization fails
	}

	urlService := service.NewURLService(urlRepo)
	urlHandler := handlers.NewUrlHandler(urlService)

	// Define the routes
	r.HandleFunc("/urls", urlHandler.CreateShortUrl).Methods("POST")
	r.HandleFunc("/{shortURL}", urlHandler.RedirectToLongURL).Methods("GET")

	// Optionally, you can add more routes here, for example:
	// r.HandleFunc("/urls/{shortURL}", urlHandler.GetURLDetails).Methods("GET")
	// r.HandleFunc("/urls", urlHandler.ListURLs).Methods("GET")
}
