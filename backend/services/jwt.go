package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"opendataug.org/commons"
)

type TokenDetails struct {
	AccessToken           *string
	RefreshToken          *string
	AccessTokenUUID       string
	RefreshTokenUUID      string
	UserNumber            string
	AccessTokenExpiresIn  *int64
	RefreshTokenExpiresIn *int64
}

const (
	AccessTokenDuration  = time.Minute * 15
	RefreshTokenDuration = time.Hour * 24 * 7
	MaxTokens            = 5
)

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) CreateToken(userNumber string, userRole string) (*TokenDetails, error) {
	privateKey := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	if privateKey == "" {
		return nil, fmt.Errorf("missing ACCESS_TOKEN_PRIVATE_KEY environment variable")
	}

	now := time.Now().UTC()
	td := &TokenDetails{
		AccessToken:           new(string),
		RefreshToken:          new(string),
		AccessTokenExpiresIn:  new(int64),
		RefreshTokenExpiresIn: new(int64),
		AccessTokenUUID:       commons.UUIDGenerator(),
		RefreshTokenUUID:      commons.UUIDGenerator(),
		UserNumber:            userNumber,
	}

	*td.AccessTokenExpiresIn = now.Add(AccessTokenDuration).Unix()
	*td.RefreshTokenExpiresIn = now.Add(RefreshTokenDuration).Unix()

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token private key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("create: parse token private key: %w", err)
	}

	baseURL := os.Getenv("BASE_URL")

	accessClaims := jwt.MapClaims{
		"user_number": userNumber,
		"token_uuid":  td.AccessTokenUUID,
		"exp":         td.AccessTokenExpiresIn,
		"iat":         now.Unix(),
		"nbf":         now.Unix(),
		"iss":         baseURL,
		"aud":         baseURL,
		"type":        "access",
		"user_role":   userRole,
	}

	refreshClaims := jwt.MapClaims{
		"user_number": userNumber,
		"token_uuid":  td.RefreshTokenUUID,
		"exp":         td.RefreshTokenExpiresIn,
		"iat":         now.Unix(),
		"nbf":         now.Unix(),
		"iss":         baseURL,
		"aud":         baseURL,
		"type":        "refresh",
		"user_role":   userRole,
	}

	*td.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims).SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("create: sign access token: %w", err)
	}

	*td.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims).SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("create: sign refresh token: %w", err)
	}

	return td, nil
}

func (s JWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	publicKey := os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")
	if publicKey == "" {
		return nil, fmt.Errorf("missing ACCESS_TOKEN_PUBLIC_KEY environment variable")
	}

	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validating token: %w", err)
	}

	baseURL := os.Getenv("BASE_URL")

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	if err := claims.Valid(); err != nil {
		return nil, fmt.Errorf("validate: invalid claims: %w", err)
	}

	// Validate required fields and their types
	requiredClaims := []string{"user_number", "token_uuid", "exp", "iat", "nbf", "iss", "aud", "type", "user_role"}
	for _, claim := range requiredClaims {
		if _, ok := claims[claim]; !ok {
			return nil, fmt.Errorf("validate: missing required claim: %s", claim)
		}
	}

	// Validate issuer and audience
	if iss, ok := claims["iss"].(string); !ok || iss != baseURL {
		return nil, fmt.Errorf("validate: invalid issuer")
	}

	if aud, ok := claims["aud"].(string); !ok || aud != baseURL {
		return nil, fmt.Errorf("validate: invalid audience")
	}

	return claims, nil
}

func (s *JWTService) RefreshToken(refreshToken string) (*TokenDetails, error) {
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Validate it's a refresh token
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
		return nil, fmt.Errorf("invalid token type: expected refresh token")
	}

	userNumber := fmt.Sprint(claims["user_number"])
	userRole := fmt.Sprint(claims["user_role"])
	newTokens, err := s.CreateToken(userNumber, userRole)
	if err != nil {
		return nil, fmt.Errorf("creating new tokens: %w", err)
	}

	return newTokens, nil
}

func (s *JWTService) ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

func (s *JWTService) GenerateTokenWithClaims(claims jwt.MapClaims, tx *gorm.DB) (string, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return "", fmt.Errorf("failed to decode private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}
