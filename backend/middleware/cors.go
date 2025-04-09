package middleware

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("FRONTEND_URL"),
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"x-api-key",
			"User-Number",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Set-Cookie",
		},
		AllowCredentials:          true,
		OptionsResponseStatusCode: 204,
		MaxAge:                    12 * time.Hour,
	})
}
