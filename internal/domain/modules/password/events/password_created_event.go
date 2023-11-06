package events

import "github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"

var (
	PasswordCreatedEventName = "password_created_event"
)

type PasswordCreatedEvent struct {
	password entities.PasswordEntity
	unhash   string
}

func NewPasswordCreatedEvent(password entities.PasswordEntity, unhash string) *PasswordCreatedEvent {
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
