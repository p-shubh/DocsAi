package database

import (
	"github.com/hexanecreations/docsai/backend/internal/database/schema"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&schema.Document{},
		&schema.DocumentChunk{},
		&schema.DocumentEmbedding{},
	)
}
