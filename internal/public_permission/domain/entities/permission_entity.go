package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type PermissionStatus string

const (
	PermissionCreated  PermissionStatus = "created"
	PermissionArchived PermissionStatus = "archived"
)

type PermissionRole string

const (
	PermissionUser          PermissionRole = "user"
	PermissionAdministrator PermissionRole = "administrator"
)

type PermissionEntity struct {
	ID         uuid.UUID
	Status     PermissionStatus
	Role       PermissionRole
	CreatedAt  time.Time
	ArchivedAt *time.Time
	UpdatedAt  time.Time
	UserID     uuid.UUID
	CompanyID  uuid.UUID
	AccountID  uuid.UUID
}

func NewPermissionEntity(role PermissionRole, userID uuid.UUID, accountID uuid.UUID, companyID uuid.UUID) *PermissionEntity {
	permission := PermissionEntity{
		ID:        uuid.NewV4(),
		Status:    PermissionCreated,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		AccountID: accountID,
		CompanyID: companyID,
	}
	return &permission
}

func (p *PermissionEntity) GetID() uuid.UUID {
	return p.ID
}

func (p *PermissionEntity) GetStatus() PermissionStatus {
	return p.Status
}

func (p *PermissionEntity) GetRole() PermissionRole {
	return p.Role
}

func (p *PermissionEntity) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *PermissionEntity) GetArchivedAt() *time.Time {
	return p.ArchivedAt
}

func (p *PermissionEntity) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *PermissionEntity) GetUserID() uuid.UUID {
	return p.UserID
}

func (p *PermissionEntity) GetCompanyID() uuid.UUID {
	return p.CompanyID
}

func (p *PermissionEntity) GetAccountID() uuid.UUID {
	return p.AccountID
}

func (p *PermissionEntity) update() {
	now := time.Now()
	p.UpdatedAt = now
}
