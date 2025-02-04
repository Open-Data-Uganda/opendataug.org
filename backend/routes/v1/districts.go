package v1

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
		districts.GET("", h.handleAllDistricts)
		districts.POST("", h.createDistrict)
		districts.GET("/:number", h.handleDistrictByNumber)
		districts.GET("/name/:name", h.handleDistrictByName)
		districts.DELETE("/:number", h.deleteDistrict)
	}
}

func (h *DistrictHandler) createDistrict(c *gin.Context) {
	var payload models.District
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	district := models.District{
		Number:       commons.UUIDGenerator(),
		Name:         payload.Name,
		RegionNumber: payload.RegionNumber,
		Size:         payload.Size,
		TownStatus:   payload.TownStatus,
	}

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", payload.RegionNumber).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid region number",
		})
		return
	}

	if err := h.db.DB.Create(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "District created successfully",
	})
}

func (h *DistrictHandler) handleAllDistricts(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var districts []models.District
	if err := h.db.DB.Preload("Region").
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := make([]models.DistrictResponse, len(districts))
	for i, district := range districts {
		response[i] = models.DistrictResponse{
			Number:       district.Number,
			Name:         district.Name,
			Size:         district.Size,
			TownStatus:   district.TownStatus,
			RegionNumber: district.RegionNumber,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *DistrictHandler) handleDistrictByNumber(c *gin.Context) {
	districtNumber := c.Param("number")

	var district models.District
	if err := h.db.DB.Preload("Region").
		Where("number = ?", districtNumber).
		First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "District not found"})
		return
	}

	response := models.DistrictResponse{
		Number:       district.Number,
		Name:         district.Name,
		Size:         district.Size,
		TownStatus:   district.TownStatus,
		RegionNumber: district.RegionNumber,
		RegionName:   district.Region.Name,
	}

	c.JSON(http.StatusOK, response)
}

func (h *DistrictHandler) handleDistrictByName(c *gin.Context) {
	districtName := c.Param("name")

	var district models.District
	if err := h.db.DB.Preload("Region").
		Where("name = ?", districtName).
		First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "District not found"})
		return
	}

	response := models.DistrictResponse{
		Number:       district.Number,
		Name:         district.Name,
		Size:         district.Size,
		TownStatus:   district.TownStatus,
		RegionNumber: district.RegionNumber,
		RegionName:   district.Region.Name,
	}

	c.JSON(http.StatusOK, response)
}

func (h *DistrictHandler) deleteDistrict(c *gin.Context) {
	districtNumber := c.Param("number")
	if districtNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "District number is required"})
		return
	}

	var district models.District
	if err := h.db.DB.Where("number = ?", districtNumber).First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "District not found"})
		return
	}

	if err := h.db.DB.Delete(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "District deleted successfully"})
}
