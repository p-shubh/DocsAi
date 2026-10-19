package bootstrap

import "github.com/hexanecreations/docsai/backend/internal/config"

func (a *Application) loadConfig() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	a.Config = cfg
	return nil
}
