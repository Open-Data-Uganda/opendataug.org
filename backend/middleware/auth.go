package middleware

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uganda-data/controllers"
	"github.com/uganda-data/database"
)

func APIKeyAuth(db *database.Database) gin.HandlerFunc {
	apiKeyController := controllers.NewAPIKeyController(db)

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No API key provided"})
			c.Abort()
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		apiKey, err := apiKeyController.GetAPIKeyByNumber(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		if apiKey.ExpiresAt != nil && apiKey.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key has expired"})
			c.Abort()
			return
		}

		go func() {
			if err := apiKeyController.UpdateAPIKeyUsage(apiKey.Number); err != nil {
				slog.Error("failed to update API key usage",
					"error", err,
					"api_key_number", apiKey.Number,
				)
			}
		}()

		c.Set("api_key", apiKey)
		c.Next()
	}
}
