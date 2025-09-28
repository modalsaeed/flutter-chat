package dto

type APIResponse struct {
	Code  int         `json:"code"`            // HTTP status code (e.g., 200, 400, 401)
	Data  interface{} `json:"data,omitempty"`  // Present on success
	Error *ErrorData  `json:"error,omitempty"` // Present on error
}

type PaginatedData[T any] struct {
	Items  []T `json:"items"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorData struct {
	Code    int          `json:"-"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields,omitempty"`
}
