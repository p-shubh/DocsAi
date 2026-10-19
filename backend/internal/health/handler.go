package health

import "net/http"

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "DocsAI Backend",
		"version": "v1",
	})
}