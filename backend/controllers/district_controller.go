package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type DistrictController struct {
	DB *gorm.DB
}

func NewDistrictController(db *gorm.DB) *DistrictController {
	return &DistrictController{DB: db}
}

func (dc *DistrictController) GetDistricts(c *gin.Context) {
	var districts []models.District
	if err := dc.DB.Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching districts"})
		return
	}
	c.JSON(http.StatusOK, districts)
}

func (dc *DistrictController) GetDistrict(c *gin.Context) {
	id := c.Param("id")
	var district models.District
	if err := dc.DB.Preload("Counties").First(&district, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}
	c.JSON(http.StatusOK, district)
}

func (dc *DistrictController) CreateDistrict(c *gin.Context) {
	var payload models.District
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	district := models.District{
		Number:       payload.Number,
		Name:         payload.Name,
		TownStatus:   payload.TownStatus,
		RegionNumber: payload.RegionNumber,
	}

	if err := dc.DB.Create(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating district"})
		return
	}
	c.JSON(http.StatusCreated, district)
}

func (dc *DistrictController) UpdateDistrict(c *gin.Context) {
	id := c.Param("id")
	var district models.District
	if err := dc.DB.First(&district, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	if err := c.ShouldBindJSON(&district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dc.DB.Save(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating district"})
		return
	}
	c.JSON(http.StatusOK, district)
}

func (dc *DistrictController) DeleteDistrict(c *gin.Context) {
	id := c.Param("id")
	if err := dc.DB.Delete(&models.District{}, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting district"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "District deleted successfully"})
}
