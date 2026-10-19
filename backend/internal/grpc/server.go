package grpc

import (
	"fmt"
	"log/slog"
	"net"

	documentpb "github.com/hexanecreations/docsai/backend/proto/documentpb"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	logger *slog.Logger
}

func New(
	logger *slog.Logger,
	callbackServer documentpb.DocumentCallbackServer,
) *Server {

	grpcServer := grpc.NewServer()

	documentpb.RegisterDocumentCallbackServer(
		grpcServer,
		callbackServer,
	)

	return &Server{
		server: grpcServer,
		logger: logger,
	}
}

func (s *Server) Run(
	port string,
) error {

	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%s", port),
	)

	if err != nil {
		return err
	}

	s.logger.Info(
		"gRPC callback server started",
		"port",
		port,
	)

	return s.server.Serve(listener)
}
