package document

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/google/uuid"

	"github.com/hexanecreations/docsai/backend/internal/ai"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	"github.com/hexanecreations/docsai/backend/internal/config"
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
	documentdto "github.com/hexanecreations/docsai/backend/internal/document/dto"
	"github.com/hexanecreations/docsai/backend/internal/storage"
	"github.com/hexanecreations/docsai/backend/proto/documentpb"
)

const filePresignExpiry = 15 * time.Minute

type gormService struct {
	repo    Repository
	storage storage.Service
	ai      ai.Service
	logger  *slog.Logger
	config  *config.Config
}

func NewService(

	repo Repository,
	storage storage.Service,
	ai ai.Service,
	logger *slog.Logger,
	config *config.Config,

) Service {
	return &gormService{
		repo:    repo,
		storage: storage,
		ai:      ai,
		logger:  logger,
		config:  config,
	}
}

func (s *gormService) GetDocument(
	ctx context.Context,
	id uuid.UUID,
) (*documentdto.GetDocumentResponse, error) {

	document, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &documentdto.GetDocumentResponse{
		ID: document.ID,

		OriginalName: document.OriginalName,

		Size: document.Size,

		MimeType: document.MimeType,

		Status: string(document.Status),

		PageCount: 0,

		SummaryAvailable: false,

		OCRCompleted: false,

		EmbeddingGenerated: false,

		UploadedAt: document.CreatedAt,
	}, nil
}

func (s *gormService) GetDocumentFile(
	ctx context.Context,
	id uuid.UUID,
) (*documentdto.DocumentFileResponse, error) {

	document, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	fileURL, err := s.storage.GetPresignedURL(
		ctx,
		document.ObjectKey,
		filePresignExpiry,
	)
	if err != nil {
		return nil, err
	}

	return &documentdto.DocumentFileResponse{
		FileURL:      fileURL,
		MimeType:     document.MimeType,
		OriginalName: document.OriginalName,
	}, nil
}

func (s *gormService) GetDocumentChunks(
	ctx context.Context,
	id uuid.UUID,
) (*documentdto.ListChunksResponse, error) {

	chunks, err := s.repo.ListChunksByDocumentID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &documentdto.ListChunksResponse{
		Chunks: make([]documentdto.ChunkResponse, 0, len(chunks)),
	}

	for _, chunk := range chunks {
		response.Chunks = append(response.Chunks, documentdto.ChunkResponse{
			ChunkIndex: chunk.ChunkIndex,
			Page:       chunk.Page,
			Section:    chunk.Section,
			Text:       chunk.Text,
		})
	}

	return response, nil
}

func (s *gormService) ListDocuments(
	ctx context.Context,
	filter ListDocumentsFilter,
) (*documentdto.ListDocumentsResponse, error) {

	documents, total, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	response := &documentdto.ListDocumentsResponse{
		Documents: make([]documentdto.DocumentResponse, 0, len(documents)),
		Total:     total,
		Page:      (filter.Offset / filter.Limit) + 1,
		Limit:     filter.Limit,
	}

	for _, document := range documents {
		response.Documents = append(
			response.Documents,
			*ToDocumentResponse(&document),
		)
	}

	return response, nil
}

func (s *gormService) RenameDocument(
	ctx context.Context,
	id uuid.UUID,
	request documentdto.RenameDocumentRequest,
) (*documentdto.DocumentResponse, error) {

	document, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	document.OriginalName = request.Name

	if err := s.repo.Update(ctx, document); err != nil {
		return nil, err
	}

	return ToDocumentResponse(document), nil
}

func (s *gormService) DeleteDocument(
	ctx context.Context,
	id uuid.UUID,
) (*documentdto.DeleteDocumentResponse, error) {

	document, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(
		ctx,
		document.ObjectKey,
	); err != nil {
		return nil, err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return nil, err
	}

	return &documentdto.DeleteDocumentResponse{
		Message: "Document deleted successfully",
	}, nil
}

func (s *gormService) UploadDocument(
	ctx context.Context,
	req documentdto.UploadDocumentRequest,
) (*documentdto.UploadDocumentResponse, error) {

	documentID := uuid.New()

	extension := filepath.Ext(req.File.Filename)

	objectKey := fmt.Sprintf(
		"%s%s",
		documentID.String(),
		extension,
	)
	_, err := s.storage.Upload(
		ctx,
		req.File,
		objectKey,
	)
	if err != nil {
		return nil, err
	}

	document := &schema.Document{
		ID: documentID,

		OriginalName: req.File.Filename,
		StoredName:   objectKey,

		MimeType:  req.File.Header.Get("Content-Type"),
		Extension: extension,

		Size: req.File.Size,

		Bucket:    s.config.MinioBucket,
		ObjectKey: objectKey,

		Status: enums.DocumentStatusUploaded,
	}

	if err := s.repo.Create(ctx, document); err != nil {
		return nil, err
	}

	go func(document *schema.Document) {

		_, err := s.ai.ProcessDocument(
			context.Background(),
			&documentpb.ProcessDocumentRequest{
				DocumentId: document.ID.String(),
				Bucket:     document.Bucket,
				ObjectKey:  document.ObjectKey,
				MimeType:   document.MimeType,
			},
		)

		if err != nil {
			s.logger.Error(
				"failed to process document",
				"document_id", document.ID,
				"error", err,
			)
		}

	}(document)

	return &documentdto.UploadDocumentResponse{
		ID:           document.ID,
		OriginalName: document.OriginalName,
		Size:         document.Size,
		MimeType:     document.MimeType,
		Status:       string(document.Status),
		UploadedAt:   document.CreatedAt,
	}, nil
}
