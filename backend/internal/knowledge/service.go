package knowledge

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/hexanecreations/docsai/backend/internal/ai"
	"github.com/hexanecreations/docsai/backend/internal/document"
	knowledgedto "github.com/hexanecreations/docsai/backend/internal/knowledge/dto"
	"github.com/hexanecreations/docsai/backend/proto/documentpb"
)

const (
	topK           = 5
	requestTimeout = 30 * time.Second
)

var ErrNoRelevantContext = errors.New("no relevant context found")

type Service interface {
	AskQuestion(
		ctx context.Context,
		question string,
		documentID *uuid.UUID,
	) (*knowledgedto.AskQuestionResponse, error)

	AskQuestionStream(
		ctx context.Context,
		question string,
		documentID *uuid.UUID,
		onSources func([]knowledgedto.SourceCitation),
		onDelta func(string),
	) (model string, err error)
}

type gormService struct {
	documentRepo document.Repository
	ai           ai.Service
	logger       *slog.Logger
}

func NewService(
	documentRepo document.Repository,
	ai ai.Service,
	logger *slog.Logger,
) Service {
	return &gormService{
		documentRepo: documentRepo,
		ai:           ai,
		logger:       logger,
	}
}

// retrieve embeds the question, searches for relevant chunks (optionally
// scoped to a single document), and returns both the gRPC context payload
// and the HTTP-facing source citations derived from the same chunks.
func (s *gormService) retrieve(
	ctx context.Context,
	question string,
	documentID *uuid.UUID,
) (pbContext []*documentpb.ContextChunk, sources []knowledgedto.SourceCitation, err error) {

	vector, _, err := s.ai.EmbedQuery(ctx, question)
	if err != nil {
		return nil, nil, fmt.Errorf("embedding question: %w", err)
	}

	chunks, err := s.documentRepo.SearchRelevantChunks(ctx, vector, topK, documentID)
	if err != nil {
		return nil, nil, fmt.Errorf("searching relevant chunks: %w", err)
	}

	pbContext = make([]*documentpb.ContextChunk, 0, len(chunks))
	sources = make([]knowledgedto.SourceCitation, 0, len(chunks))

	for _, chunk := range chunks {
		pbContext = append(pbContext, &documentpb.ContextChunk{
			ChunkId:      chunk.ChunkID.String(),
			Text:         chunk.Text,
			DocumentName: chunk.DocumentName,
			Page:         chunk.Page,
			Section:      chunk.Section,
		})
		sources = append(sources, knowledgedto.SourceCitation{
			DocumentName: chunk.DocumentName,
			Page:         chunk.Page,
			Section:      chunk.Section,
			ChunkID:      chunk.ChunkID.String(),
		})
	}

	return pbContext, sources, nil
}

func (s *gormService) AskQuestion(
	ctx context.Context,
	question string,
	documentID *uuid.UUID,
) (*knowledgedto.AskQuestionResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	pbContext, sources, err := s.retrieve(ctx, question, documentID)
	if err != nil {
		return nil, err
	}

	if len(pbContext) == 0 {
		return nil, ErrNoRelevantContext
	}

	answer, model, err := s.ai.GenerateAnswer(ctx, question, pbContext)
	if err != nil {
		return nil, fmt.Errorf("generating answer: %w", err)
	}

	return &knowledgedto.AskQuestionResponse{
		Answer:  answer,
		Model:   model,
		Sources: sources,
	}, nil
}

func (s *gormService) AskQuestionStream(
	ctx context.Context,
	question string,
	documentID *uuid.UUID,
	onSources func([]knowledgedto.SourceCitation),
	onDelta func(string),
) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	pbContext, sources, err := s.retrieve(ctx, question, documentID)
	if err != nil {
		return "", err
	}

	if len(pbContext) == 0 {
		onSources(nil)
		onDelta("I don't have any relevant information to answer that question.")
		return "", nil
	}

	onSources(sources)

	stream, err := s.ai.GenerateAnswerStream(ctx, question, pbContext)
	if err != nil {
		return "", fmt.Errorf("generating answer: %w", err)
	}

	var model string

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("streaming answer: %w", err)
		}
		if chunk.Done {
			model = chunk.Model
			break
		}
		onDelta(chunk.TextDelta)
	}

	return model, nil
}
