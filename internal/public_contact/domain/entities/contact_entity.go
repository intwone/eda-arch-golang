package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ContactStatus string

const (
	ContactCreated  ContactStatus = "created"
	ContactVerified ContactStatus = "verified"
	ContactAccepted ContactStatus = "accepted"
	ContactRejected ContactStatus = "rejected"
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
	VerifiedAt *time.Time
	AcceptedAt *time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
}

func NewContactEntity(value string, kind ContactKind, userID uuid.UUID) *ContactEntity {
	contact := ContactEntity{
		ID:        uuid.NewV4(),
		Status:    ContactCreated,
		Kind:      kind,
		Value:     value,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
	}
	return &contact
}

func (c *ContactEntity) GetID() uuid.UUID {
	return c.ID
}

func (c *ContactEntity) GetUserID() uuid.UUID {
	return c.UserID
}

func (c *ContactEntity) SetStatus(status ContactStatus) {
	c.Status = status
	c.update()
}

func (c *ContactEntity) SetVerified(date time.Time) {
	c.VerifiedAt = &date
	c.update()
}

func (c *ContactEntity) update() {
	now := time.Now()
	c.UpdatedAt = now
}
