package server

import (
	"github.com/hexanecreations/docsai/backend/internal/document"
	"github.com/hexanecreations/docsai/backend/internal/health"
	"github.com/hexanecreations/docsai/backend/internal/knowledge"
)

func (s *Server) registerRoutes() {

	api := s.router.Group("/api")

	v1 := api.Group("/v1")

	health.RegisterRoutes(v1)

	document.RegisterRoutes(
		v1,
		document.NewHandler(
			document.NewService(
				document.NewRepository(
					s.app.DB,
					s.app.Logger,
				),
				s.app.Storage,
				s.app.AI,
				s.app.Logger,
				s.app.Config,
			),
		),
	)

	knowledge.RegisterRoutes(
		v1,
		knowledge.NewHandler(
			knowledge.NewService(
				document.NewRepository(
					s.app.DB,
					s.app.Logger,
				),
				s.app.AI,
				s.app.Logger,
			),
		),
	)

}
