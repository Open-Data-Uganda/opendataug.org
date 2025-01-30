package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uganda-data/commons"
	"github.com/uganda-data/database"
	"github.com/uganda-data/models"
)

type RegionHandler struct {
	db *database.Database
}

func NewRegionHandler(db *database.Database) *RegionHandler {
	return &RegionHandler{
		db: db,
	}
}

func (h *RegionHandler) RegisterRoutes(r *gin.RouterGroup) {
	regions := r.Group("/regions")
	{
		regions.GET("", h.handleAllRegions)
		regions.GET("/:id", h.handleGetRegion)
		regions.POST("", h.createRegion)
		regions.PUT("/:id", h.updateRegion)
		regions.DELETE("/:id", h.deleteRegion)
		regions.GET("/:id/districts", h.getDistricts)
	}
}

// @Summary Get all regions
// @Description Get a paginated list of all regions
// @Tags regions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.RegionResponse
// @Failure 500 {object} gin.H
// @Router /v1/regions [get]
func (h *RegionHandler) handleAllRegions(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var regions []models.Region
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&regions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []models.RegionResponse
	for _, region := range regions {
		response = append(response, models.RegionResponse{
			Number: region.Number,
			Name:   region.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get a region by ID
// @Description Get detailed information about a specific region
// @Tags regions
// @Accept json
// @Produce json
// @Param id path string true "Region ID"
// @Success 200 {object} models.RegionResponse
// @Failure 404 {object} gin.H
// @Router /v1/regions/{id} [get]
func (h *RegionHandler) handleGetRegion(c *gin.Context) {
	id := c.Param("id")

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
		return
	}

	response := models.RegionResponse{
		Number: region.Number,
		Name:   region.Name,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Create a new region
// @Description Create a new region with the provided information
// @Tags regions
// @Accept json
// @Produce json
// @Param region body models.Region true "Region information"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/regions [post]
func (h *RegionHandler) createRegion(c *gin.Context) {
	var payload models.Region
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingRegion models.Region
	if err := h.db.DB.Where("name = ?", payload.Name).First(&existingRegion).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region with this name already exists"})
		return
	}

	region := models.Region{
		Number: commons.UUIDGenerator(),
		Name:   payload.Name,
	}

	if err := h.db.DB.Create(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Region created successfully",
	})
}

// @Summary Update a region
// @Description Update an existing region's information
// @Tags regions
// @Accept json
// @Produce json
// @Param id path string true "Region ID"
// @Param region body models.Region true "Updated region information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/regions/{id} [put]
func (h *RegionHandler) updateRegion(c *gin.Context) {
	id := c.Param("id")

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
		return
	}

	var payload models.Region
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if another region already has this name
	var existingRegion models.Region
	if err := h.db.DB.Where("name = ? AND number != ?", payload.Name, id).First(&existingRegion).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region with this name already exists"})
		return
	}

	region.Name = payload.Name

	if err := h.db.DB.Save(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Region updated successfully",
	})
}

// @Summary Delete a region
// @Description Delete an existing region
// @Tags regions
// @Accept json
// @Produce json
// @Param id path string true "Region ID"
// @Success 200 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/regions/{id} [delete]
func (h *RegionHandler) deleteRegion(c *gin.Context) {
	id := c.Param("id")

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
		return
	}

	if err := h.db.DB.Delete(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Region deleted successfully",
	})
}

// @Summary Get districts in a region
// @Description Get all districts belonging to a specific region
// @Tags regions
// @Accept json
// @Produce json
// @Param id path string true "Region ID"
// @Success 200 {array} models.District
// @Failure 404 {object} gin.H
// @Router /v1/regions/{id}/districts [get]
func (h *RegionHandler) getDistricts(c *gin.Context) {
	id := c.Param("id")

	var region models.Region
	if err := h.db.DB.Preload("Districts").First(&region, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
		return
	}

	c.JSON(http.StatusOK, region.Districts)
}
