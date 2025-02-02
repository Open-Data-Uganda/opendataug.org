package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type RegionHandler struct {
	DB *gorm.DB
}

func NewRegionHandler(db *gorm.DB) *RegionHandler {
	return &RegionHandler{DB: db}
}

// CreateRegion creates a new region
func (h *RegionHandler) CreateRegion(c *gin.Context) {
	var input models.Region
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result := h.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create region",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": input,
	})
}

// GetRegion retrieves a region by ID
func (h *RegionHandler) GetRegion(c *gin.Context) {
	var region models.Region
	result := h.DB.First(&region, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Region not found",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": region,
	})
}

// ListRegions retrieves all regions
func (h *RegionHandler) ListRegions(c *gin.Context) {
	var regions []models.Region
	result := h.DB.Find(&regions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch regions",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": regions,
	})
}

// UpdateRegion updates a region by ID
func (h *RegionHandler) UpdateRegion(c *gin.Context) {
	var region models.Region
	if result := h.DB.First(&region, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Region not found",
			"error":   result.Error.Error(),
		})
		return
	}

	var input models.Region
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result := h.DB.Model(&region).Updates(input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update region",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": region,
	})
}

// DeleteRegion deletes a region by ID
func (h *RegionHandler) DeleteRegion(c *gin.Context) {
	var region models.Region
	if result := h.DB.First(&region, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Region not found",
			"error":   result.Error.Error(),
		})
		return
	}

	if result := h.DB.Delete(&region); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete region",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Region deleted successfully",
	})
}

// CreateSubRegion creates a new subregion
func (h *RegionHandler) CreateSubRegion(c *gin.Context) {
	var input models.SubRegion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result := h.DB.Create(&input)
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
