package knowledge

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
) {
	chat := router.Group("/chat")
	{
		chat.POST("", handler.AskQuestion)
		chat.POST("/stream", handler.AskQuestionStream)
	}
}
