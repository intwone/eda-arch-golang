package entities

import (
	"time"

	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
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
	Value      contactValueObject.Email
	IsActive   bool
	CreatedAt  time.Time
	VerifiedAt *time.Time
	AcceptedAt *time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
}

func NewContactEntity(value contactValueObject.Email, kind ContactKind, userID uuid.UUID) *ContactEntity {
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

func (c *ContactEntity) GetStatus() ContactStatus {
	return c.Status
}

func (c *ContactEntity) GetKind() ContactKind {
	return c.Kind
}

func (c *ContactEntity) GetValue() contactValueObject.Email {
	return c.Value
}

func (c *ContactEntity) GetIsActive() bool {
	return c.IsActive
}

func (c *ContactEntity) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *ContactEntity) GetVerifiedAt() *time.Time {
	return c.VerifiedAt
}

func (c *ContactEntity) GetAcceptedAt() *time.Time {
	return c.AcceptedAt
}

func (c *ContactEntity) GetUpdatedAt() time.Time {
	return c.UpdatedAt
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
