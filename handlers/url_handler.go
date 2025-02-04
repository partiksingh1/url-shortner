package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/service"

	"github.com/gorilla/mux"
)

type UrlHandler struct {
	urlService *service.URLService
}

func NewUrlHandler(urlService *service.URLService) *UrlHandler {
	return &UrlHandler{urlService: urlService}
}

type CreateUrlRequest struct {
	LongUrl string `json:"long_url"`
}

func (h *UrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var req CreateUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	url, err := h.urlService.CreateShortUrl(req.LongUrl)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

func (h *UrlHandler) RedirectToLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]
	url, err := h.urlService.GetLongUrl(shortURL)
	if err != nil {
		http.Error(w, "Failed to retrieve URL", http.StatusInternalServerError)
		return
	}
	if url == nil {
		http.Error(w, "URL not found", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url.LongUrl, http.StatusMovedPermanently)
}
