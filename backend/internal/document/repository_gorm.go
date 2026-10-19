package document

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
	"github.com/pgvector/pgvector-go"
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
func (r *gormRepository) Create(
	ctx context.Context,
	document *schema.Document,
) error {
	return r.db.WithContext(ctx).
		Create(document).
		Error
}
func (r *gormRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*schema.Document, error) {

	var document schema.Document

	err := r.db.
		WithContext(ctx).
		First(&document, "id = ?", id).
		Error

	if err != nil {
		return nil, err
	}

	return &document, nil
}
func (r *gormRepository) Update(
	ctx context.Context,
	document *schema.Document,
) error {

	return r.db.
		WithContext(ctx).
		Save(document).
		Error
}
func (r *gormRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Delete(
			&schema.Document{},
			"id = ?",
			id,
		).
		Error
}

func (r *gormRepository) List(
	ctx context.Context,
	filter ListDocumentsFilter,
) ([]schema.Document, int64, error) {

	var (
		documents []schema.Document
		total     int64
	)

	query := r.db.
		WithContext(ctx).
		Model(&schema.Document{})

	if filter.Search != "" {
		query = query.Where(
			"original_name ILIKE ?",
			"%"+filter.Search+"%",
		)
	}

	if filter.Status != nil {
		query = query.Where(
			"status = ?",
			*filter.Status,
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err := query.
		Order("created_at DESC").
		Find(&documents).
		Error

	if err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}

func (r *gormRepository) SaveChunks(
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

func (r *gormRepository) SaveEmbeddings(
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

func (r *gormRepository) ListChunksByDocumentID(
	ctx context.Context,
	documentID uuid.UUID,
) ([]schema.DocumentChunk, error) {

	var chunks []schema.DocumentChunk

	err := r.db.
		WithContext(ctx).
		Where("document_id = ?", documentID).
		Order("chunk_index ASC").
		Find(&chunks).
		Error

	if err != nil {
		return nil, err
	}

	return chunks, nil
}

func (r *gormRepository) SearchRelevantChunks(
	ctx context.Context,
	embedding []float32,
	limit int,
	documentID *uuid.UUID,
) ([]RetrievedChunk, error) {
	var results []RetrievedChunk

	if limit <= 0 {
		limit = 5
	}

	queryVector := pgvector.NewVector(embedding)

	sql := `
		SELECT
			dc.id           AS chunk_id,
			dc.document_id  AS document_id,
			d.original_name AS document_name,
			dc.page         AS page,
			dc.section      AS section,
			dc.text         AS text,
			de.vector <=> ? AS distance
		FROM document_embeddings de
		JOIN document_chunks dc
			ON dc.id = de.document_chunk_id
		JOIN documents d
			ON d.id = dc.document_id
		WHERE de.deleted_at IS NULL
			AND dc.deleted_at IS NULL
			AND d.deleted_at IS NULL
	`
	args := []any{queryVector}

	if documentID != nil {
		sql += " AND dc.document_id = ?"
		args = append(args, *documentID)
	}

	sql += " ORDER BY de.vector <=> ? LIMIT ?"
	args = append(args, queryVector, limit)

	err := r.db.
		WithContext(ctx).
		Raw(sql, args...).
		Scan(&results).
		Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *gormRepository) UpdateStatus(
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
