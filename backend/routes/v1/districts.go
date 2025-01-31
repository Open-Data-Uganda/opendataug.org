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
	}
}

func (h *DistrictHandler) createDistrict(c *gin.Context) {
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
		TownStatus: payload.TownStatus,
	}

	if err := h.db.DB.Create(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "District created successfully",
	})
}

func (h *DistrictHandler) handleAllDistricts(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var districts []models.District
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, districts)
}

func (h *DistrictHandler) handleDistrictByNumber(c *gin.Context) {
	districtNumber := c.Param("number")

	var district models.District
	if err := h.db.DB.Where("number = ?", districtNumber).First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	c.JSON(http.StatusOK, district)
}

func (h *DistrictHandler) handleDistrictByName(c *gin.Context) {
	districtName := c.Param("name")

	var district models.District
	if err := h.db.DB.Where("name = ?", districtName).First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	c.JSON(http.StatusOK, district)
}
