package bootstrap

import "github.com/hexanecreations/docsai/backend/internal/storage"

func (a *Application) loadStorage() error {

	client, err := storage.New(a.Config)
	if err != nil {
		return err
	}

	a.Storage = storage.NewService(
		client,
		a.Config.MinioBucket,
	)

	a.Logger.Info("MinIO connected")

	return nil
}
