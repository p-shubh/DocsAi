package bootstrap

import (
	"github.com/hexanecreations/docsai/backend/internal/ai"
)

func (a *Application) loadAI() error {

	a.Logger.Info(
		"Initializing AI client",
		"host", a.Config.GRPCHost,
		"port", a.Config.GRPCPort,
	)

	service, err := ai.NewService(
		a.Config.GRPCHost,
		a.Config.GRPCPort,
	)
	if err != nil {
		return err
	}

	a.AI = service

	a.Logger.Info("AI gRPC client initialized")

	return nil
}
