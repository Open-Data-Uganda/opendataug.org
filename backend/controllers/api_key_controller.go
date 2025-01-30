package controllers

import (
	"time"

	"github.com/google/uuid"
	"github.com/uganda-data/database"
	"github.com/uganda-data/models"
	"gorm.io/gorm"
)

type APIKeyController struct {
	db *database.Database
}

func NewAPIKeyController(db *database.Database) *APIKeyController {
	return &APIKeyController{db: db}
}

func (c *APIKeyController) CreateAPIKey(apiKey *models.APIKey) error {
	return c.db.DB.Create(apiKey).Error
}

func (c *APIKeyController) DeleteAPIKey(userID string, keyID uuid.UUID) error {
	result := c.db.DB.Where("user_number = ? AND number = ?", userID, keyID).Delete(&models.APIKey{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (c *APIKeyController) GetAPIKeys(userID string) ([]models.APIKey, error) {
	var keys []models.APIKey
	err := c.db.DB.Where("user_number = ?", userID).Find(&keys).Error
	return keys, err
}

func (c *APIKeyController) GetAPIKeyByNumber(key string) (*models.APIKey, error) {
	var apiKey models.APIKey
	err := c.db.DB.Where("number = ?", key).First(&apiKey).Error
	if err != nil {
		return nil, err
	}
	return &apiKey, nil
}

func (c *APIKeyController) UpdateAPIKeyLastUsed(keyID uuid.UUID) error {
	now := time.Now()
	return c.db.DB.Model(&models.APIKey{}).Where("number = ?", keyID).
		Update("last_used_at", now).Error
}

func (c *APIKeyController) UpdateAPIKeyUsage(apiKeyID string) error {
	result := c.db.DB.Model(&models.APIKey{}).
		Where("number = ?", apiKeyID).
		Updates(map[string]interface{}{
			"last_used_at": time.Now(),
			"usage_count":  gorm.Expr("usage_count + 1"),
		})
	return result.Error
}
