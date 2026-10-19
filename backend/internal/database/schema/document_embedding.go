package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type DocumentEmbedding struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	DocumentID uuid.UUID `gorm:"type:uuid;not null;index"`

	DocumentChunkID uuid.UUID `gorm:"type:uuid;not null;index"`

	Model string `gorm:"type:varchar(255);not null"`

	Vector pgvector.Vector `gorm:"type:vector(1024);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}