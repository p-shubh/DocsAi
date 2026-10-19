package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Page   int
	Limit  int
	Offset int

	Search string
}

func ParseQuery(c *gin.Context) Query {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	return Query{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
		Search: c.Query("search"),
	}
}