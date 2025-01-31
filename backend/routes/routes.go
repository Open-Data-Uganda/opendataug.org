package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"opendataug.org/database"
	"opendataug.org/middleware"
	v1 "opendataug.org/routes/v1"
)

func SetupRouter(db *database.Database) *gin.Engine {
	router := gin.Default()

	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	v1Group := router.Group("/v1")
	{
		// Public routes
		authHandler := v1.NewAuthHandler(db)
		authHandler.RegisterRoutes(v1Group)

		// Protected routes
		protected := v1Group.Group("")
		protected.Use(middleware.APIKeyAuth(db))
		{
			regionHandler := v1.NewRegionHandler(db)
			regionHandler.RegisterRoutes(protected)

			districtHandler := v1.NewDistrictHandler(db)
			districtHandler.RegisterRoutes(protected)

			countyHandler := v1.NewCountyHandler(db)
			countyHandler.RegisterRoutes(protected)

			subCountyHandler := v1.NewSubcountyHandler(db)
			subCountyHandler.RegisterRoutes(protected)

			parishHandler := v1.NewParishHandler(db)
			parishHandler.RegisterRoutes(protected)

			villageHandler := v1.NewVillageHandler(db)
			villageHandler.RegisterRoutes(protected)

			apiKeyHandler := v1.NewAPIKeyHandler(db)
			apiKeyHandler.RegisterRoutes(protected)

		}
	}

	return router
}
