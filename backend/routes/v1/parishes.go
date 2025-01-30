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

// @Summary Create parish
// @Description Create a new parish
// @Tags Parishes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param parish body models.Parish true "Parish details"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/parishes [post]
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

// @Summary List parishes
// @Description Get a paginated list of all parishes
// @Tags Parishes
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.Parish
// @Failure 500 {object} gin.H
// @Router /v1/parishes [get]
func (h *ParishHandler) handleAllParishes(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var parishes []models.Parish
	if err := h.db.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&parishes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parishes)
}

// @Summary Get parish
// @Description Get a parish by its number
// @Tags Parishes
// @Accept json
// @Produce json
// @Param number path string true "Parish Number"
// @Success 200 {object} models.Parish
// @Failure 404 {object} gin.H
// @Router /v1/parishes/{number} [get]
func (h *ParishHandler) handleParish(c *gin.Context) {
	parishNumber := c.Param("number")

	var parish models.Parish
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parish not found"})
		return
	}

	c.JSON(http.StatusOK, parish)
}

// @Summary Update parish
// @Description Update an existing parish
// @Tags Parishes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param number path string true "Parish Number"
// @Param parish body models.Parish true "Updated parish details"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/parishes/{number} [put]
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

// @Summary Delete parish
// @Description Delete a parish
// @Tags Parishes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param number path string true "Parish Number"
// @Success 200 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /v1/parishes/{number} [delete]
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
