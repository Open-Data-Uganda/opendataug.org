package handlers

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
		counties.GET("", h.handleGetCounties)
		counties.GET("/:number", h.handleGetCounty)
		counties.POST("", h.handleCreateCounty)
		counties.PUT("/:number", h.handleUpdateCounty)
		counties.DELETE("/:number", h.handleDeleteCounty)
		counties.GET("/:number/subcounties", h.handleGetCountySubCounties)
	}
}

func (h *CountyHandler) handleGetCounties(c *gin.Context) {
	var counties []models.County
	if err := h.db.DB.Find(&counties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counties)
}

func (h *CountyHandler) handleGetCounty(c *gin.Context) {
	number := c.Param("number")
	var county models.County
	if err := h.db.DB.Preload("SubCounties").First(&county, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}
	c.JSON(http.StatusOK, county)
}

func (h *CountyHandler) handleCreateCounty(c *gin.Context) {
	var payload models.County
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusCreated, county)
}

func (h *CountyHandler) handleUpdateCounty(c *gin.Context) {
	number := c.Param("number")
	var county models.County
	if err := h.db.DB.First(&county, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	if err := c.ShouldBindJSON(&county); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB.Save(&county).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, county)
}

func (h *CountyHandler) handleDeleteCounty(c *gin.Context) {
	number := c.Param("number")
	if err := h.db.DB.Delete(&models.County{}, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "County deleted successfully"})
}

func (h *CountyHandler) handleGetCountySubCounties(c *gin.Context) {
	number := c.Param("number")
	var subcounties []models.SubCounty
	if err := h.db.DB.Where("county_number = ?", number).Find(&subcounties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcounties)
}
