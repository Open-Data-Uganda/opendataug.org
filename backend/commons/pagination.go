package commons

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaginationParams struct {
	Page  int
	Limit int
}

func GetPaginationParams(c *gin.Context) PaginationParams {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}
}
