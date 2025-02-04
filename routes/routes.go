package routes

import (
	"url-shortener/handlers"
	"url-shortener/repository"
	"url-shortener/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(r *mux.Router, db *gorm.DB) {
	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handlers.NewUrlHandler(urlService)

	r.HandleFunc("/urls", urlHandler.CreateShortUrl).Methods("POST")
	r.HandleFunc("/{shortURL}", urlHandler.RedirectToLongURL).Methods("GET")
}
