package document

import (
	"context"

	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
)

type Repository interface {
	Create(ctx context.Context, document *schema.Document) error

	GetByID(ctx context.Context, id uuid.UUID) (*schema.Document, error)

	Update(ctx context.Context, document *schema.Document) error

	Delete(ctx context.Context, id uuid.UUID) error

	List(ctx context.Context, filter ListDocumentsFilter) ([]schema.Document, int64, error)

	SaveChunks(ctx context.Context, chunks []schema.DocumentChunk) error

	ListChunksByDocumentID(ctx context.Context, documentID uuid.UUID) ([]schema.DocumentChunk, error)

	SaveEmbeddings(ctx context.Context, embeddings []schema.DocumentEmbedding) error

	SearchRelevantChunks(ctx context.Context, embedding []float32, limit int, documentID *uuid.UUID) ([]RetrievedChunk, error)

	UpdateStatus(ctx context.Context, documentID uuid.UUID, status enums.DocumentStatus) error
}
