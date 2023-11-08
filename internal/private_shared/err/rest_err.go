package err

import "net/http"

type RestError struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewCause(field string, message string) *Cause {
	return &Cause{
		Field:   field,
		Message: message,
	}
}

func NewConflictError(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Err:     "conflict",
		Code:    http.StatusConflict,
		Causes:  causes,
	}
}

func NewBadRequestError(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
