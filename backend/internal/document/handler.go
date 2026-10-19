package document

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	documentdto "github.com/hexanecreations/docsai/backend/internal/document/dto"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Upload(c *gin.Context) {

	var req documentdto.UploadDocumentRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	resp, err := h.service.UploadDocument(
		c.Request.Context(),
		req,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		resp,
	)
}

func (h *Handler) GetDocument(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid document id",
			},
		)
		return
	}

	response, err := h.service.GetDocument(
		c.Request.Context(),
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *Handler) GetDocumentFile(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid document id",
			},
		)
		return
	}

	response, err := h.service.GetDocumentFile(
		c.Request.Context(),
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *Handler) GetDocumentChunks(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid document id",
			},
		)
		return
	}

	response, err := h.service.GetDocumentChunks(
		c.Request.Context(),
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *Handler) ListDocuments(c *gin.Context) {

	filter := ListDocumentsFilter{
		Limit:  10,
		Offset: 0,
	}

	response, err := h.service.ListDocuments(
		c.Request.Context(),
		filter,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *Handler) RenameDocument(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid document id",
			},
		)
		return
	}

	var request documentdto.RenameDocumentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	response, err := h.service.RenameDocument(
		c.Request.Context(),
		id,
		request,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *Handler) DeleteDocument(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid document id",
			},
		)
		return
	}

	response, err := h.service.DeleteDocument(
		c.Request.Context(),
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
