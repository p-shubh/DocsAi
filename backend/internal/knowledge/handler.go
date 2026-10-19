package knowledge

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	knowledgedto "github.com/hexanecreations/docsai/backend/internal/knowledge/dto"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AskQuestion(c *gin.Context) {

	var req knowledgedto.AskQuestionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var documentID *uuid.UUID

	if req.DocumentID != nil {
		parsed, err := uuid.Parse(*req.DocumentID)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "invalid document id",
				},
			)
			return
		}
		documentID = &parsed
	}

	response, err := h.service.AskQuestion(
		c.Request.Context(),
		req.Question,
		documentID,
	)
	if err != nil {

		if errors.Is(err, ErrNoRelevantContext) {
			c.JSON(
				http.StatusOK,
				knowledgedto.AskQuestionResponse{
					Answer:  "I don't have any relevant information to answer that question.",
					Sources: []knowledgedto.SourceCitation{},
				},
			)
			return
		}

		if st, ok := status.FromError(err); ok && st.Code() == codes.Unavailable {
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"error": "AI service unavailable",
				},
			)
			return
		}

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

type sseEvent struct {
	Type    string                        `json:"type"`
	Sources []knowledgedto.SourceCitation `json:"sources,omitempty"`
	Text    string                        `json:"text,omitempty"`
	Model   string                        `json:"model,omitempty"`
	Message string                        `json:"message,omitempty"`
}

func (h *Handler) AskQuestionStream(c *gin.Context) {

	var req knowledgedto.AskQuestionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	var documentID *uuid.UUID

	if req.DocumentID != nil {
		parsed, err := uuid.Parse(*req.DocumentID)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "invalid document id",
				},
			)
			return
		}
		documentID = &parsed
	}

	headersSent := false

	writeEvent := func(payload sseEvent) {
		if !headersSent {
			c.Writer.Header().Set("Content-Type", "text/event-stream")
			c.Writer.Header().Set("Cache-Control", "no-cache")
			c.Writer.Header().Set("Connection", "keep-alive")
			c.Writer.WriteHeader(http.StatusOK)
			headersSent = true
		}

		data, _ := json.Marshal(payload)
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		c.Writer.Flush()
	}

	model, err := h.service.AskQuestionStream(
		c.Request.Context(),
		req.Question,
		documentID,
		func(sources []knowledgedto.SourceCitation) {
			writeEvent(sseEvent{Type: "sources", Sources: sources})
		},
		func(delta string) {
			writeEvent(sseEvent{Type: "delta", Text: delta})
		},
	)
	if err != nil {

		if !headersSent {
			if st, ok := status.FromError(err); ok && st.Code() == codes.Unavailable {
				c.JSON(
					http.StatusBadGateway,
					gin.H{
						"error": "AI service unavailable",
					},
				)
				return
			}

			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		writeEvent(sseEvent{Type: "error", Message: err.Error()})
		return
	}

	writeEvent(sseEvent{Type: "done", Model: model})
}
