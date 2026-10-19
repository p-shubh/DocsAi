package document

import (
	"github.com/google/uuid"
	"github.com/hexanecreations/docsai/backend/internal/common/enums"
)

type ListDocumentsFilter struct {
	Limit  int
	Offset int

	Search string

	Status *enums.DocumentStatus

	FolderID *uuid.UUID
}
