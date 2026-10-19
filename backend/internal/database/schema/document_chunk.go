package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DocumentChunk struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	DocumentID uuid.UUID `gorm:"type:uuid;not null;index"`

	ChunkIndex int32 `gorm:"not null"`

	Page int32 `gorm:"not null"`

	Section string `gorm:"type:text"`

	Text string `gorm:"type:text;not null"`

	Metadata datatypes.JSON `gorm:"type:jsonb"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
