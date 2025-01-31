package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
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

func (h *RegionHandler) getDistricts(c *gin.Context) {
	id := c.Param("id")

	var region models.Region
	if err := h.db.DB.Preload("Districts").First(&region, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
		return
	}

	c.JSON(http.StatusOK, region.Districts)
}
