package v1

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/controllers"
	"opendataug.org/database"
	"opendataug.org/models"
)

type APIKeyHandler struct {
	controller *controllers.APIKeyController
}

func NewAPIKeyHandler(db *database.Database) *APIKeyHandler {
	return &APIKeyHandler{
		controller: controllers.NewAPIKeyController(db),
	}
}

func (h *APIKeyHandler) RegisterRoutes(r *gin.RouterGroup) {
	keys := r.Group("/api-keys")
	{
		keys.GET("", h.listAPIKeys)
		keys.POST("", h.createAPIKey)
		keys.DELETE("/:id", h.deleteAPIKey)
	}
}

type CreateAPIKeyRequest struct {
	Name      string     `json:"name" binding:"required"`
	ExpiresAt *time.Time `json:"expires_at"`
}

func generateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// @Summary Create API key
// @Description Create a new API key for the authenticated user
// @Tags API Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateAPIKeyRequest true "API key creation request"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/api-keys [post]
func (h *APIKeyHandler) createAPIKey(c *gin.Context) {
	var req CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	currentUser := user.(*models.User)

	key, err := generateAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate API key"})
		return
	}

	apiKey := &models.APIKey{
		UserNumber: currentUser.Number,
		Number:     commons.UUIDGenerator(),
		Name:       req.Name,
		Key:        key,
		ExpiresAt:  req.ExpiresAt,
	}

	if err := h.controller.CreateAPIKey(apiKey); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create API key"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "API key created successfully",
		"key":     key,
		"id":      apiKey.Number,
	})
}

// @Summary List API keys
// @Description Get a list of all API keys for the authenticated user
// @Tags API Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} APIKeyResponse
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/api-keys [get]
func (h *APIKeyHandler) listAPIKeys(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	currentUser := user.(*models.User)

	keys, err := h.controller.GetAPIKeys(currentUser.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch API keys"})
		return
	}

	type APIKeyResponse struct {
		Number     string     `json:"number"`
		Name       string     `json:"name"`
		LastUsedAt *time.Time `json:"last_used_at"`
		ExpiresAt  *time.Time `json:"expires_at"`
		CreatedAt  time.Time  `json:"created_at"`
	}

	response := make([]APIKeyResponse, len(keys))
	for i, key := range keys {
		response[i] = APIKeyResponse{
			Number:     key.Number,
			Name:       key.Name,
			LastUsedAt: key.LastUsedAt,
			ExpiresAt:  key.ExpiresAt,
			CreatedAt:  key.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Delete API key
// @Description Delete an API key by its ID
// @Tags API Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "API Key ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/api-keys/{id} [delete]
func (h *APIKeyHandler) deleteAPIKey(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	currentUser := user.(*models.User)

	keyNumber, err := uuid.Parse(c.Param("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid API key number"})
		return
	}

	if err := h.controller.DeleteAPIKey(currentUser.Number, keyNumber); err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": "Failed to delete API key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "API key deleted successfully"})
}
