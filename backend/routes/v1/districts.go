package v1

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/commons/constants"
	"opendataug.org/database"
	customerrors "opendataug.org/errors"
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

func (h *DistrictHandler) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	districts := r.Group("/districts")
	{
		apiProtected := districts.Group("")
		apiProtected.Use(authHandler.APIAuthMiddleware())
		{
			districts.GET("", h.handleAllDistricts)
			districts.GET("/:id", h.handleDistrictByNumber)
			districts.GET("/name/:name", h.handleDistrictByName)
		}

		private := districts.Group("")
		private.Use(authHandler.TokenAuthMiddleware())
		{
			districts.POST("", h.createDistrict)
			districts.DELETE("/:id", h.deleteDistrict)
		}
	}
}

func (h *DistrictHandler) createDistrict(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	var payload models.District
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError("Invalid request payload"))
		return
	}

	district := models.District{
		Number:       commons.UUIDGenerator(),
		Name:         payload.Name,
		RegionNumber: payload.RegionNumber,
		TownStatus:   payload.TownStatus,
	}

	var region models.Region
	if err := h.db.DB.First(&region, "number = ?", payload.RegionNumber).Error; err != nil {
		c.JSON(http.StatusBadRequest, customerrors.NewValidationError("Invalid region number"))
		return
	}

	if err := h.db.DB.Create(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Failed to create district"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "District created successfully",
	})
}

func (h *DistrictHandler) handleAllDistricts(c *gin.Context) {
	pagination := commons.GetPaginationParams(c)

	var districts []models.District
	if err := h.db.DB.Preload("Region").
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := make([]models.DistrictResponse, len(districts))
	for i, district := range districts {
		response[i] = models.DistrictResponse{
			ID:         district.Number,
			Name:       district.Name,
			TownStatus: district.TownStatus,
			RegionID:   district.RegionNumber,
			RegionName: district.Region.Name,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *DistrictHandler) handleDistrictByNumber(c *gin.Context) {
	districtNumber := c.Param("id")
	if districtNumber == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("District id is required"))
		return
	}
	districtNumber = commons.Sanitize(districtNumber)

	var district models.District
	if err := h.db.DB.Preload("Region").
		Where("number = ?", districtNumber).
		First(&district).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("District not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Failed to fetch district"))
		return
	}

	response := models.DistrictResponse{
		ID:         district.Number,
		Name:       district.Name,
		TownStatus: district.TownStatus,
		RegionID:   district.RegionNumber,
		RegionName: district.Region.Name,
	}

	c.JSON(http.StatusOK, response)
}

func (h *DistrictHandler) handleDistrictByName(c *gin.Context) {
	districtName := c.Param("name")
	if districtName == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("District name is required"))
		return
	}
	districtName = commons.Sanitize(districtName)

	var district models.District
	if err := h.db.DB.Preload("Region").
		Where("name = ?", districtName).
		First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("District not found"))
		return
	}

	response := models.DistrictResponse{
		ID:         district.Number,
		Name:       district.Name,
		TownStatus: district.TownStatus,
		RegionID:   district.RegionNumber,
		RegionName: district.Region.Name,
	}

	c.JSON(http.StatusOK, response)
}

func (h *DistrictHandler) deleteDistrict(c *gin.Context) {
	user, _ := commons.GetUserFromHeader(c, h.db.DB)
	if user.Role != constants.RoleAdmin {
		c.JSON(http.StatusUnauthorized, customerrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	districtNumber := c.Param("id")
	if districtNumber == "" {
		c.JSON(http.StatusBadRequest, customerrors.NewBadRequestError("District number is required"))
		return
	}
	districtNumber = commons.Sanitize(districtNumber)

	var district models.District
	if err := h.db.DB.Where("number = ?", districtNumber).First(&district).Error; err != nil {
		c.JSON(http.StatusNotFound, customerrors.NewNotFoundError("District not found"))
		return
	}

	if err := h.db.DB.Delete(&district).Error; err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.NewDatabaseError("Failed to delete district"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "District deleted successfully"})
}
