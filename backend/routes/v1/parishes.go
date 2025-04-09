package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/commons/constants"
	"opendataug.org/database"
	customerrors "opendataug.org/errors"
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

func (h *ParishHandler) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	parishes := r.Group("/parishes")
	{
		apiProtected := parishes.Group("")
		apiProtected.Use(authHandler.APIAuthMiddleware())
		{
			parishes.GET("", h.handleAllParishes)
			parishes.GET("/:id", h.handleParish)
			parishes.GET("/:id/villages", h.handleParishVillages)
		}

		private := parishes.Group("")
		private.Use(authHandler.TokenAuthMiddleware())
		{
			parishes.POST("", h.createParish)
			parishes.PUT("/:id", h.handleUpdateParish)
			parishes.DELETE("/:id", h.handleDeleteParish)
		}
	}
}

func (h *ParishHandler) createParish(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin && !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	var payload models.Parish
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError("Failed to parse requestbody"))
		return
	}

	var existingParish models.Parish
	if err := h.db.DB.Where("name = ? AND sub_county_number = ?", payload.Name, payload.SubCountyNumber).
		First(&existingParish).Error; err == nil {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("Parish with this name already exists in this subcounty"))
		return
	}

	parish := models.Parish{
		Number:          commons.UUIDGenerator(),
		SubCountyNumber: payload.SubCountyNumber,
		Name:            payload.Name,
	}

	if err := h.db.DB.Create(&parish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
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
	parishNumber := c.Param("id")

	if parishNumber == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("Invalid parish number"))
		return
	}
	parishNumber = commons.Sanitize(parishNumber)

	var parish models.Parish
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Parish not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	response := models.ParishResponse{
		Name: parish.Name,
		ID:   parish.Number,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ParishHandler) handleUpdateParish(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin && !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	parishNumber := c.Param("id")

	var payload models.Parish
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError("Failed to parse request body"))
		return
	}

	var parish models.Parish
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Parish not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	parish.Name = payload.Name

	if err := h.db.DB.Save(&parish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Parish updated successfully",
	})
}

func (h *ParishHandler) handleDeleteParish(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin && !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	parishNumber := c.Param("id")

	var parish models.District
	if err := h.db.DB.Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Parish not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	if err := h.db.DB.Delete(&parish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Parish deleted successfully",
	})
}

func (h *ParishHandler) handleParishVillages(c *gin.Context) {
	parishNumber := c.Param("id")
	parishNumber = commons.Sanitize(parishNumber)

	if parishNumber == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("Invalid parish number"))
		return
	}

	var parish models.Parish
	if err := h.db.DB.Preload("Villages").Where("number = ?", parishNumber).First(&parish).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Parish not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	var villages []models.VillageResponse
	for _, village := range parish.Villages {
		villages = append(villages, models.VillageResponse{
			Name: village.Name,
			ID:   village.Number,
		})
	}

	c.JSON(http.StatusOK, villages)
}
