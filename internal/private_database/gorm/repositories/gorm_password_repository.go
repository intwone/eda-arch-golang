package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	passwordEntities "github.com/intwone/eda-arch-golang/internal/public_password/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GORMPasswordRepository struct {
	DB *gorm.DB
}

func NewGORMPasswordRepository(db *gorm.DB) *GORMPasswordRepository {
	return &GORMPasswordRepository{
		DB: db,
	}
}

func (r *GORMPasswordRepository) Create(password passwordEntities.PasswordEntity) (*passwordEntities.PasswordEntity, error) {
	passwordModel := models.PasswordModel{
		ID:        uuid.UUID(password.ID),
		Kind:      password.Kind,
		Hash:      password.Hash,
		IsActive:  password.IsActive,
		CreatedAt: password.CreatedAt,
		ContactID: password.ContactID,
	}
	if err := r.DB.Create(passwordModel).Error; err != nil {
		return nil, err
	}
	return &password, nil
}

func (r *GORMPasswordRepository) FindFirstActiveByContactID(contactID uuid.UUID) (*passwordEntities.PasswordEntity, error) {
	var password models.PasswordModel
	err := r.DB.Where(&models.PasswordModel{ContactID: contactID, IsActive: true}).Take(&password).Error
	if err != nil {
		return nil, err
	}
	return &passwordEntities.PasswordEntity{
		ID:        uuid.UUID(password.ID),
		Kind:      password.Kind,
		Hash:      password.Hash,
		IsActive:  password.IsActive,
		CreatedAt: password.CreatedAt,
		ContactID: password.ContactID,
	}, nil
}
