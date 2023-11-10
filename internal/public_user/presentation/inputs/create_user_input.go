package inputs

import (
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
)

type CreateUserInput struct {
	Name      string                   `json:"name" validate:"required,mix:3"`
	Email     contactValueObject.Email `json:"email" validate:"required,email"`
	Cpf       personValueObject.Cpf    `json:"cpf" validate:"required,min:11,max:11"`
	Birthdate string                   `json:"birthdate" validate:"required"`
}
