package models

import (
	"time"

	"github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ContactModel struct {
	ID         uuid.UUID              `gorm:"column:contact_id;type:uuid;primaryKey;unique;not null"`
	Status     entities.ContactStatus `gorm:"column:status;type:text;not null"`
	Kind       entities.ContactKind   `gorm:"column:kind;type:text;not null"`
	Value      string                 `gorm:"column:value;type:text;not null"`
	IsActive   bool                   `gorm:"column:is_active;type:boolean;not null"`
	CreatedAt  time.Time              `gorm:"column:created_at;type:timestamptz(6);not null"`
	VeriedAt   time.Time              `gorm:"column:verified_at;type:timestamptz(6)"`
	AcceptedAt time.Time              `gorm:"column:accepted_at;type:timestamptz(6)"`
	UpdatedAt  time.Time              `gorm:"column:updated_at;type:timestamptz(6);not null"`
	UserID     uuid.UUID              `gorm:"column:user_id;foreignKey:User;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User       UserModel              `gorm:"foreignKey:UserID"`
}

func MigrateContact(db *gorm.DB) error {
	err := db.Table("contact").AutoMigrate(&ContactModel{})
	if err != nil {
		return err
	}
	return nil
}

func (ContactModel) TableName() string {
	return "contact"
}
