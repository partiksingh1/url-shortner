package repository

import (
	"errors"
	"url-shortener/models"

	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *models.URL) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) FindByShortURL(shortURL string) (*models.URL, error) {
	var url models.URL
	if err := r.db.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) IncrementClickCount(shortURL string) error {
	return r.db.Model(&models.URL{}).Where("short_url = ?", shortURL).UpdateColumn("click_count", gorm.Expr("click_count + ?", 1)).Error
}
