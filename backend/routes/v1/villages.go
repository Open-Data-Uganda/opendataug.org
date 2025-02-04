package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opendataug.org/controllers"
	"opendataug.org/database"
)

type VillageHandler struct {
	controller *controllers.VillageController
}

func NewVillageHandler(db *database.Database) *VillageHandler {
	return &VillageHandler{
		controller: controllers.NewVillageController(db),
	}
}

func (h *VillageHandler) RegisterRoutes(r *gin.RouterGroup, authHandler *AuthHandler) {
	villages := r.Group("/villages")
	{
		apiProtected := villages.Group("")
		apiProtected.Use(authHandler.APIAuthMiddleware())
		{
			villages.GET("", h.handleAllVillages)
			villages.GET("/:id", h.handleGetVillage)
		}

		private := villages.Group("")
		private.Use(authHandler.TokenAuthMiddleware())
		{
			villages.POST("", h.createVillage)
			villages.PUT("/:id", h.updateVillage)
			villages.DELETE("/:id", h.deleteVillage)
		}
	}
}

func (h *VillageHandler) createVillage(c *gin.Context) {
	if err := h.controller.CreateVillage(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Village created successfully",
	})
}

func (h *VillageHandler) updateVillage(c *gin.Context) {
	if err := h.controller.UpdateVillage(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Village updated successfully",
	})
}

func (h *VillageHandler) deleteVillage(c *gin.Context) {
	if err := h.controller.DeleteVillage(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Village deleted successfully",
	})
}

func (h *VillageHandler) handleAllVillages(c *gin.Context) {
	villages, err := h.controller.GetAllVillages(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, villages)
}

func (h *VillageHandler) handleGetVillage(c *gin.Context) {
	village, err := h.controller.GetVillage(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, village)
}
