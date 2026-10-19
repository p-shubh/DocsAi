package knowledgedto

type AskQuestionRequest struct {
	Question string `json:"question" binding:"required"`

	DocumentID *string `json:"document_id,omitempty"`
}
