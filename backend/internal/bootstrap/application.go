package bootstrap

import (
	"log/slog"

	"github.com/hexanecreations/docsai/backend/internal/ai"
	"github.com/hexanecreations/docsai/backend/internal/config"
	"github.com/hexanecreations/docsai/backend/internal/grpc"
	"github.com/hexanecreations/docsai/backend/internal/storage"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Config        *config.Config
	Logger        *slog.Logger
	DB            *gorm.DB
	Redis         *redis.Client
	StorageClient *minio.Client
	AI            ai.Service
	Storage       storage.Service
	GRPCServer    *grpc.Server
}

func New() (*Application, error) {

	application := &Application{}

	if err := application.loadConfig(); err != nil {
		return nil, err
	}

	if err := application.loadLogger(); err != nil {
		return nil, err
	}

	if err := application.loadDatabase(); err != nil {
		return nil, err
	}

	if err := application.loadRedis(); err != nil {
		return nil, err
	}

	if err := application.loadStorage(); err != nil {
		return nil, err
	}

	if err := application.loadAI(); err != nil {
		return nil, err
	}

	if err := application.loadGRPC(); err != nil {
		return nil, err
	}

	return application, nil
}
