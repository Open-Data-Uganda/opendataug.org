package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type CountyController struct {
	DB *gorm.DB
}

func NewCountyController(db *gorm.DB) *CountyController {
	return &CountyController{DB: db}
}

func (cc *CountyController) GetCounties(c *gin.Context) {
	var counties []models.County
	if err := cc.DB.Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching counties"})
		return
	}
	c.JSON(http.StatusOK, counties)
}

func (cc *CountyController) GetCounty(c *gin.Context) {
	id := c.Param("id")
	var county models.County
	if err := cc.DB.Preload("SubCounties").First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}
	c.JSON(http.StatusOK, county)
}

func (cc *CountyController) CreateCounty(c *gin.Context) {
	var county models.County
	if err := c.ShouldBindJSON(&county); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Create(&county).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating county"})
		return
	}
	c.JSON(http.StatusCreated, county)
}

func (cc *CountyController) UpdateCounty(c *gin.Context) {
	id := c.Param("id")
	var county models.County
	if err := cc.DB.First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	if err := c.ShouldBindJSON(&county); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Save(&county).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating county"})
		return
	}
	c.JSON(http.StatusOK, county)
}

func (cc *CountyController) DeleteCounty(c *gin.Context) {
	id := c.Param("id")
	if err := cc.DB.Delete(&models.County{}, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting county"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "County deleted successfully"})
}

func (cc *CountyController) GetDistrictCounties(c *gin.Context) {
	districtNumber := c.Param("id")
	var counties []models.County
	if err := cc.DB.Where("district_number = ?", districtNumber).Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching district counties"})
		return
	}
	c.JSON(http.StatusOK, counties)
}
