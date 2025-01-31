package commons

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userNumber string) (*TokenDetails, error) {
	td := &TokenDetails{}
	now := time.Now()

	expiry := now.Add(15 * time.Minute).Unix()
	td.AccessTokenExpiresIn = &expiry
	td.RefreshTokenExpiresIn = &expiry

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_number": userNumber,
		"exp":         expiry,
		"type":        "access",
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_number": userNumber,
		"exp":         expiry,
		"type":        "refresh",
	})

	strAccessToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	td.AccessToken = &strAccessToken

	strRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	td.RefreshToken = &strRefreshToken

	return td, nil
}

func ValidateToken(tokenString, tokenType string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["type"] != tokenType {
			return nil, fmt.Errorf("invalid token type")
		}
		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func RefreshToken(refreshToken string) (*TokenDetails, error) {
	token, err := ValidateToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userNumber := claims["user_number"].(string)

	return CreateToken(userNumber)
}
