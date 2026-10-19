package processing

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	documentpb "github.com/hexanecreations/docsai/backend/proto/documentpb"
)

type Service interface {
	SaveProcessingResult(
		ctx context.Context,
		request *documentpb.SaveProcessingResultRequest,
	) (*documentpb.SaveProcessingResultResponse, error)
}

type service struct {
	logger     *slog.Logger
	repository Repository
}

func NewService(
	logger *slog.Logger,
	repository Repository,
) Service {

	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) SaveProcessingResult(
	ctx context.Context,
	request *documentpb.SaveProcessingResultRequest,
) (*documentpb.SaveProcessingResultResponse, error) {

	result := request.GetResult()

	documentID, err := uuid.Parse(request.GetDocumentId())
	if err != nil {
		return nil, err
	}

	s.logger.Info(
		"received processing result",
		"document_id", request.GetDocumentId(),
		"page_count", result.GetPageCount(),
		"chunks", len(result.GetChunks()),
		"embeddings", len(result.GetEmbeddings()),
	)

	chunks := mapChunks(
		documentID,
		result.GetChunks(),
	)

	embeddings := mapEmbeddings(
		documentID,
		chunks,
		result.GetEmbeddings(),
	)

	if err := s.repository.SaveProcessingResult(
		ctx,
		documentID,
		enums.DocumentStatusCompleted,
		chunks,
		embeddings,
	); err != nil {
		return nil, err
	}

	s.logger.Info(
		"processing result saved",
		"document_id", documentID,
		"chunks_saved", len(chunks),
		"embeddings_saved", len(embeddings),
	)

	return &documentpb.SaveProcessingResultResponse{
		Success: true,
	}, nil

}
