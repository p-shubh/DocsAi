package processing

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
	"gorm.io/gorm"
)

type gormRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewRepository(
	db *gorm.DB,
	logger *slog.Logger,
) Repository {
	return &gormRepository{
		db:     db,
		logger: logger,
	}
}

func (r *gormRepository) SaveProcessingResult(
	ctx context.Context,
	documentID uuid.UUID,
	status enums.DocumentStatus,
	chunks []schema.DocumentChunk,
	embeddings []schema.DocumentEmbedding,
) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := r.saveChunksTx(tx, chunks); err != nil {
			return err
		}

		if err := r.saveEmbeddingsTx(tx, embeddings); err != nil {
			return err
		}

		if err := r.updateDocumentStatusTx(tx, documentID, status); err != nil {
			return err
		}

		return nil
	})
}

func (r *gormRepository) saveChunks(
	ctx context.Context,
	chunks []schema.DocumentChunk,
) error {
	if len(chunks) == 0 {
		return nil
	}

	return r.db.
		WithContext(ctx).
		Create(&chunks).
		Error
}

func (r *gormRepository) saveEmbeddings(
	ctx context.Context,
	embeddings []schema.DocumentEmbedding,
) error {
	if len(embeddings) == 0 {
		return nil
	}

	return r.db.
		WithContext(ctx).
		Create(&embeddings).
		Error
}

func (r *gormRepository) updateDocumentStatus(
	ctx context.Context,
	documentID uuid.UUID,
	status enums.DocumentStatus,
) error {
	return r.db.
		WithContext(ctx).
		Model(&schema.Document{}).
		Where("id = ?", documentID).
		Update("status", status).
		Error
}

func (r *gormRepository) saveChunksTx(
	tx *gorm.DB,
	chunks []schema.DocumentChunk,
) error {
	if len(chunks) == 0 {
		return nil
	}

	return tx.CreateInBatches(chunks, 100).Error
}

func (r *gormRepository) saveEmbeddingsTx(
	tx *gorm.DB,
	embeddings []schema.DocumentEmbedding,
) error {
	if len(embeddings) == 0 {
		return nil
	}

	return tx.CreateInBatches(embeddings, 100).Error
}

func (r *gormRepository) updateDocumentStatusTx(
	tx *gorm.DB,
	documentID uuid.UUID,
	status enums.DocumentStatus,
) error {
	return tx.Model(&schema.Document{}).
		Where("id = ?", documentID).
		Update("status", status).
		Error
}