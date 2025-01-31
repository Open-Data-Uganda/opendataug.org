package commons

import (
	"errors"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"opendataug.org/models"
)

func UUIDGenerator() string {
	u := uuid.NewV4().String()
	return strings.ReplaceAll(u, "-", "")
}

func ComparePassword(hashedPassword, password string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		return false, err
	}
	if !match {
		return false, errors.New("password does not match")
	}

	return true, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return hashedPassword, nil
}

func GetUserFromHeader(c *gin.Context, db *gorm.DB) (*models.User, error) {
	userNumber := c.Request.Header["User-Number"][0]
	var user models.User

	if err := db.Where("number = ?", userNumber).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
