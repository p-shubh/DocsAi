package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hexanecreations/docsai/backend/internal/bootstrap"
	"github.com/hexanecreations/docsai/backend/internal/middleware"
)

type Server struct {
	router *gin.Engine
	app    *bootstrap.Application
}

func New(app *bootstrap.Application) (*Server, error) {

	router := gin.New()

	router.Use(middleware.CORS(
		middleware.ParseOrigins(app.Config.CORSAllowedOrigins),
	))

	s := &Server{
		router: router,
		app:    app,
	}

	s.registerRoutes()

	return s, nil
}

func (s *Server) Run() error {

	go func() {
		_ = s.app.GRPCServer.Run(
			s.app.Config.GRPCServerPort,
		)
	}()

	return s.router.Run(
		":" + s.app.Config.AppPort,
	)
}
