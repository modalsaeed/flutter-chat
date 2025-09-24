package dto

type APIResponse struct {
	Success bool        `json:success`
	Code    int         `json:"code"`            // HTTP status code (e.g., 200, 400, 401)
	Data    interface{} `json:"data,omitempty"`  // Present on success
	Error   string      `json:"error,omitempty"` // Present on error
}

type PaginatedData[T any] struct {
	Items  []T `json:"items"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
