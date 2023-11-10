package models

import (
	"time"

	companyEntities "github.com/intwone/eda-arch-golang/internal/public_company/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CompanyModel struct {
	ID        uuid.UUID                     `gorm:"column:company_id;type:uuid;primaryKey;unique;not null"`
	Name      string                        `gorm:"column:name;type:text;not null"`
	Status    companyEntities.CompanyStatus `gorm:"column:status;type:text;not null"`
	CreatedAt time.Time                     `gorm:"column:created_at;type:timestamptz(6);not null"`
	UpdatedAt time.Time                     `gorm:"column:updated_at;type:timestamptz(6);not null"`
}

func MigrateCompany(db *gorm.DB) error {
	err := db.Table("company").AutoMigrate(&CompanyModel{})
	if err != nil {
		return err
	}
	return nil
}

func (CompanyModel) TableName() string {
	return "company"
}
