package service

import (
	"crypto/rand"
	"encoding/base64"
	"time"
	"url-shortener/models"
	"url-shortener/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) CreateShortUrl(longUrl string) (*models.URL, error) {
	shortUrl, err := generateShortUrl()
	if err != nil {
		return nil, err
	}
	url := &models.URL{
		LongUrl:   longUrl,
		ShortUrl:  shortUrl,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().AddDate(0, 1, 0),
	}
	if err := s.repo.Create(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *URLService) GetLongUrl(shortURL string) (*models.URL, error) {
	url, err := s.repo.FindByShortURL(shortURL)
	if err != nil {
		return nil, err
	}
	if url != nil {
		s.repo.IncrementClickCount(shortURL)
	}
	return url, nil
}

func generateShortUrl() (string, error) {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:6], nil
}
