package routes

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"opendataug.org/database"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "https://opendataug.org"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key", "User-Number"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
		// OptionsResponseStatusCode: 200,
		MaxAge: 12 * time.Hour,
	}))

	v1Group := router.Group("/v1")
	{
		// Public routes
		authHandler := v1.NewAuthHandler(db)
		authHandler.RegisterRoutes(v1Group)

		// Protected routes
		protected := v1Group.Group("")
		protected.Use(v1.NewAuthHandler(db).TokenAuthMiddleware())
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
