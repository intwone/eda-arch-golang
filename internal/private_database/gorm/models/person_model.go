package models

import (
	"time"

	personEntities "github.com/intwone/eda-arch-golang/internal/public_person/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PersonModel struct {
	ID                 uuid.UUID                   `gorm:"column:person_id;type:uuid;primaryKey;unique;not null"`
	Status             personEntities.PersonStatus `gorm:"column:status;type:text;not null"`
	Name               string                      `gorm:"column:name;type:text;not null"`
	Cpf                string                      `gorm:"column:cpf;type:text;not null"`
	LegalName          *string                     `gorm:"column:legal_name;type:text;not null"`
	DefaultProbability *int8                       `gorm:"column:default_probability;type:integer"`
	DefaultScore       *int8                       `gorm:"column:default_score;type:integer"`
	DefaultDebits      *int8                       `gorm:"column:default_debits;type:integer"`
	DefaultProtests    *int8                       `gorm:"column:default_protests;type:integer"`
	IsActive           bool                        `gorm:"column:is_active;type:boolean;not null"`
	CreatedAt          time.Time                   `gorm:"column:created_at;type:timestamptz(6);not null"`
	AcceptedAt         *time.Time                  `gorm:"column:accepted_at;type:timestamptz(6)"`
	RejectedAt         *time.Time                  `gorm:"column:rejected_at;type:timestamptz(6)"`
	FailedAt           *time.Time                  `gorm:"column:failed_at;type:timestamptz(6)"`
	UpdatedAt          time.Time                   `gorm:"column:updated_at;type:timestamptz(6);not null"`
	UserID             uuid.UUID                   `gorm:"column:user_id;type:uuid;foreignKey:User;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User               UserModel                   `gorm:"foreignKey:UserID"`
}

func MigratePersonModel(db *gorm.DB) error {
	err := db.Table("person").AutoMigrate(&PasswordModel{})
	if err != nil {
		return err
	}
	return nil
}

func (PersonModel) TableName() string {
	return "person"
}
