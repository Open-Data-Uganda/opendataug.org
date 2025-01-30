package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uganda-data/commons"
	"github.com/uganda-data/database"
	"github.com/uganda-data/models"
)

type VillageHandler struct {
	db *database.Database
}

func NewVillageHandler(db *database.Database) *VillageHandler {
	return &VillageHandler{
		db: db,
	}
}

func (h *VillageHandler) RegisterRoutes(r *gin.RouterGroup) {
	villages := r.Group("/villages")
	{
		villages.GET("", h.handleAllVillages)
		villages.GET("/:number", h.handleGetVillage)
		villages.POST("", h.createVillage)
		villages.PUT("/:number", h.updateVillage)
		villages.DELETE("/:number", h.deleteVillage)
	}
}

// @Summary Create village
// @Description Create a new village
// @Tags Villages
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param village body models.Village true "Village details"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/villages [post]
func (h *VillageHandler) createVillage(c *gin.Context) {
	var payload models.Village
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	village := models.Village{
		Number:       commons.UUIDGenerator(),
		Name:         payload.Name,
		ParishNumber: payload.ParishNumber,
	}

	if err := h.db.DB.Create(&village).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Village created successfully",
	})
}

// @Summary Update village
// @Description Update an existing village
// @Tags Villages
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param number path string true "Village Number"
// @Param village body models.Village true "Updated village details"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/villages/{number} [put]
func (h *VillageHandler) updateVillage(c *gin.Context) {
	id := c.Param("number")

	var village models.Village
	if err := h.db.DB.First(&village, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Village not found"})
		return
	}

	var payload models.Village
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	village.Name = payload.Name
	village.ParishNumber = payload.ParishNumber

	if err := h.db.DB.Save(&village).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Village updated successfully",
	})
}

// @Summary Delete village
// @Description Delete a village
// @Tags Villages
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param number path string true "Village Number"
// @Success 200 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/villages/{number} [delete]
func (h *VillageHandler) deleteVillage(c *gin.Context) {
	id := c.Param("number")

	var village models.Village
	if err := h.db.DB.First(&village, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Village not found"})
		return
	}

	if err := h.db.DB.Delete(&village).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Village deleted successfully",
	})
}

// @Summary List villages
// @Description Get a paginated list of all villages
// @Tags Villages
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.Village
// @Failure 500 {object} gin.H
// @Router /v1/villages [get]
func (h *VillageHandler) handleAllVillages(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var villages []models.Village
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&villages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, villages)
}

// @Summary Get village
// @Description Get a village by its number
// @Tags Villages
// @Accept json
// @Produce json
// @Param number path string true "Village Number"
// @Success 200 {object} models.Village
// @Failure 404 {object} gin.H
// @Router /v1/villages/{number} [get]
func (h *VillageHandler) handleGetVillage(c *gin.Context) {
	id := c.Param("number")

	var village models.Village
	if err := h.db.DB.First(&village, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Village not found"})
		return
	}

	c.JSON(http.StatusOK, village)
}
