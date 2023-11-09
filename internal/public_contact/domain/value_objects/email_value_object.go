package valueobjects

import (
	"encoding/json"
	"regexp"

	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
)

type Email struct {
	Value string
}

func NewEmail(value string) (*Email, error) {
	if !isValidEmail(value) {
		return nil, err.NewInvalidValueError("email")
	}
	email := Email{Value: value}
	return &email, nil
}

func isValidEmail(value string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, value)
	return match
}

func (e *Email) UnmarshalJSON(data []byte) error {
	var email string
	if err := json.Unmarshal(data, &email); err != nil {
		return err
	}
	if !isValidEmail(email) {
		return err.NewInvalidValueError("email")
	}
	e.Value = email
	return nil
}
