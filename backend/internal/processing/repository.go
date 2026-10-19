package processing

import (
	"context"

	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
)

type Repository interface {
	SaveProcessingResult(
		ctx context.Context,
		documentID uuid.UUID,
		status enums.DocumentStatus,
		chunks []schema.DocumentChunk,
		embeddings []schema.DocumentEmbedding,
	) error
}
