package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
)

type CountyHandler struct {
	db *database.Database
}

func NewCountyHandler(db *database.Database) *CountyHandler {
	return &CountyHandler{
		db: db,
	}
}

func (h *CountyHandler) RegisterRoutes(r *gin.RouterGroup) {
	counties := r.Group("/counties")
	{
		counties.GET("", h.handleAllCounties)
		counties.POST("", h.createCounty)
		counties.GET("/:id", h.handleGetCounty)
		counties.PUT("/:id", h.updateCounty)
		counties.DELETE("/:id", h.deleteCounty)
		counties.GET("/:id/subcounties", h.getSubCounties)
	}
}

func (h *CountyHandler) handleAllCounties(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var counties []models.County
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, counties)
}

func (h *CountyHandler) createCounty(c *gin.Context) {
	var payload models.County
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCounty models.County
	if err := h.db.DB.Where("name = ? AND district_number = ?", payload.Name, payload.DistrictNumber).First(&existingCounty).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "County with this name already exists in this district"})
		return
	}

	county := models.County{
		Number:         commons.UUIDGenerator(),
		Name:           payload.Name,
		DistrictNumber: payload.DistrictNumber,
	}

	if err := h.db.DB.Create(&county).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "County created successfully",
	})
}

func (h *CountyHandler) handleGetCounty(c *gin.Context) {
	id := c.Param("id")

	var county models.County
	if err := h.db.DB.First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	c.JSON(http.StatusOK, county)
}

func (h *CountyHandler) updateCounty(c *gin.Context) {
	id := c.Param("id")

	var county models.County
	if err := h.db.DB.First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	var payload models.County
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCounty models.County
	if err := h.db.DB.Where("name = ? AND district_number = ? AND number != ?",
		payload.Name, payload.DistrictNumber, id).First(&existingCounty).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "County with this name already exists in this district"})
		return
	}

	county.Name = payload.Name
	county.DistrictNumber = payload.DistrictNumber

	if err := h.db.DB.Save(&county).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "County updated successfully",
	})
}

func (h *CountyHandler) deleteCounty(c *gin.Context) {
	id := c.Param("id")

	var county models.County
	if err := h.db.DB.First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	if err := h.db.DB.Delete(&county).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "County deleted successfully",
	})
}

func (h *CountyHandler) getSubCounties(c *gin.Context) {
	id := c.Param("id")

	var county models.County
	if err := h.db.DB.Preload("SubCounties").First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	c.JSON(http.StatusOK, county.SubCounties)
}
