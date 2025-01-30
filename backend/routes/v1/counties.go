package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uganda-data/commons"
	"github.com/uganda-data/database"
	"github.com/uganda-data/models"
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

// @Summary Get all counties
// @Description Get a paginated list of all counties
// @Tags Counties
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.County
// @Failure 500 {object} gin.H
// @Router /v1/counties [get]
func (h *CountyHandler) handleAllCounties(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var counties []models.County
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, counties)
}

// @Summary Create a new county
// @Description Create a new county with the provided details
// @Tags Counties
// @Accept json
// @Produce json
// @Param county body models.County true "County details"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/counties [post]
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

// @Summary Get a county by ID
// @Description Get detailed information about a specific county
// @Tags Counties
// @Accept json
// @Produce json
// @Param id path string true "County ID"
// @Success 200 {object} models.County
// @Failure 404 {object} gin.H
// @Router /v1/counties/{id} [get]
func (h *CountyHandler) handleGetCounty(c *gin.Context) {
	id := c.Param("id")

	var county models.County
	if err := h.db.DB.First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	c.JSON(http.StatusOK, county)
}

// @Summary Update a county
// @Description Update an existing county's information
// @Tags Counties
// @Accept json
// @Produce json
// @Param id path string true "County ID"
// @Param county body models.County true "Updated county details"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/counties/{id} [put]
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

// @Summary Delete a county
// @Description Delete a county by its ID
// @Tags Counties
// @Accept json
// @Produce json
// @Param id path string true "County ID"
// @Success 200 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/counties/{id} [delete]
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

// @Summary Get subcounties of a county
// @Description Get all subcounties belonging to a specific county
// @Tags Counties
// @Accept json
// @Produce json
// @Param id path string true "County ID"
// @Success 200 {array} models.SubCounty
// @Failure 404 {object} gin.H
// @Router /v1/counties/{id}/subcounties [get]
func (h *CountyHandler) getSubCounties(c *gin.Context) {
	id := c.Param("id")

	var county models.County
	if err := h.db.DB.Preload("SubCounties").First(&county, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	c.JSON(http.StatusOK, county.SubCounties)
}
