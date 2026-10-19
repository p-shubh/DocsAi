package processing

import (
	"context"

	"github.com/hexanecreations/docsai/backend/proto/documentpb"
)

type Server struct {
	documentpb.UnimplementedDocumentCallbackServer

	service Service
}

func NewServer(
	service Service,
) *Server {

	return &Server{
		service: service,
	}
}

func (s *Server) SaveProcessingResult(
	ctx context.Context,
	request *documentpb.SaveProcessingResultRequest,
) (*documentpb.SaveProcessingResultResponse, error) {

	return s.service.SaveProcessingResult(
		ctx,
		request,
	)

}
