package documentdto

import (
	"time"

	"github.com/google/uuid"
)

type UploadDocumentResponse struct {
	ID uuid.UUID `json:"id"`

	OriginalName string `json:"original_name"`

	Size int64 `json:"size"`

	MimeType string `json:"mime_type"`

	Status string `json:"status"`

	UploadedAt time.Time `json:"uploaded_at"`
}

type DocumentResponse struct {
	ID uuid.UUID `json:"id"`

	OriginalName string `json:"original_name"`

	Size int64 `json:"size"`

	MimeType string `json:"mime_type"`

	Status string `json:"status"`

	PageCount int `json:"page_count"`

	SummaryAvailable bool `json:"summary_available"`

	OCRCompleted bool `json:"ocr_completed"`

	EmbeddingGenerated bool `json:"embedding_generated"`

	UploadedAt time.Time `json:"uploaded_at"`
}

type GetDocumentResponse = DocumentResponse

type ListDocumentsResponse struct {
	Documents []DocumentResponse `json:"documents"`

	Total int64 `json:"total"`

	Page int `json:"page"`

	Limit int `json:"limit"`
}

type DeleteDocumentResponse struct {
	Message string `json:"message"`
}

type DocumentFileResponse struct {
	FileURL string `json:"file_url"`

	MimeType string `json:"mime_type"`

	OriginalName string `json:"original_name"`
}

type ChunkResponse struct {
	ChunkIndex int32 `json:"chunk_index"`

	Page int32 `json:"page"`

	Section string `json:"section"`

	Text string `json:"text"`
}

type ListChunksResponse struct {
	Chunks []ChunkResponse `json:"chunks"`
}

type RenameDocumentRequest struct {
	Name string `json:"name" binding:"required"`
}
