package inputs

import (
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
)

type AuthenticateInput struct {
	Email    contactValueObject.Email `json:"email" binding:"required,email"`
	Password string                   `json:"password" binding:"required,min=6,max=6"`
}
