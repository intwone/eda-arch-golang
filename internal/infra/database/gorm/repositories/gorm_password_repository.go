package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/domain/auth/entities"
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
	err := r.DB.Save(password).Error
	if err != nil {
		return nil, err
	}
	return &password, nil
}
