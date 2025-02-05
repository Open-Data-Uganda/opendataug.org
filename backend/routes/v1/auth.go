package v1

import (
	"net/http"
	"os"
	"strings"
	"time"

	"encoding/base64"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/controllers"
	"opendataug.org/database"
	"opendataug.org/models"
	"opendataug.org/services"
)

type AuthHandler struct {
	db             *database.Database
	userController *controllers.UserController
	jwtService     *services.JWTService
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	jwtService := services.NewJWTService()
	return &AuthHandler{
		db:             db,
		userController: controllers.NewUserController(db, jwtService),
		jwtService:     jwtService,
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
		auth.POST("/register", h.RegisterUser)

		protected := auth.Group("")
		protected.Use(h.TokenAuthMiddleware())
		{
			protected.POST("/logout", h.LogoutUser)
			protected.GET("/profile", h.Profile)
		}
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload models.SignInRequest
	if err := c.ShouldBindJSON(&payload); err != nil || payload.Validate() != nil {
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

	claims, err := h.jwtService.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token"})
		return
	}

	userNumber, ok := claims["user_number"].(string)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
		return
	}

	var user models.User
	if err := h.db.DB.Where("number = ?", userNumber).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
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
		"user_number":  userNumber,
		"role":         user.Role,
	})
}

func (h *AuthHandler) setAuthCookies(c *gin.Context, tokens *services.TokenDetails) {
	if tokens == nil || tokens.AccessToken == nil || tokens.RefreshToken == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token details"})
		return
	}

	secureValue := true
	domainValue := os.Getenv("BASE_URL")
	sameSite := http.SameSiteStrictMode

	if os.Getenv("ENVIRONMENT") == "dev" {
		secureValue = false
		domainValue = ""
		sameSite = http.SameSiteLaxMode
	}

	c.SetSameSite(sameSite)

	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    *tokens.AccessToken,
		Path:     "/",
		Domain:   domainValue,
		MaxAge:   int(services.AccessTokenDuration.Seconds()),
		Secure:   secureValue,
		HttpOnly: true,
		SameSite: sameSite,
	}
	http.SetCookie(c.Writer, accessCookie)

	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    *tokens.RefreshToken,
		Path:     "/",
		Domain:   domainValue,
		MaxAge:   int(services.RefreshTokenDuration.Seconds()),
		Secure:   secureValue,
		HttpOnly: true,
		SameSite: sameSite,
	}
	http.SetCookie(c.Writer, refreshCookie)

	loggedInCookie := &http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		Domain:   domainValue,
		MaxAge:   int(services.AccessTokenDuration.Seconds()),
		Secure:   secureValue,
		HttpOnly: false,
		SameSite: sameSite,
	}
	http.SetCookie(c.Writer, loggedInCookie)
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

	claims, err := h.jwtService.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "err": err.Error()})
		c.Abort()
		return
	}

	userNumber, ok := claims["user_number"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
		c.Abort()
		return
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "password_reset" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token type"})
		return
	}

	if err := h.userController.SetNewPassword(tokenString, userNumber, payload.Password, payload.ConfirmPassword); err != nil {
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
	domainValue := os.Getenv("BASE_URL")
	if os.Getenv("ENVIRONMENT") == "dev" {
		domainValue = ""
	}

	for _, cookieName := range []string{"access_token", "refresh_token", "logged_in"} {
		cookie := &http.Cookie{
			Name:     cookieName,
			Value:    "",
			Path:     "/",
			Domain:   domainValue,
			MaxAge:   -1,
			Secure:   os.Getenv("ENVIRONMENT") != "dev",
			HttpOnly: cookieName != "logged_in",
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(c.Writer, cookie)
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var payload models.SignUpInput
	if err := c.ShouldBindJSON(&payload); err != nil || payload.Validate() != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	payload.Prepare()
	if err := payload.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Failed to process input"})
		return
	}

	userNumber := commons.UUIDGenerator()

	if err := checkmail.ValidateFormat(payload.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email address"})
		return
	}

	emailExists, _ := h.userController.CheckEmailExists(strings.ToLower(payload.Email))
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is already in use"})
		return
	}

	user := models.User{
		Number:    userNumber,
		FirstName: payload.FirstName,
		OtherName: payload.OtherName,
		Email:     strings.ToLower(payload.Email),
		Role:      map[bool]string{true: "ADMIN", false: "USER"}[payload.Role == "ADMIN"],
		Status:    "INACTIVE",
		IsAdmin:   payload.Role == "ADMIN",
	}

	emailExists, _ = h.userController.CheckEmailExists(user.Email)
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is already in use"})
		return
	}

	tx := h.db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to register user"})
		return
	}

	baseURL := os.Getenv("BASE_URL")

	claims := jwt.MapClaims{
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
		"type":        "password_reset",
		"user_number": userNumber,
		"token_uuid":  commons.UUIDGenerator(),
		"iat":         time.Now().UTC().Unix(),
		"nbf":         time.Now().UTC().Unix(),
		"iss":         baseURL,
		"aud":         baseURL,
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to decode private key"})
		return
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse private key"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	userToken, err := token.SignedString(privateKey)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate reset token"})
		return
	}

	saveUserPasswordToken := models.PasswordReset{
		Number:     commons.UUIDGenerator(),
		UserNumber: userNumber,
		Token:      userToken,
		Status:     "ACTIVE",
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

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := h.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		userNumber, ok := claims["user_number"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
			c.Abort()
			return
		}

		var user models.User
		if err := h.db.DB.Where("number = ?", userNumber).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "usernumber": user.Name})
			c.Abort()
			return
		}

		c.Set("user", &user)
		c.Next()
	}
}

func (h *AuthHandler) APIAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("x-api-key")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No API Key provided"})
			c.Abort()
			return
		}

		var apiKeyModel models.APIKey
		if err := h.db.DB.Where("key = ? AND is_active = ?", apiKey, true).First(&apiKeyModel).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
			c.Abort()
			return
		}

		if apiKeyModel.ExpiresAt != nil && apiKeyModel.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "API key has expired"})
			c.Abort()
			return
		}

		var user models.User
		if err := h.db.DB.Where("number = ?", apiKeyModel.UserNumber).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			c.Abort()
			return
		}

		updates := map[string]interface{}{
			"last_used_at": time.Now(),
			"usage_count":  gorm.Expr("usage_count + ?", 1),
		}

		h.db.DB.Model(&apiKeyModel).Updates(updates)

		c.Set("api_key", &apiKeyModel)
		c.Set("user", &user)
		c.Next()
	}
}
