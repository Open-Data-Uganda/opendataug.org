package controllers

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"opendataug.org/database"
	"opendataug.org/models"
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

func (c *APIKeyController) DeleteAPIKey(userID string, keyID string) error {
	result := c.db.DB.Where("user_number = ? AND number = ?", userID, keyID).Delete(&models.APIKey{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (c *APIKeyController) GetAPIKeys(userID string) ([]models.APIKey, error) {
	var keys []models.APIKey
	err := c.db.DB.Where("user_number = ?", userID).Order("created_at DESC").Find(&keys).Error
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

func (c *APIKeyController) APIKeyNameExists(userNumber string, name string) (bool, error) {
	var count int64
	result := c.db.DB.Model(&models.APIKey{}).Where("user_number = ? AND name = ?", userNumber, name).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}
