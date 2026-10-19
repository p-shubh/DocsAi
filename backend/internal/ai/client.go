package ai

import (
	"context"

	"github.com/hexanecreations/docsai/backend/proto/documentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type service struct {
	client documentpb.DocumentAIClient
	conn   *grpc.ClientConn
}

func NewService(host, port string) (Service, error) {

	conn, err := grpc.NewClient(
		host+":"+port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}

	return &service{
		client: documentpb.NewDocumentAIClient(conn),
		conn:   conn,
	}, nil
}


func (s *service) ProcessDocument(
	ctx context.Context,
	req *documentpb.ProcessDocumentRequest,
) (*documentpb.ProcessDocumentResponse, error) {
	return s.client.ProcessDocument(ctx, req)
}

func (s *service) EmbedQuery(
	ctx context.Context,
	text string,
) ([]float32, string, error) {
	resp, err := s.client.EmbedQuery(
		ctx,
		&documentpb.EmbedQueryRequest{
			Text: text,
		},
	)
	if err != nil {
		return nil, "", err
	}

	return resp.Vector, resp.Model, nil
}

func (s *service) GenerateAnswer(
	ctx context.Context,
	question string,
	chunks []*documentpb.ContextChunk,
) (string, string, error) {
	resp, err := s.client.GenerateAnswer(
		ctx,
		&documentpb.GenerateAnswerRequest{
			Question: question,
			Context:  chunks,
		},
	)
	if err != nil {
		return "", "", err
	}

	return resp.Answer, resp.Model, nil
}

func (s *service) GenerateAnswerStream(
	ctx context.Context,
	question string,
	chunks []*documentpb.ContextChunk,
) (documentpb.DocumentAI_GenerateAnswerStreamClient, error) {
	return s.client.GenerateAnswerStream(
		ctx,
		&documentpb.GenerateAnswerRequest{
			Question: question,
			Context:  chunks,
		},
	)
}

func (s *service) Close() error {
	return s.conn.Close()
}