package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"opendataug.org/controllers"
	"opendataug.org/database"
	"opendataug.org/models"
)

type Provider struct {
	config         *oauth2.Config
	name           string
	userController *controllers.UserController
}

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func NewGithubProvider(db *database.Database, cfg Config) *Provider {
	return &Provider{
		name:           "github",
		userController: controllers.NewUserController(db),
		config: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		},
	}
}

func NewGoogleProvider(db *database.Database, cfg Config) *Provider {
	return &Provider{
		name:           "google",
		userController: controllers.NewUserController(db),
		config: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (p *Provider) HandleLogin(c *gin.Context) {
	url := p.config.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (p *Provider) HandleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := p.config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
		return
	}

	userInfo, err := p.getUserInfo(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
		return
	}

	user, err := p.userController.GetUserByAuthID(p.name, userInfo.ID)
	if err != nil {
		user = &models.User{
			Email:     userInfo.Email,
			Name:      userInfo.Name,
			Provider:  p.name,
			AuthID:    userInfo.ID,
			AvatarURL: userInfo.AvatarURL,
		}
		if err := p.userController.CreateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully authenticated",
		"user":    user,
	})
}
