package bootstrap

import (
	grpcserver "github.com/hexanecreations/docsai/backend/internal/grpc"
	"github.com/hexanecreations/docsai/backend/internal/processing"
)

func (a *Application) loadGRPC() error {

	processingRepository := processing.NewRepository(
		a.DB,
		a.Logger,
	)

	callbackService := processing.NewService(
		a.Logger,
		processingRepository,
	)

	callbackServer := processing.NewServer(
		callbackService,
	)

	a.GRPCServer = grpcserver.New(
		a.Logger,
		callbackServer,
	)

	return nil
}
