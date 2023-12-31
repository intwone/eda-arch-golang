package models

import (
	"time"

	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID           uuid.UUID               `gorm:"column:user_id;type:uuid;primaryKey;unique;not null"`
	Status       userEntities.UserStatus `gorm:"column:status;type:text;not null"`
	CreatedAt    time.Time               `gorm:"column:created_at;type:timestamptz(6);not null"`
	VerifiedAt   *time.Time              `gorm:"column:verified_at;type:timestamptz(6)"`
	PendingAt    *time.Time              `gorm:"column:pending_at;type:timestamptz(6)"`
	WaitlistedAt *time.Time              `gorm:"column:waitlisted_at;type:timestamptz(6)"`
	SuspendedAt  *time.Time              `gorm:"column:suspended_at;type:timestamptz(6)"`
	MemberAt     *time.Time              `gorm:"column:member_at;type:timestamptz(6)"`
	ChurnedAt    *time.Time              `gorm:"column:churned_at;type:timestamptz(6)"`
	ArchivedAt   *time.Time              `gorm:"column:archived_at;type:timestamptz(6)"`
	SettledAt    *time.Time              `gorm:"column:settled_at;type:timestamptz(6)"`
	UpdatedAt    time.Time               `gorm:"column:updated_at;type:timestamptz(6);not null"`
	CrmSyncedAt  *time.Time              `gorm:"column:crm_synced_at;type:timestamptz(6)"`
	IsCrmSynced  bool                    `gorm:"column:is_crm_synced;type:boolean;not null"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.Table("user").AutoMigrate(&UserModel{})
	if err != nil {
		return err
	}
	return nil
}

func (UserModel) TableName() string {
	return "user"
}
