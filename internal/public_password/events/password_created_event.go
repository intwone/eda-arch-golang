package events

import (
	passwordEntities "github.com/intwone/eda-arch-golang/internal/public_password/domain/entities"
)

var (
	PasswordCreatedEventName = "password_created_event"
)

type PasswordCreatedEvent struct {
	password passwordEntities.PasswordEntity
	unhash   string
}

func NewPasswordCreatedEvent(password passwordEntities.PasswordEntity, unhash string) *PasswordCreatedEvent {
	p := PasswordCreatedEvent{
		password: password,
		unhash:   unhash,
	}
	return &p
}

func (e *PasswordCreatedEvent) GetName() string {
	return PasswordCreatedEventName
}

func (e *PasswordCreatedEvent) GetUnhash() string {
	return e.unhash
}
