package document

import (
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
	documentdto "github.com/hexanecreations/docsai/backend/internal/document/dto"
)

func ToDocumentResponse(
	document *schema.Document,
) *documentdto.DocumentResponse {

	return &documentdto.DocumentResponse{
		ID:                 document.ID,
		OriginalName:       document.OriginalName,
		Size:               document.Size,
		MimeType:           document.MimeType,
		Status:             string(document.Status),
		PageCount:          0,
		SummaryAvailable:   false,
		OCRCompleted:       false,
		EmbeddingGenerated: false,
		UploadedAt:         document.CreatedAt,
	}
}
