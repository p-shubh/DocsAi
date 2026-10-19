package document

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
) {
	documents := router.Group("/documents")
	{
		documents.POST("", handler.Upload)
		documents.GET("/:id", handler.GetDocument)
		documents.GET("/:id/file", handler.GetDocumentFile)
		documents.GET("/:id/chunks", handler.GetDocumentChunks)
		documents.GET("", handler.ListDocuments)
		documents.PATCH("/:id", handler.RenameDocument)
		documents.DELETE("/:id", handler.DeleteDocument)
	}
}
