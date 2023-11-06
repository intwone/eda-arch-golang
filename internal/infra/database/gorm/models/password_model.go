package models

import (
	"time"

	"github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PasswordModel struct {
	ID        uuid.UUID             `gorm:"column:password_id;type:uuid;primaryKey;unique;not null"`
	Kind      entities.PasswordKind `gorm:"column:kind;type:text;not null"`
	Hash      string                `gorm:"column:hash;type:text;not null"`
	IsActive  bool                  `gorm:"column:is_active;type:boolean;not null"`
	CreatedAt time.Time             `gorm:"column:created_at;type:timestamptz(6);not null"`
	ContactID uuid.UUID             `gorm:"column:contact_id;foreignKey:Contact;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Contact   ContactModel          `gorm:"foreignKey:ContactID"`
}

func MigratePassword(db *gorm.DB) error {
	err := db.Table("password").AutoMigrate(&PasswordModel{})
	if err != nil {
		return err
	}
	return nil
}

func (PasswordModel) TableName() string {
	return "password"
}
