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

func (u *UserEntity) GetStatus() UserStatus {
	return u.Status
}

func (u *UserEntity) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *UserEntity) GetVerifiedAt() *time.Time {
	return u.VerifiedAt
}

func (u *UserEntity) GetPendingAt() *time.Time {
	return u.PendingAt
}

func (u *UserEntity) GetWaitlistedAt() *time.Time {
	return u.WaitlistedAt
}

func (u *UserEntity) GetSuspendedAt() *time.Time {
	return u.SuspendedAt
}

func (u *UserEntity) GetMemberAt() *time.Time {
	return u.MemberAt
}

func (u *UserEntity) GetChurnedAt() *time.Time {
	return u.ChurnedAt
}

func (u *UserEntity) GetArchivedAt() *time.Time {
	return u.ArchivedAt
}

func (u *UserEntity) GetSettledAt() *time.Time {
	return u.SettledAt
}

func (u *UserEntity) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *UserEntity) GetCrmSyncedAt() *time.Time {
	return u.CrmSyncedAt
}

func (u *UserEntity) GetIsCrmSynced() bool {
	return u.IsCrmSynced
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
