package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/models"
)

type ParishController struct {
	DB *gorm.DB
}

func NewParishController(db *gorm.DB) *ParishController {
	return &ParishController{DB: db}
}

func (pc *ParishController) GetParishes(c *gin.Context) {
	var parishes []models.Parish
	result := pc.DB.Find(&parishes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch parishes",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": parishes,
	})
}

func (pc *ParishController) GetParish(c *gin.Context) {
	id := c.Param("id")
	var parish models.Parish

	result := pc.DB.First(&parish, "number = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Parish not found",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": parish,
	})
}

func (pc *ParishController) CreateParish(c *gin.Context) {
	var input models.Parish
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result := pc.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create parish",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": input,
	})
}

func (pc *ParishController) UpdateParish(c *gin.Context) {
	id := c.Param("id")
	var parish models.Parish

	result := pc.DB.First(&parish, "number = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Parish not found",
			"error":   result.Error.Error(),
		})
		return
	}

	var input models.Parish
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	result = pc.DB.Model(&parish).Updates(input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update parish",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": parish,
	})
}

func (pc *ParishController) DeleteParish(c *gin.Context) {
	id := c.Param("id")
	var parish models.Parish

	result := pc.DB.First(&parish, "number = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Parish not found",
			"error":   result.Error.Error(),
		})
		return
	}

	result = pc.DB.Delete(&parish)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete parish",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Parish deleted successfully",
	})
}

func (pc *ParishController) GetParishesByDistrict(c *gin.Context) {
	districtID := c.Param("id")
	var parishes []models.Parish

	result := pc.DB.Where("district_number = ?", districtID).Find(&parishes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch parishes",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": parishes,
	})
}
