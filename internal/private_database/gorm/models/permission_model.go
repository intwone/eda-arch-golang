package models

import (
	"time"

	permissionEntities "github.com/intwone/eda-arch-golang/internal/public_permission/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PermissionModel struct {
	ID         uuid.UUID                           `gorm:"column:permission_id;type:uuid;primaryKey;unique;not null"`
	Status     permissionEntities.PermissionStatus `gorm:"column:status;type:text;not null"`
	Role       permissionEntities.PermissionRole   `gorm:"column:role;type:text;not null"`
	CreatedAt  time.Time                           `gorm:"column:created_at;type:timestamptz(6);not null"`
	ArchivedAt *time.Time                          `gorm:"column:archived_at;type:timestamptz(6)"`
	UpdatedAt  time.Time                           `gorm:"column:updated_at;type:timestamptz(6);not null"`
	UserID     uuid.UUID                           `gorm:"column:user_id;type:uuid;foreignKey:User;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CompanyID  uuid.UUID                           `gorm:"column:company_id;type:uuid;foreignKey:Company;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccountID  uuid.UUID                           `gorm:"column:account_id;type:uuid;foreignKey:Account;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User       UserModel                           `gorm:"foreignKey:UserID"`
	Company    CompanyModel                        `gorm:"foreignKey:CompanyID"`
	Account    AccountModel                        `gorm:"foreignKey:AccountID"`
}

func MigratePermission(db *gorm.DB) error {
	err := db.Table("permission").AutoMigrate(&PermissionModel{})
	if err != nil {
		return err
	}
	return nil
}

func (PermissionModel) TableName() string {
	return "permission"
}
