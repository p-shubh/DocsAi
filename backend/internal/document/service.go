package document

import (
	"context"

	"github.com/google/uuid"

	documentdto "github.com/hexanecreations/docsai/backend/internal/document/dto"
)

type Service interface {
	UploadDocument(
		ctx context.Context,
		request documentdto.UploadDocumentRequest,
	) (*documentdto.UploadDocumentResponse, error)

	GetDocument(
		ctx context.Context,
		id uuid.UUID,
	) (*documentdto.GetDocumentResponse, error)

	GetDocumentFile(
		ctx context.Context,
		id uuid.UUID,
	) (*documentdto.DocumentFileResponse, error)

	GetDocumentChunks(
		ctx context.Context,
		id uuid.UUID,
	) (*documentdto.ListChunksResponse, error)

	ListDocuments(
		ctx context.Context,
		filter ListDocumentsFilter,
	) (*documentdto.ListDocumentsResponse, error)

	RenameDocument(
		ctx context.Context,
		id uuid.UUID,
		request documentdto.RenameDocumentRequest,
	) (*documentdto.DocumentResponse, error)

	DeleteDocument(
		ctx context.Context,
		id uuid.UUID,
	) (*documentdto.DeleteDocumentResponse, error)
}
