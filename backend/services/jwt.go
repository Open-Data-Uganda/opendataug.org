package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

// Constants for token configuration
const (
	AccessTokenDuration  = time.Minute * 15   // 15 minutes
	RefreshTokenDuration = time.Hour * 24 * 7 // 7 days
	MaxTokens            = 5                  // Maximum concurrent tokens per user
)

func CreateToken(userNumber string) (*TokenDetails, error) {
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

	accessClaims := jwt.MapClaims{
		"sub":        userNumber,
		"token_uuid": td.AccessTokenUUID,
		"exp":        td.AccessTokenExpiresIn,
		"iat":        now.Unix(),
		"nbf":        now.Unix(),
		"iss":        "opendataug.org",
		"aud":        "opendataug.org",
		"type":       "access",
	}

	refreshClaims := jwt.MapClaims{
		"sub":        userNumber,
		"token_uuid": td.RefreshTokenUUID,
		"exp":        td.RefreshTokenExpiresIn,
		"iat":        now.Unix(),
		"nbf":        now.Unix(),
		"iss":        "opendataug.org",
		"aud":        "opendataug.org",
		"type":       "refresh",
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

func ValidateToken(token string, tokenType string) (*TokenDetails, error) {
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

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validating token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	if err := claims.Valid(); err != nil {
		return nil, fmt.Errorf("validate: invalid claims: %w", err)
	}

	// Validate token type
	if claimType, ok := claims["type"].(string); !ok || claimType != tokenType {
		return nil, fmt.Errorf("validate: invalid token type")
	}

	// Validate required fields and their types
	requiredClaims := []string{"sub", "token_uuid", "exp", "iat", "nbf", "iss", "aud", "type"}
	for _, claim := range requiredClaims {
		if _, ok := claims[claim]; !ok {
			return nil, fmt.Errorf("validate: missing required claim: %s", claim)
		}
	}

	// Validate issuer and audience
	if iss, ok := claims["iss"].(string); !ok || iss != "opendataug.org" {
		return nil, fmt.Errorf("validate: invalid issuer")
	}

	if aud, ok := claims["aud"].(string); !ok || aud != "opendataug.org" {
		return nil, fmt.Errorf("validate: invalid audience")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("validate: invalid expiration claim type")
	}

	td := &TokenDetails{
		AccessTokenUUID:      fmt.Sprint(claims["token_uuid"]),
		UserNumber:           fmt.Sprint(claims["sub"]),
		AccessTokenExpiresIn: new(int64),
	}
	*td.AccessTokenExpiresIn = int64(exp)

	return td, nil
}

func RefreshToken(refreshToken string) (*TokenDetails, error) {
	// Validate the refresh token
	td, err := ValidateToken(refreshToken, "refresh")
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Create new tokens
	newTokens, err := CreateToken(td.UserNumber)
	if err != nil {
		return nil, fmt.Errorf("creating new tokens: %w", err)
	}

	return newTokens, nil
}
