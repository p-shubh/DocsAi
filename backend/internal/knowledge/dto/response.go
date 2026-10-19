package knowledgedto

type SourceCitation struct {
	DocumentName string `json:"document_name"`

	Page int32 `json:"page"`

	Section string `json:"section"`

	ChunkID string `json:"chunk_id"`
}

type AskQuestionResponse struct {
	Answer string `json:"answer"`

	Model string `json:"model"`

	Sources []SourceCitation `json:"sources"`
}
