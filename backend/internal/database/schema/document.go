package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
	"gorm.io/gorm"
)

type DocumentStatus string

const (
	DocumentStatusUploaded   DocumentStatus = "uploaded"
	DocumentStatusProcessing DocumentStatus = "processing"
	DocumentStatusCompleted  DocumentStatus = "completed"
	DocumentStatusFailed     DocumentStatus = "failed"
)

type Document struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OriginalName string `gorm:"type:text;not null"`
	StoredName   string `gorm:"type:text;not null"`

	MimeType  string `gorm:"type:text;not null"`
	Extension string `gorm:"type:varchar(20);not null"`

	Size int64 `gorm:"not null"`

	Bucket    string `gorm:"type:text;not null"`
	ObjectKey string `gorm:"type:text;uniqueIndex;not null"`

	Status enums.DocumentStatus `gorm:"type:varchar(20);not null;default:'uploaded'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
