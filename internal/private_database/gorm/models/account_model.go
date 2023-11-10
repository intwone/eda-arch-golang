package models

import (
	"time"

	accountEntities "github.com/intwone/eda-arch-golang/internal/public_account/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AccountModel struct {
	ID        uuid.UUID                   `gorm:"column:account_id;type:uuid;primaryKey;unique;not null"`
	Kind      accountEntities.AccountKind `gorm:"column:kind;type:text;not null"`
	CreatedAt time.Time                   `gorm:"column:created_at;type:timestamptz(6);not null"`
	UpdatedAt time.Time                   `gorm:"column:updated_at;type:timestamptz(6);not null"`
	CohortAt  *time.Time                  `gorm:"column:updated_at;type:timestamptz(6)"`
}

func MigrateAccount(db *gorm.DB) error {
	err := db.Table("account").AutoMigrate(&AccountModel{})
	if err != nil {
		return err
	}
	return nil
}

func (ContactModel) TableName() string {
	return "account"
}
