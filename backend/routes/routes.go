package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"opendataug.org/database"
	_ "opendataug.org/docs"
	"opendataug.org/middleware"
	v1 "opendataug.org/routes/v1"
)

// @title Uganda Data API
// @version 1.0
// @description This is the API server for Uganda Data.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @description API Key authentication

func SetupRouter(db *database.Database) *gin.Engine {
	router := gin.Default()

	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Group := router.Group("/api/v1")
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
