package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
)

type ParishHandler struct {
	db *database.Database
}

func NewParishHandler(db *database.Database) *ParishHandler {
	return &ParishHandler{
		db: db,
	}
}

func (h *ParishHandler) RegisterRoutes(r *gin.RouterGroup) {
	parishes := r.Group("/parishes")
	{
		parishes.GET("", h.handleAllParishes)
		parishes.POST("", h.createParish)
		parishes.GET("/:number", h.handleParish)
		parishes.PUT("/:number", h.handleUpdateParish)
		parishes.DELETE("/:number", h.handleDeleteParish)
	}
}

func (h *ParishHandler) createParish(c *gin.Context) {
	var payload models.Parish
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parish := models.Parish{
		Number:          commons.UUIDGenerator(),
		SubCountyNumber: payload.SubCountyNumber,
		Name:            payload.Name,
	}

	if err := h.db.DB.Create(&parish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Parish created successfully",
	})
}

func (h *ParishHandler) handleAllParishes(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var parishes []models.Parish
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&parishes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parishes)
}

func (h *ParishHandler) handleParish(c *gin.Context) {
	parishNumber := c.Param("number")

	var parish models.Parish
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parish not found"})
		return
	}

	c.JSON(http.StatusOK, parish)
}

func (h *ParishHandler) handleUpdateParish(c *gin.Context) {
	parishNumber := c.Param("number")

	var payload models.Parish
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var parish models.Parish
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parish not found"})
		return
	}

	parish.Name = payload.Name

	if err := h.db.DB.Save(&parish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Parish updated successfully",
	})
}

func (h *ParishHandler) handleDeleteParish(c *gin.Context) {
	parishNumber := c.Param("number")

	var parish models.District
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parish not found"})
		return
	}

	if err := h.db.DB.Delete(&parish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Parish deleted successfully",
	})
}
