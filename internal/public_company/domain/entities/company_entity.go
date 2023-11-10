package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CompanyStatus string

const (
	CompanyCreated CompanyStatus = "created"
)

type CompanyEntity struct {
	ID        uuid.UUID
	Name      string
	Status    CompanyStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCompanyEntity(name string) *CompanyEntity {
	company := CompanyEntity{
		ID:        uuid.NewV4(),
		Name:      name,
		Status:    CompanyCreated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &company
}

func (c *CompanyEntity) GetID() uuid.UUID {
	return c.ID
}

func (c *CompanyEntity) GetName() string {
	return c.Name
}

func (c *CompanyEntity) GetStatus() CompanyStatus {
	return c.Status
}

func (c *CompanyEntity) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *CompanyEntity) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

func (c *CompanyEntity) update() {
	now := time.Now()
	c.UpdatedAt = now
}
