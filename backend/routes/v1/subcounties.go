package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uganda-data/commons"
	"github.com/uganda-data/database"
	"github.com/uganda-data/models"
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

// @Summary Create subcounty
// @Description Create a new subcounty
// @Tags Subcounties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param subcounty body models.SubCounty true "Subcounty details"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/subcounties [post]
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

// @Summary List subcounties
// @Description Get a paginated list of all subcounties
// @Tags Subcounties
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.SubCounty
// @Failure 500 {object} gin.H
// @Router /v1/subcounties [get]
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

// @Summary Get subcounty
// @Description Get a subcounty by its number
// @Tags Subcounties
// @Accept json
// @Produce json
// @Param number path string true "Subcounty Number"
// @Success 200 {object} models.SubCounty
// @Failure 404 {object} gin.H
// @Router /v1/subcounties/{number} [get]
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

// @Summary Update subcounty
// @Description Update an existing subcounty
// @Tags Subcounties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param number path string true "Subcounty Number"
// @Param subcounty body models.SubCounty true "Updated subcounty details"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/subcounties/{number} [put]
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

// @Summary Delete subcounty
// @Description Delete a subcounty
// @Tags Subcounties
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param number path string true "Subcounty Number"
// @Success 200 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/subcounties/{number} [delete]
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

// @Summary List parishes in subcounty
// @Description Get a paginated list of all parishes in a subcounty
// @Tags Subcounties
// @Accept json
// @Produce json
// @Param number path string true "Subcounty Number"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.Parish
// @Failure 500 {object} gin.H
// @Router /v1/subcounties/{number}/parishes [get]
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
