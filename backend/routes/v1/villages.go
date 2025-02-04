package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
)

type VillageHandler struct {
	db *database.Database
}

func NewVillageHandler(db *database.Database) *VillageHandler {
	return &VillageHandler{
		db: db,
	}
}

func (h *VillageHandler) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	villages := r.Group("/villages")
	{

		apiProtected := villages.Group("")
		apiProtected.Use(authHandler.APIAuthMiddleware())
		{
			villages.GET("", h.handleAllVillages)
			villages.GET("/:number", h.handleGetVillage)
		}

		villages.POST("", h.createVillage)
		villages.PUT("/:number", h.updateVillage)
		villages.DELETE("/:number", h.deleteVillage)
	}
}

func (h *VillageHandler) createVillage(c *gin.Context) {
	var payload models.Village
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var parish models.Parish
	if err := h.db.DB.First(&parish, "number = ?", payload.ParishNumber).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parish number - parish does not exist"})
		return
	}

	// Check if village with same name exists in the same parish
	var existingVillage models.Village
	if err := h.db.DB.Where("name = ? AND parish_number = ?", payload.Name, payload.ParishNumber).
		First(&existingVillage).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Village with this name already exists in this parish"})
		return
	}

	village := models.Village{
		Number:       commons.UUIDGenerator(),
		Name:         payload.Name,
		ParishNumber: payload.ParishNumber,
	}

	if err := h.db.DB.Create(&village).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Village created successfully",
	})
}

func (h *VillageHandler) updateVillage(c *gin.Context) {
	id := c.Param("number")

	var village models.Village
	if err := h.db.DB.First(&village, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Village not found"})
		return
	}

	var payload models.Village
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	village.Name = payload.Name
	village.ParishNumber = payload.ParishNumber

	if err := h.db.DB.Save(&village).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Village updated successfully",
	})
}

func (h *VillageHandler) deleteVillage(c *gin.Context) {
	id := c.Param("number")

	var village models.Village
	if err := h.db.DB.First(&village, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Village not found"})
		return
	}

	if err := h.db.DB.Delete(&village).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Village deleted successfully",
	})
}

func (h *VillageHandler) handleAllVillages(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var villages []models.Village
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&villages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, villages)
}

func (h *VillageHandler) handleGetVillage(c *gin.Context) {
	id := c.Param("number")

	var village models.Village
	if err := h.db.DB.First(&village, "number = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Village not found"})
		return
	}

	c.JSON(http.StatusOK, village)
}
