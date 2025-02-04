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

func (h *RegionHandler) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	regions := r.Group("/regions")
	{
		apiProtected := regions.Group("")
		apiProtected.Use(authHandler.APIAuthMiddleware())
		{
			apiProtected.GET("", h.handleAllRegions)
			apiProtected.GET("/:id", h.handleGetRegion)
			regions.GET("/:id/districts", h.getDistricts)
		}

		private := regions.Group("")
		private.Use(authHandler.TokenAuthMiddleware())
		{
			private.POST("", h.createRegion)
			private.PUT("/:id", h.updateRegion)
			private.DELETE("/:id", h.deleteRegion)
		}
	}
}

func (h *RegionHandler) handleAllRegions(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var regions []models.Region
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&regions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var response []models.RegionResponse
	for _, region := range regions {
		response = append(response, models.RegionResponse{
			ID:   region.Number,
			Name: region.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *RegionHandler) handleGetRegion(c *gin.Context) {
	number := c.Param("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid region id"})
		return
	}
	number = commons.Sanitize(number)

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Region not found"})
		return
	}

	response := models.RegionResponse{
		ID:   region.Number,
		Name: region.Name,
	}

	c.JSON(http.StatusOK, response)
}

func (h *RegionHandler) createRegion(c *gin.Context) {
	var payload models.Region
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var existingRegion models.Region
	if err := h.db.DB.Where("name = ?", payload.Name).First(&existingRegion).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Region with this name already exists"})
		return
	}

	region := models.Region{
		Number: commons.UUIDGenerator(),
		Name:   payload.Name,
	}

	if err := h.db.DB.Create(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Region created successfully",
	})
}

func (h *RegionHandler) updateRegion(c *gin.Context) {
	number := c.Param("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Region id is required"})
		return
	}

	number = commons.Sanitize(number)

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Region not found"})
		return
	}

	var payload models.Region
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var existingRegion models.Region
	if err := h.db.DB.Where("name = ? AND number != ?", payload.Name, number).First(&existingRegion).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Region with this name already exists"})
		return
	}

	region.Name = payload.Name

	if err := h.db.DB.Save(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Region updated successfully",
	})
}

func (h *RegionHandler) deleteRegion(c *gin.Context) {
	number := c.Param("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Region id is required"})
		return
	}

	number = commons.Sanitize(number)

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Region not found"})
		return
	}

	if err := h.db.DB.Delete(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Region deleted successfully"})
}

func (h *RegionHandler) getDistricts(c *gin.Context) {
	number := c.Param("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Region number is required"})
		return
	}
	number = commons.Sanitize(number)

	var region models.Region
	if err := h.db.DB.Preload("Districts").First(&region, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Region not found"})
		return
	}

	var districts []models.DistrictResponse
	for _, district := range region.Districts {
		districts = append(districts, models.DistrictResponse{
			ID:         district.Number,
			Name:       district.Name,
			Size:       district.Size,
			TownStatus: district.TownStatus,
			RegionID:   district.RegionNumber,
			RegionName: district.Region.Name,
		})
	}

	c.JSON(http.StatusOK, districts)
}
