package events

import (
	contactEntities "github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
)

var (
	AuthenticatedEventName = "authenticated_event"
)

type AuthenticatedEvent struct {
	user    userEntities.UserEntity
	contact contactEntities.ContactEntity
}

func NewAuthenticatedEvent(user userEntities.UserEntity, contact contactEntities.ContactEntity) *AuthenticatedEvent {
	p := AuthenticatedEvent{
		user:    user,
		contact: contact,
	}
	return &p
}

func (e *AuthenticatedEvent) GetName() string {
	return AuthenticatedEventName
}

func (e *AuthenticatedEvent) GetUser() userEntities.UserEntity {
	return e.user
}

func (e *AuthenticatedEvent) GetContactID() contactEntities.ContactEntity {
	return e.contact
}
