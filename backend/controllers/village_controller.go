package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/errors"
	"opendataug.org/models"
)

type VillageController struct {
	db *database.Database
}

func NewVillageController(db *database.Database) *VillageController {
	return &VillageController{
		db: db,
	}
}

func (c *VillageController) CreateVillage(ctx *gin.Context) error {
	var payload models.Village
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		return errors.NewValidationError("Failed to parse request body")
	}

	var parish models.Parish
	if err := c.db.DB.First(&parish, "number = ?", payload.ParishNumber).Error; err != nil {
		return errors.NewNotFoundError("Parish not found")
	}

	var existingVillage models.Village
	if err := c.db.DB.Where("name = ? AND parish_number = ?", payload.Name, payload.ParishNumber).
		First(&existingVillage).Error; err == nil {
		return errors.NewBadRequestError("Village with this name already exists in this parish")
	}

	village := models.Village{
		Number:       commons.UUIDGenerator(),
		Name:         payload.Name,
		ParishNumber: payload.ParishNumber,
	}

	if err := c.db.DB.Create(&village).Error; err != nil {
		return errors.NewDatabaseError("Database level error occurred")
	}

	return nil
}

func (c *VillageController) UpdateVillage(ctx *gin.Context) error {
	id := ctx.Param("id")

	var village models.Village
	if err := c.db.DB.First(&village, "number = ?", id).Error; err != nil {
		return errors.NewNotFoundError("Village not found")
	}

	var payload models.Village
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		return errors.NewValidationError("Failed to parse request body")
	}

	village.Name = payload.Name
	village.ParishNumber = payload.ParishNumber

	if err := c.db.DB.Save(&village).Error; err != nil {
		return errors.NewDatabaseError("Database level error occurred")
	}

	return nil
}

func (c *VillageController) DeleteVillage(ctx *gin.Context) error {
	id := ctx.Param("id")

	var village models.Village
	if err := c.db.DB.First(&village, "number = ?", id).Error; err != nil {
		return errors.NewNotFoundError("Village not found")
	}

	if err := c.db.DB.Delete(&village).Error; err != nil {
		return errors.NewDatabaseError("Database level error occurred")
	}

	return nil
}

func (c *VillageController) GetAllVillages(ctx *gin.Context) ([]models.Village, error) {
	pagination := commons.GetPaginationParams(ctx)

	var villages []models.Village
	if err := c.db.DB.Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&villages).Error; err != nil {
		return nil, errors.NewDatabaseError("Database level error occurred")
	}

	return villages, nil
}

func (c *VillageController) GetVillage(ctx *gin.Context) (*models.Village, error) {
	id := ctx.Param("id")
	if id == "" {
		return nil, errors.NewBadRequestError("Village ID is required")
	}

	var village models.Village
	if err := c.db.DB.First(&village, "number = ?", id).Error; err != nil {
		return nil, errors.NewNotFoundError("Village not found")
	}

	return &village, nil
}

func (c *VillageController) GetDB() *gorm.DB {
	return c.db.DB
}
