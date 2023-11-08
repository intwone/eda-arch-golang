package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type PasswordKind string

const (
	PasswordEmail PasswordKind = "email"
	PasswordPhone PasswordKind = "phone"
)

type PasswordEntity struct {
	ID        uuid.UUID
	Kind      PasswordKind
	Hash      string
	IsActive  bool
	CreatedAt time.Time
	ContactID uuid.UUID
}

func NewPasswordEntity(kind PasswordKind, hash string, contactID uuid.UUID) *PasswordEntity {
	password := PasswordEntity{
		ID:        uuid.NewV4(),
		Kind:      kind,
		Hash:      hash,
		IsActive:  true,
		CreatedAt: time.Now(),
		ContactID: contactID,
	}
	return &password
}

func (p *PasswordEntity) GetID() uuid.UUID {
	return p.ID
}

func (p *PasswordEntity) GetHash() string {
	return p.Hash
}

func (p *PasswordEntity) GetContactID() uuid.UUID {
	return p.ContactID
}
