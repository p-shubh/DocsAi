package pagination

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

type Result struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}