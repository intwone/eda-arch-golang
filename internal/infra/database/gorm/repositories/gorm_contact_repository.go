package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/domain/auth/entities"
	"github.com/intwone/eda-arch-golang/internal/infra/database/gorm/models"
	"gorm.io/gorm"
)

type GORMContactRepository struct {
	DB *gorm.DB
}

func NewGORMContactRepository(db *gorm.DB) *GORMContactRepository {
	return &GORMContactRepository{
		DB: db,
	}
}

func (r *GORMContactRepository) FindFirstActiveByValue(value string) (*entities.ContactEntity, error) {
	var contact entities.ContactEntity
	err := r.DB.Where(&models.ContactModel{Value: value, IsActive: true}).Take(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}
