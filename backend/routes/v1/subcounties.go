package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/commons/constants"
	"opendataug.org/database"
	customerrors "opendataug.org/errors"
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

func (h *SubcountyHandle) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	subcounties := r.Group("/subcounties")
	{
		apiProtected := subcounties.Group("")
		apiProtected.Use(authHandler.APIAuthMiddleware())
		{
			subcounties.GET("", h.handleAllSubCounties)
			subcounties.GET("/:id", h.handleGetSubCounty)
			subcounties.GET("/:id/parishes", h.handleParishes)
		}

		private := subcounties.Group("")
		private.Use(authHandler.TokenAuthMiddleware())
		{
			subcounties.POST("", h.createSubcounty)
			subcounties.PUT("/:id", h.updateSubCounty)
			subcounties.DELETE("/:id", h.deleteSubCounty)
		}
	}
}

func (h *SubcountyHandle) createSubcounty(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin && !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	var payload models.SubCounty
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError("Failed to parse request body"))
		return
	}

	var existingSubcounty models.SubCounty
	if err := h.db.DB.Where("name = ? AND county_number = ?", payload.Name, payload.CountyNumber).
		First(&existingSubcounty).Error; err == nil {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("Subcounty with this name already exists in this county"))
		return
	}

	subcounty := models.SubCounty{
		Number:       commons.UUIDGenerator(),
		CountyNumber: payload.CountyNumber,
		Name:         payload.Name,
	}

	if err := h.db.DB.Create(&subcounty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subcounty created successfully"})
}

func (h *SubcountyHandle) handleAllSubCounties(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var subcounties []models.SubCounty
	if err := h.db.DB.Preload("Parishes").
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&subcounties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	c.JSON(http.StatusOK, subcounties)
}

func (h *SubcountyHandle) handleGetSubCounty(c *gin.Context) {
	number := c.Param("id")

	var subcounty models.SubCounty
	if err := h.db.DB.Preload("Parishes").
		First(&subcounty, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Sub county not found"))
		return
	}

	c.JSON(http.StatusOK, subcounty)
}

func (h *SubcountyHandle) updateSubCounty(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin && !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	number := c.Param("id")

	var subcounty models.SubCounty
	if err := h.db.DB.First(&subcounty, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Sub county not found"))
		return
	}

	var payload models.SubCounty
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError("Failed to parse request body"))
		return
	}

	subcounty.Name = payload.Name
	subcounty.CountyNumber = payload.CountyNumber

	if err := h.db.DB.Save(&subcounty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcounty updated successfully"})
}

func (h *SubcountyHandle) deleteSubCounty(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin && !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	number := c.Param("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("ID is required"))
		return
	}
	number = commons.Sanitize(number)

	var subcounty models.SubCounty
	if err := h.db.DB.First(&subcounty, "number = ?", number).Error; err != nil {
		c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("Subcounty not found"))
		return
	}

	if err := h.db.DB.Delete(&subcounty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcounty deleted successfully"})
}

func (h *SubcountyHandle) handleParishes(c *gin.Context) {
	number := c.Param("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("Number is required"))
		return
	}

	number = commons.Sanitize(number)

	var parishes []models.Parish
	if err := h.db.DB.Where("sub_county_number = ?", number).
		Find(&parishes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Database level error occurred"))
		return
	}

	var response []models.ParishResponse
	for _, parish := range parishes {
		response = append(response, models.ParishResponse{
			Name: parish.Name,
			ID:   parish.Number,
		})
	}

	c.JSON(http.StatusOK, response)
}
