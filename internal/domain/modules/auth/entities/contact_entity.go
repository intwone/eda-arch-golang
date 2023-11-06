package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ContactStatus string

const (
	Created  ContactStatus = "created"
	Verified ContactStatus = "verified"
	Accepted ContactStatus = "accepted"
	Rejected ContactStatus = "rejected"
)

type ContactKind string

const (
	ContactEmail ContactKind = "email"
	ContactPhone ContactKind = "phone"
)

type ContactEntity struct {
	ID         uuid.UUID
	Status     ContactStatus
	Kind       ContactKind
	Value      string
	IsActive   bool
	CreatedAt  time.Time
	VeriedAt   time.Time
	AcceptedAt time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
}

func NewContactEntity(value string, kind ContactKind, userID uuid.UUID) *ContactEntity {
	contact := ContactEntity{
		ID:         uuid.NewV4(),
		Status:     Created,
		Kind:       kind,
		Value:      value,
		IsActive:   true,
		CreatedAt:  time.Now(),
		VeriedAt:   time.Time{},
		AcceptedAt: time.Time{},
		UpdatedAt:  time.Now(),
		UserID:     userID,
	}
	return &contact
}
