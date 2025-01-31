package v1

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"opendataug.org/commons"
	"opendataug.org/controllers"
	"opendataug.org/database"
	"opendataug.org/models"
	"opendataug.org/services"
)

type AuthHandler struct {
	db             *database.Database
	userController *controllers.UserController
}

func NewAuthHandler(db *database.Database) *AuthHandler {

	return &AuthHandler{
		db:             db,
		userController: controllers.NewUserController(db),
	}
}

func (h *AuthHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.RefreshAccessToken)
		auth.POST("/reset-password", h.ResetPassword)
		auth.POST("/set-password", h.SetPassword)
		auth.POST("/logout", h.LogoutUser)
		auth.GET("/profile", h.Profile)
		auth.POST("/register", h.RegisterUser)
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload models.SignInRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	payload.Prepare()
	if err := payload.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	user, err := h.userController.AuthenticateUser(payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response, tokens, err := h.userController.CreateLoginSession(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	h.setAuthCookies(c, tokens)
	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) RefreshAccessToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing refresh token"})
		return
	}

	newTokens, err := h.userController.RefreshUserSession(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	h.setAuthCookies(c, newTokens)
	c.JSON(http.StatusOK, gin.H{
		"access_token": newTokens.AccessToken,
		"expires_in":   *newTokens.AccessTokenExpiresIn,
	})
}

func (h *AuthHandler) setAuthCookies(c *gin.Context, tokens *commons.TokenDetails) {
	if tokens == nil || tokens.AccessToken == nil || tokens.RefreshToken == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token details"})
		return
	}

	secureValue := true
	domainValue := os.Getenv("BASE_URL")
	if os.Getenv("ENVIRONMENT") == "dev" {
		secureValue = false
		domainValue = ""
	}

	c.SetCookie(
		"access_token",
		*tokens.AccessToken,
		int(services.AccessTokenDuration.Seconds()),
		"/",
		domainValue,
		secureValue,
		true,
	)

	c.SetCookie(
		"refresh_token",
		*tokens.RefreshToken,
		int(services.RefreshTokenDuration.Seconds()),
		"/",
		domainValue,
		secureValue,
		true,
	)

	c.SetCookie(
		"logged_in",
		"true",
		int(services.AccessTokenDuration.Seconds()),
		"/",
		domainValue,
		secureValue,
		false,
	)

	c.SetSameSite(http.SameSiteStrictMode)
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var payload controllers.ResetPasswordInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	userToken, err := h.userController.InitiatePasswordReset(payload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailService := services.Info{
		Email:       payload.Email,
		Token:       userToken,
		MailType:    "Password reset - RB Buildings CRM",
		UserName:    "",
		CurrentYear: time.Now().Year(),
		Type:        "password_reset",
	}

	if err := emailService.SendEmail(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to your email"})
}

func (h *AuthHandler) SetPassword(c *gin.Context) {
	var payload models.ResetPassword
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Reset token is required"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired reset token"})
		return
	}

	// Verify token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid reset token"})
		return
	}

	// Verify token type
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "password_reset" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token type"})
		return
	}

	// Get user number from claims
	userNumber, ok := claims["user_number"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
		return
	}

	if err := h.userController.SetNewPassword(userNumber, payload.Password, payload.ConfirmPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your account is now active. Log in to your account"})
}

func (h *AuthHandler) LogoutUser(c *gin.Context) {
	refreshToken, _ := c.Cookie("refresh_token")
	if err := h.userController.InvalidateSession(refreshToken); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	h.clearAuthCookies(c)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	user, err := commons.GetUserFromHeader(c, h.db.DB)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	profile, err := h.userController.GetUserProfile(user.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

func (h *AuthHandler) clearAuthCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.SetCookie("logged_in", "", -1, "/", "", false, true)
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var payload models.SignUpInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	payload.Prepare()
	if err := payload.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	userNumber := commons.UUIDGenerator()
	userStatus := "INACTIVE"

	if err := checkmail.ValidateFormat(payload.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email address"})
		return
	}

	var (
		adminRole bool
		role      string
	)
	if payload.Role == "ADMIN" {
		adminRole = true
		role = "ADMIN"
	} else {
		adminRole = false
		role = "USER"
	}

	user := models.User{
		Number:    userNumber,
		FirstName: payload.FirstName,
		OtherName: payload.OtherName,
		Email:     strings.ToLower(payload.Email),
		Role:      role,
		Status:    userStatus,
		IsAdmin:   adminRole,
	}

	emailExists, _ := h.userController.CheckEmailExists(user.Email)
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is already in use"})
		return
	}

	tx := h.db.DB.Begin()

	if err := tx.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to register user", "err": err.Error()})
		return
	}

	claims := jwt.MapClaims{
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
		"type":        "password_reset",
		"user_number": userNumber,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	userToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate reset token"})
		return
	}

	saveUserPasswordToken := models.PasswordReset{
		Number:     commons.UUIDGenerator(),
		UserNumber: userNumber,
		Token:      userToken,
		Status:     userStatus,
	}

	if err := tx.Create(&saveUserPasswordToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to save user password"})
		return
	}

	emailService := services.Info{
		Email:       user.Email,
		Token:       userToken,
		MailType:    "Uganda Open Data - Registration",
		UserName:    user.FirstName,
		CurrentYear: time.Now().Year(),
		Type:        "registration",
	}

	if err := emailService.SendEmail(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to send email"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered sucessfully"})
}
