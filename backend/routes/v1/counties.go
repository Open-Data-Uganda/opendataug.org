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
		counties.GET("/:number", h.handleGetCounty)
		counties.PUT("/:number", h.updateCounty)
		counties.DELETE("/:number", h.deleteCounty)
		counties.GET("/:number/subcounties", h.getSubCounties)
		counties.GET("/district/:number", h.getCountiesByDistrict)
	}
}

func (h *CountyHandler) toCountyResponse(county models.County) models.CountyResponse {
	return models.CountyResponse{
		Number:         county.Number,
		Name:           county.Name,
		DistrictNumber: county.DistrictNumber,
		DistrictName:   county.District.Name,
	}
}

func (h *CountyHandler) handleAllCounties(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var counties []models.County
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Preload("District").Limit(pagination.Limit).Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []models.CountyResponse
	for _, county := range counties {
		response = append(response, h.toCountyResponse(county))
	}

	c.JSON(http.StatusOK, response)
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
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid county number"})
		return
	}

	var county models.County
	if err := h.db.DB.Preload("District").First(&county, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	c.JSON(http.StatusOK, h.toCountyResponse(county))
}

func (h *CountyHandler) updateCounty(c *gin.Context) {
	id := c.Param("number")

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
	id := c.Param("number")

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
	number := c.Param("number")
	number = commons.Sanitize(number)
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid county number"})
		return
	}

	var county models.County
	if err := h.db.DB.Preload("SubCounties").First(&county, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	var response []models.SubCountyResponse
	for _, subCounty := range county.SubCounties {
		response = append(response, models.SubCountyResponse{
			Name:   subCounty.Name,
			Number: subCounty.Number,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *CountyHandler) getCountiesByDistrict(c *gin.Context) {
	districtNumber := c.Param("number")
	if districtNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "District number is required"})
		return
	}

	districtNumber = commons.Sanitize(districtNumber)

	var counties []models.County
	if err := h.db.DB.Where("district_number = ?", districtNumber).Preload("District").Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []models.CountyResponse
	for _, county := range counties {
		response = append(response, models.CountyResponse{
			Number:         county.Number,
			Name:           county.Name,
			DistrictNumber: county.DistrictNumber,
			DistrictName:   county.District.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}
