package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.New("database host is required")
	}
	if c.User == "" {
		return errors.New("database user is required")
	}
	if c.Password == "" {
		return errors.New("database password is required")
	}
	if c.DBName == "" {
		return errors.New("database name is required")
	}
	if c.Port == "" {
		return errors.New("database port is required")
	}
	return nil
}

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config *Config) (*Database, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	dsn := "host=" + config.Host + " user=" + config.User +
		" password=" + config.Password + " dbname=" + config.DBName +
		" port=" + config.Port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Enable UUID extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.APIKey{},
		&models.UserPassword{},
		&models.PasswordReset{},
		&models.Region{},
		&models.District{},
		&models.County{},
		&models.SubCounty{},
		&models.Parish{},
		&models.Village{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Database{DB: db}, nil
}
