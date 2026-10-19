package ai

import (
	"context"

	"github.com/hexanecreations/docsai/backend/proto/documentpb"
)

type Service interface {
	ProcessDocument(
		ctx context.Context,
		req *documentpb.ProcessDocumentRequest,
	) (*documentpb.ProcessDocumentResponse, error)

	EmbedQuery(
		ctx context.Context,
		text string,
	) ([]float32, string, error)

	GenerateAnswer(
		ctx context.Context,
		question string,
		chunks []*documentpb.ContextChunk,
	) (answer string, model string, err error)

	GenerateAnswerStream(
		ctx context.Context,
		question string,
		chunks []*documentpb.ContextChunk,
	) (documentpb.DocumentAI_GenerateAnswerStreamClient, error)
}