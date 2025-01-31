package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type SubRegionHandler struct {
	DB *gorm.DB
}

func NewSubRegionHandler(db *gorm.DB) *SubRegionHandler {
	return &SubRegionHandler{DB: db}
}

func (sc *SubRegionHandler) GetSubRegions(c *gin.Context) {
	var subregions []models.SubRegion
	result := sc.DB.Find(&subregions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch subregions",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subregions,
	})
}

func (sc *SubRegionHandler) GetSubRegion(c *gin.Context) {
	id := c.Param("id")
	var subregion models.SubRegion

	result := sc.DB.First(&subregion, "number = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subregion not found",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subregion,
	})
}

func (sc *SubRegionHandler) CreateSubRegion(c *gin.Context) {
	var input models.SubRegion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result := sc.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create subregion",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": input,
	})
}

func (sc *SubRegionHandler) UpdateSubRegion(c *gin.Context) {
	id := c.Param("id")
	var subregion models.SubRegion

	result := sc.DB.First(&subregion, "number = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subregion not found",
			"error":   result.Error.Error(),
		})
		return
	}

	var input models.SubRegion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result = sc.DB.Model(&subregion).Updates(input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update subregion",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subregion,
	})
}

func (sc *SubRegionHandler) DeleteSubRegion(c *gin.Context) {
	id := c.Param("id")
	var subregion models.SubRegion

	result := sc.DB.First(&subregion, "number = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subregion not found",
			"error":   result.Error.Error(),
		})
		return
	}

	result = sc.DB.Delete(&subregion)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete subregion",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Subregion deleted successfully",
	})
}

func (sc *SubRegionHandler) GetSubRegionsByRegion(c *gin.Context) {
	regionID := c.Param("regionId")
	var subregions []models.SubRegion

	result := sc.DB.Where("region_number = ?", regionID).Find(&subregions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch subregions",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subregions,
	})
}
