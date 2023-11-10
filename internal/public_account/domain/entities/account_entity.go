package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type AccountKind string

const (
	AccountB2B   AccountKind = "b2b"
	AccountB2C   AccountKind = "b2c"
	AccountB2B2C AccountKind = "b2b2c"
)

type AccountEntity struct {
	ID        uuid.UUID
	Kind      AccountKind
	CreatedAt time.Time
	UpdatedAt time.Time
	CohortAt  *time.Time
}

func NewAccountEntity(kind AccountKind) *AccountEntity {
	account := AccountEntity{
		ID:        uuid.NewV4(),
		Kind:      kind,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &account
}

func (a *AccountEntity) GetID() uuid.UUID {
	return a.ID
}

func (a *AccountEntity) GetKind() AccountKind {
	return a.Kind
}

func (a *AccountEntity) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *AccountEntity) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a *AccountEntity) GetCohortAt() *time.Time {
	return a.CohortAt
}

func (a *AccountEntity) update() {
	now := time.Now()
	a.UpdatedAt = now
}
