package sort

import "github.com/gin-gonic/gin"

func Parse(c *gin.Context) Sort {

	by := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	return Sort{
		By:    by,
		Order: order,
	}
}