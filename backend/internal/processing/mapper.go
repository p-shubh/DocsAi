package processing

import (
	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
	documentpb "github.com/hexanecreations/docsai/backend/proto/documentpb"
	"github.com/pgvector/pgvector-go"
	"gorm.io/datatypes"
)

func mapChunks(
	documentID uuid.UUID,
	chunks []*documentpb.Chunk,
) []schema.DocumentChunk {

	result := make([]schema.DocumentChunk, 0, len(chunks))

	for _, chunk := range chunks {
		result = append(result, schema.DocumentChunk{
			ID:         uuid.New(),
			DocumentID: documentID,
			ChunkIndex: chunk.GetId(),
			Page:       chunk.GetPage(),
			Section:    chunk.GetSection(),
			Text:       chunk.GetText(),
			Metadata:   datatypes.JSON([]byte("{}")),
		})
	}

	return result
}

func mapEmbeddings(
	documentID uuid.UUID,
	chunks []schema.DocumentChunk,
	embeddings []*documentpb.Embedding,
) []schema.DocumentEmbedding {

	result := make([]schema.DocumentEmbedding, 0, len(embeddings))

	chunkMap := make(map[int32]uuid.UUID)
	for _, chunk := range chunks {
		chunkMap[chunk.ChunkIndex] = chunk.ID
	}

	for _, embedding := range embeddings {
		chunkID, ok := chunkMap[embedding.GetChunkId()]
		if !ok {
			continue
		}

		result = append(result, schema.DocumentEmbedding{
			ID:              uuid.New(),
			DocumentID:      documentID,
			DocumentChunkID: chunkID,
			Model:           embedding.GetModel(),
			Vector:          pgvector.NewVector(embedding.GetVector()),
		})
	}

	return result
}
