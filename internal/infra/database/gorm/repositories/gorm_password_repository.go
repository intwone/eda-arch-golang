package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"
	"github.com/intwone/eda-arch-golang/internal/infra/database/gorm/models"
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

func (r *GORMPasswordRepository) Save(password entities.PasswordEntity) (*entities.PasswordEntity, error) {
	passwordModel := models.PasswordModel{
		ID:        uuid.UUID(password.ID),
		Kind:      password.Kind,
		Hash:      password.Hash,
		IsActive:  password.IsActive,
		CreatedAt: password.CreatedAt,
		ContactID: password.ContactID,
	}
	err := r.DB.Create(passwordModel).Error
	if err != nil {
		return nil, err
	}
	return &password, nil
}
