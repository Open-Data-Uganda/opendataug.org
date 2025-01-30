package controllers

import (
	"opendataug.org/database"
	"opendataug.org/models"
)

type UserController struct {
	db *database.Database
}

func NewUserController(db *database.Database) *UserController {
	return &UserController{db: db}
}

func (c *UserController) CreateUser(user *models.User) error {
	return c.db.DB.Create(user).Error
}

func (c *UserController) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := c.db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *UserController) GetUserByAuthID(provider, authID string) (*models.User, error) {
	var user models.User
	result := c.db.DB.Where("provider = ? AND auth_number = ?", provider, authID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *UserController) GetUserByNumber(number string) (*models.User, error) {
	var user models.User
	result := c.db.DB.Where("number = ?", number).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
