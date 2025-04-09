package routes

import (
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"opendataug.org/commons"
	"opendataug.org/commons/constants"
	"opendataug.org/database"
	"opendataug.org/middleware"
	v1 "opendataug.org/routes/v1"
)

func SetupRouter(db *database.Database) *gin.Engine {
	router := gin.Default()

	if os.Getenv("ENVIRONMENT") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router.Use(static.Serve("./templates/*", static.LocalFile("./templates/*", false)))
	router.Use(middleware.CorsMiddleware())
	router.NoRoute(commons.RouteNotFound)

	v1Group := router.Group("v1")
	if os.Getenv("ENVIRONMENT") == constants.ENVIRONMENT_PROD {
		v1Group.Use(middleware.RateLimit(60, time.Minute, 1))
	}

	v1Group.Use(middleware.TimeoutMiddleware(30 * time.Second))

	{
		// Public routes
		authHandler := v1.NewAuthHandler(db)
		authHandler.RegisterRoutes(v1Group)

		// Protected routes
		protected := v1Group.Group("")
		{
			regionHandler := v1.NewRegionHandler(db)
			regionHandler.RegisterRoutes(protected, authHandler)

			districtHandler := v1.NewDistrictHandler(db)
			districtHandler.RegisterRoutes(protected, authHandler)

			countyHandler := v1.NewCountyHandler(db)
			countyHandler.RegisterRoutes(protected, authHandler)

			subCountyHandler := v1.NewSubcountyHandler(db)
			subCountyHandler.RegisterRoutes(protected, authHandler)

			parishHandler := v1.NewParishHandler(db)
			parishHandler.RegisterRoutes(protected, authHandler)

			villageHandler := v1.NewVillageHandler(db)
			villageHandler.RegisterRoutes(protected, authHandler)

			apiKeyHandler := v1.NewAPIKeyHandler(db)
			apiKeyHandler.RegisterRoutes(protected, authHandler)
		}
	}

	return router
}
