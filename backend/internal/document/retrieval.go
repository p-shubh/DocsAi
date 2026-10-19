package document

import "github.com/google/uuid"

type RetrievedChunk struct {
	ChunkID uuid.UUID

	DocumentID uuid.UUID

	DocumentName string

	Page int32

	Section string

	Text string

	Distance float64
}
