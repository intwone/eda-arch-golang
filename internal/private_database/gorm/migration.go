package gorm

import (
	"log"

	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

func Migrate(db *gorm.DB) {
	models := []interface{}{
		&models.PasswordModel{},
		&models.UserModel{},
		&models.PersonModel{},
		&models.ContactModel{},
		&models.PermissionModel{},
		&models.AccountModel{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatal(err)
		}
	}
}
