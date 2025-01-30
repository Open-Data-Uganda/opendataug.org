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

// @Summary Create district
// @Description Create a new district
// @Tags Districts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param district body models.District true "District details"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/districts [post]
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

// @Summary List districts
// @Description Get a paginated list of all districts
// @Tags Districts
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.District
// @Failure 500 {object} gin.H
// @Router /v1/districts [get]
func (h *DistrictHandler) handleAllDistricts(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var districts []models.District
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, districts)
}

// @Summary Get district by number
// @Description Get a district by its unique number
// @Tags Districts
// @Accept json
// @Produce json
// @Param number path string true "District Number"
// @Success 200 {object} models.District
// @Failure 404 {object} gin.H
// @Router /v1/districts/{number} [get]
func (h *DistrictHandler) handleDistrictByNumber(c *gin.Context) {
	districtNumber := c.Param("number")

	var district models.District
	if err := h.db.DB.Where("number = ?", districtNumber).First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	c.JSON(http.StatusOK, district)
}

// @Summary Get district by name
// @Description Get a district by its name
// @Tags Districts
// @Accept json
// @Produce json
// @Param name path string true "District Name"
// @Success 200 {object} models.District
// @Failure 404 {object} gin.H
// @Router /v1/districts/name/{name} [get]
func (h *DistrictHandler) handleDistrictByName(c *gin.Context) {
	districtName := c.Param("name")

	var district models.District
	if err := h.db.DB.Where("name = ?", districtName).First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	c.JSON(http.StatusOK, district)
}
