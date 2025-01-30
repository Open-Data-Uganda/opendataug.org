package v1

import (
	"os"

	"github.com/gin-gonic/gin"
	"opendataug.org/database"
	"opendataug.org/pkg/oauth"
)

type AuthHandler struct {
	githubProvider *oauth.Provider
	googleProvider *oauth.Provider
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	baseURL := os.Getenv("BASE_URL")

	githubCfg := oauth.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  baseURL + "/api/v1/auth/github/callback",
	}

	googleCfg := oauth.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  baseURL + "/api/v1/auth/google/callback",
	}

	return &AuthHandler{
		githubProvider: oauth.NewGithubProvider(db, githubCfg),
		googleProvider: oauth.NewGoogleProvider(db, googleCfg),
	}
}

func (h *AuthHandler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.GET("/github", h.githubProvider.HandleLogin)
		auth.GET("/github/callback", h.githubProvider.HandleCallback)
		auth.GET("/google", h.googleProvider.HandleLogin)
		auth.GET("/google/callback", h.googleProvider.HandleCallback)
	}
}
