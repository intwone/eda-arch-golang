package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserStatus string

const (
	UserCreated  UserStatus = "created"
	UserVerified UserStatus = "verified"
	UserPending  UserStatus = "pending"
	UserAccepted UserStatus = "accepted"
	UserRejected UserStatus = "rejected"
)

type UserEntity struct {
	ID           uuid.UUID
	Status       UserStatus
	CreatedAt    time.Time
	VerifiedAt   *time.Time
	PendingAt    *time.Time
	WaitlistedAt *time.Time
	SuspendedAt  *time.Time
	MemberAt     *time.Time
	ChurnedAt    *time.Time
	ArchivedAt   *time.Time
	SettledAt    *time.Time
	UpdatedAt    time.Time
	CrmSyncedAt  *time.Time
	IsCrmSynced  bool
}

func NewUserEntity() *UserEntity {
	user := UserEntity{
		ID:          uuid.NewV4(),
		Status:      UserCreated,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsCrmSynced: false,
	}
	return &user
}

func (u *UserEntity) GetID() uuid.UUID {
	return u.ID
}

func (u *UserEntity) SetStatus(status UserStatus) {
	u.Status = status
	u.update()
}

func (u *UserEntity) SetVerified(date time.Time) {
	u.VerifiedAt = &date
	u.update()
}

func (u *UserEntity) update() {
	now := time.Now()
	u.UpdatedAt = now
}
