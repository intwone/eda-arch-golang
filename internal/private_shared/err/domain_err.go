package err

import (
	"errors"
	"fmt"
)

var (
	UnauthorizedError = errors.New("unauthorized_error")
)

func NewResourceAlreadyTakenError(resource string) error {
	err := fmt.Sprintf("%s_already_taken_error", resource)
	return errors.New(err)
}

func NewInvalidValueError(value string) error {
	err := fmt.Sprintf("%s_invalid_error", value)
	return errors.New(err)
}
