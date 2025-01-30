package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
)

type DistrictHandler struct {
	db *database.Database
}

func NewDistrictHandler(db *database.Database) *DistrictHandler {
	return &DistrictHandler{
		db: db,
	}
}

func (h *DistrictHandler) RegisterRoutes(r *gin.RouterGroup) {
	districts := r.Group("/districts")
	{
		districts.GET("", h.handleGetDistricts)
		districts.GET("/:number", h.handleGetDistrict)
		districts.POST("", h.handleCreateDistrict)
		districts.PUT("/:number", h.handleUpdateDistrict)
		districts.DELETE("/:number", h.handleDeleteDistrict)
		districts.GET("/:number/counties", h.handleGetDistrictCounties)
	}
}

func (h *DistrictHandler) handleGetDistricts(c *gin.Context) {
	var districts []models.District
	if err := h.db.DB.Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, districts)
}

func (h *DistrictHandler) handleGetDistrict(c *gin.Context) {
	number := c.Param("number")
	var district models.District
	if err := h.db.DB.Preload("Counties").First(&district, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}
	c.JSON(http.StatusOK, district)
}

func (h *DistrictHandler) handleCreateDistrict(c *gin.Context) {
	var payload models.District
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	district := models.District{
		Number:     commons.UUIDGenerator(),
		Name:       payload.Name,
		Region:     payload.Region,
		Size:       payload.Size,
		SizeUnits:  payload.SizeUnits,
		TownStatus: payload.TownStatus,
	}

	if err := h.db.DB.Create(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, district)
}

func (h *DistrictHandler) handleUpdateDistrict(c *gin.Context) {
	number := c.Param("number")
	var district models.District
	if err := h.db.DB.First(&district, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	if err := c.ShouldBindJSON(&district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB.Save(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, district)
}

func (h *DistrictHandler) handleDeleteDistrict(c *gin.Context) {
	number := c.Param("number")
	if err := h.db.DB.Delete(&models.District{}, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "District deleted successfully"})
}

func (h *DistrictHandler) handleGetDistrictCounties(c *gin.Context) {
	number := c.Param("number")
	var counties []models.County
	if err := h.db.DB.Where("district_number = ?", number).Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counties)
}
