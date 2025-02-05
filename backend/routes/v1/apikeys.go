package v1

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/controllers"
	"opendataug.org/database"
	customerrors "opendataug.org/errors"
	"opendataug.org/models"
)

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type APIKeyHandler struct {
	controller *controllers.APIKeyController
}

func NewAPIKeyHandler(db *database.Database) *APIKeyHandler {
	return &APIKeyHandler{
		controller: controllers.NewAPIKeyController(db),
	}
}

func (h *APIKeyHandler) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	keys := r.Group("/api-keys")
	keys.Use(authHandler.TokenAuthMiddleware())
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

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateAPIKey() (string, error) {
	randomStr := generateRandomString(8)
	return "opu_" + commons.UUIDGenerator() + randomStr, nil
}

func (h *APIKeyHandler) createAPIKey(c *gin.Context) {
	var payload CreateAPIKeyRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError(err.Error(), nil))
		return
	}

	if payload.Name == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("API Key name is required"))
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("User not found in context"))
		return
	}
	currentUser := user.(*models.User)

	exists, err := h.controller.APIKeyNameExists(currentUser.Number, payload.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewInternalError("Failed to check API key name"))
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("API key with this name already exists"))
		return
	}

	key, err := generateAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewInternalError("Failed to generate API key"))
		return
	}

	apiKey := &models.APIKey{
		UserNumber: currentUser.Number,
		Number:     commons.UUIDGenerator(),
		Name:       payload.Name,
		Key:        key,
		ExpiresAt:  payload.ExpiresAt,
	}

	if err := h.controller.CreateAPIKey(apiKey); err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewInternalError("Failed to create API key"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "API Key created successfully", "key": apiKey.Key})
}

func (h *APIKeyHandler) listAPIKeys(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("User not found in context"))
		return
	}
	currentUser := user.(*models.User)

	keys, err := h.controller.GetAPIKeys(currentUser.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewInternalError("Failed to fetch API keys"))
		return
	}

	response := make([]models.APIKeyResponse, len(keys))
	for i, key := range keys {
		response[i] = models.APIKeyResponse{
			ID:        key.Number,
			Name:      key.Name,
			Key:       key.Key,
			CreatedAt: key.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *APIKeyHandler) deleteAPIKey(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("User not found in context"))
		return
	}
	currentUser := user.(*models.User)

	keyNumber := c.Param("id")
	if keyNumber == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("API key number is missing"))
		return
	}

	keyNumber = commons.Sanitize(keyNumber)

	if err := h.controller.DeleteAPIKey(currentUser.Number, keyNumber); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("API key not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, customerrors.NewInternalError("Failed to delete API key"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "API key deleted successfully"})
}
