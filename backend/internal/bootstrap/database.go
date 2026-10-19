package bootstrap

import (
	"github.com/hexanecreations/docsai/backend/internal/database"
)

func (a *Application) loadDatabase() error {

	db, err := database.New(a.Config)
	if err != nil {
		return err
	}

	if err := database.Migrate(db); err != nil {
		return err
	}

	a.DB = db

	a.Logger.Info("PostgreSQL connected")

	return nil
}
