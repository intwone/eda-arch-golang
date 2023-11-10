package events

import (
	"time"

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
	birthdate  time.Time
	personName string
	userID     uuid.UUID
	companyID  uuid.UUID
}

func NewUserCreatedEvent(email contactValueObject.Email, cpf personValueObject.Cpf, birthdate time.Time, personName string, userID uuid.UUID, companyID uuid.UUID) *UserCreatedEvent {
	p := UserCreatedEvent{
		email:      email,
		cpf:        cpf,
		birthdate:  birthdate,
		personName: personName,
		userID:     userID,
		companyID:  companyID,
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

func (e *UserCreatedEvent) GetBirthdate() time.Time {
	return e.birthdate
}

func (e *UserCreatedEvent) GetUserID() uuid.UUID {
	return e.userID
}

func (e *UserCreatedEvent) GetCompanyID() uuid.UUID {
	return e.companyID
}
