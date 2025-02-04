package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"opendataug.org/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if apiErr, ok := err.(*errors.APIError); ok {
				log.Printf("API Error: %v", apiErr)

				c.JSON(apiErr.StatusCode, gin.H{
					"error": apiErr,
				})
				return
			}

			c.JSON(500, gin.H{
				"error": errors.NewInternalError("An unexpected error occurred"),
			})
		}
	}
}
