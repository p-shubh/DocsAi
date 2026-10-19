package documentdto

import "mime/multipart"

type UploadDocumentRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`

	FolderID string `form:"folder_id"`

	Tags []string `form:"tags"`

	Description string `form:"description"`
}