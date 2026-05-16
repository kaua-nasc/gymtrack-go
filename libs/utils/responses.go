package utils

// PaginatedResponse is a generic structure for all paginated API responses.
type PaginatedResponse[T any] struct {
	Data   []T    `json:"data"`
	Cursor string `json:"cursor,omitempty"`
}

// NewPaginatedResponse creates a new PaginatedResponse instance.
func NewPaginatedResponse[T any](data []T, cursor string) PaginatedResponse[T] {
	return PaginatedResponse[T]{
		Data:   data,
		Cursor: cursor,
	}
}

// ErrorResponse is a standard structure for API error messages.
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}
