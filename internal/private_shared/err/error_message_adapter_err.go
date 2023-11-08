package err

import (
	"errors"
	"fmt"
)

type ErrorType string

var ErrorMessageToResponseMap = map[string]ErrorType{
	"record not found": "not found",
}

func ParseErrorResponse(typeErr string) (*ErrorType, error) {
	result, ok := ErrorMessageToResponseMap[typeErr]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Invalid conversion from ErrorMessage %s to ErrorType", typeErr))
	}
	return &result, nil
}
