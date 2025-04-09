package commons

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

type LoginResponse struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	UserNumber   string  `json:"user_number"`
	Role         string  `json:"role"`
	ExpiresIn    *int64  `json:"expires_in"`
}

type TokenDetails struct {
	AccessToken           *string
	RefreshToken          *string
	AccessTokenExpiresIn  *int64
	RefreshTokenExpiresIn *int64
}

func (s *AuthService) CreateLoginSession(user *models.User) (*LoginResponse, *TokenDetails, error) {
	tokenDetails, err := CreateToken(user.Number)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create token: %w", err)
	}

	response := &LoginResponse{
		RefreshToken: tokenDetails.RefreshToken,
		AccessToken:  tokenDetails.AccessToken,
		UserNumber:   user.Number,
		Role:         string(user.Role),
		ExpiresIn:    tokenDetails.AccessTokenExpiresIn,
	}

	return response, tokenDetails, nil
}

func (s *AuthService) RefreshUserSession(refreshToken string) (*TokenDetails, error) {
	return RefreshToken(refreshToken)
}

func (s *AuthService) InvalidateSession(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("missing refresh token")
	}

	_, err := ValidateToken(refreshToken, "refresh")
	if err != nil {
		return fmt.Errorf("invalid refresh token: %w", err)
	}

	return nil
}

func ValidateToken(tokenString string, tokenType string) (jwt.MapClaims, error) {
	publicKey := os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token public key: %w", err)
	}

	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse token public key: %w", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims["type"] != tokenType {
		return nil, fmt.Errorf("invalid token type")
	}

	return claims, nil
}

func RefreshToken(refreshToken string) (*TokenDetails, error) {
	claims, err := ValidateToken(refreshToken, "refresh")
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	userNumber, ok := claims["user_number"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return CreateToken(userNumber)
}

func CreateToken(userNumber string) (*TokenDetails, error) {
	td := &TokenDetails{}
	now := time.Now()
	privateKey := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		log.Println(privateKey)
		return nil, fmt.Errorf("invalid token")
	}

	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("create: parse token private key: %w", err)
	}

	// Access token
	accessClaims := jwt.MapClaims{
		"user_number": userNumber,
		"type":        "access",
		"exp":         now.Add(time.Hour * 1).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(rsaPrivateKey)
	if err != nil {
		return nil, err
	}

	// Refresh token
	refreshClaims := jwt.MapClaims{
		"user_number": userNumber,
		"type":        "refresh",
		"exp":         now.Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(rsaPrivateKey)
	if err != nil {
		return nil, err
	}

	td.AccessToken = &signedAccessToken
	td.RefreshToken = &signedRefreshToken
	td.AccessTokenExpiresIn = new(int64)
	*td.AccessTokenExpiresIn = accessClaims["exp"].(int64)
	td.RefreshTokenExpiresIn = new(int64)
	*td.RefreshTokenExpiresIn = refreshClaims["exp"].(int64)

	return td, nil
}
