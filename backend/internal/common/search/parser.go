package search

import "github.com/gin-gonic/gin"

func Parse(c *gin.Context) Search {
	return Search{
		Query: c.Query("search"),
	}
}