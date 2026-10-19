package enums

type DocumentStatus string

const (
	DocumentStatusUploaded   DocumentStatus = "uploaded"
	DocumentStatusQueued     DocumentStatus = "queued"
	DocumentStatusProcessing DocumentStatus = "processing"
	DocumentStatusCompleted  DocumentStatus = "completed"
	DocumentStatusFailed     DocumentStatus = "failed"
)