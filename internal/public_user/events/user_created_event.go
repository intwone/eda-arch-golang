package events

import (
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
	uuid "github.com/satori/go.uuid"
)

var (
	UserCreatedEventName = "user_created_event"
)

type UserCreatedEvent struct {
	email      contactValueObject.Email
	cpf        personValueObject.Cpf
	personName string
	userID     uuid.UUID
}

func NewUserCreatedEvent(userID uuid.UUID, email contactValueObject.Email, cpf personValueObject.Cpf) *UserCreatedEvent {
	p := UserCreatedEvent{
		userID: userID,
		email:  email,
		cpf:    cpf,
	}
	return &p
}

func (e *UserCreatedEvent) GetName() string {
	return UserCreatedEventName
}

func (e *UserCreatedEvent) GetEmail() contactValueObject.Email {
	return e.email
}

func (e *UserCreatedEvent) GetCpf() personValueObject.Cpf {
	return e.cpf
}

func (e *UserCreatedEvent) GetPersonName() string {
	return e.personName
}

func (e *UserCreatedEvent) GetUserID() uuid.UUID {
	return e.userID
}
