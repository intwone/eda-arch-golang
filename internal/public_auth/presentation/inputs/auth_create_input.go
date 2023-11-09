package inputs

import (
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
)

type AuthCreateInput struct {
	Value contactValueObject.Email `json:"value" validate:"required,email"`
}
