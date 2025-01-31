package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
)

type SubcountyHandle struct {
	db *database.Database
}

func NewSubcountyHandler(db *database.Database) *SubcountyHandle {
	return &SubcountyHandle{
		db: db,
	}
}

func (h *SubcountyHandle) RegisterRoutes(r *gin.RouterGroup) {
	subcounties := r.Group("/subcounties")
	{
		subcounties.GET("", h.handleAllSubCounties)
		subcounties.GET("/:number", h.handleGetSubCounty)
		subcounties.GET("/:number/parishes", h.handleParishes)
		subcounties.POST("", h.createSubcounty)
		subcounties.PUT("/:number", h.updateSubCounty)
		subcounties.DELETE("/:number", h.deleteSubCounty)
	}
}

func (h *SubcountyHandle) createSubcounty(c *gin.Context) {
	var payload models.SubCounty
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subcounty := models.SubCounty{
		Number:       commons.UUIDGenerator(),
		CountyNumber: payload.CountyNumber,
		Name:         payload.Name,
	}

	if err := h.db.DB.Create(&subcounty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Subcounty created successfully",
	})
}

func (h *SubcountyHandle) handleAllSubCounties(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var subcounties []models.SubCounty
	if err := h.db.DB.Preload("Parishes").
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&subcounties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subcounties)
}

func (h *SubcountyHandle) handleGetSubCounty(c *gin.Context) {
	number := c.Param("number")

	var subcounty models.SubCounty
	if err := h.db.DB.Preload("Parishes").
		First(&subcounty, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub county not found"})
		return
	}

	c.JSON(http.StatusOK, subcounty)
}

func (h *SubcountyHandle) updateSubCounty(c *gin.Context) {
	number := c.Param("number")

	var subcounty models.SubCounty
	if err := h.db.DB.First(&subcounty, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub county not found"})
		return
	}

	var payload models.SubCounty
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subcounty.Name = payload.Name
	subcounty.CountyNumber = payload.CountyNumber

	if err := h.db.DB.Save(&subcounty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Subcounty updated successfully",
	})
}

func (h *SubcountyHandle) deleteSubCounty(c *gin.Context) {
	number := c.Param("number")

	var subcounty models.SubCounty
	if err := h.db.DB.First(&subcounty, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcounty not found"})
		return
	}

	if err := h.db.DB.Delete(&subcounty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Subcounty deleted successfully",
	})
}

func (h *SubcountyHandle) handleParishes(c *gin.Context) {
	number := c.Param("number")
	pagination := commons.GetPaginationParams(c)

	var parishes []models.Parish
	if err := h.db.DB.Where("sub_county_number = ?", number).
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&parishes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parishes)
}
