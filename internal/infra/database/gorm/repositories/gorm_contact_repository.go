package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"
	"github.com/intwone/eda-arch-golang/internal/infra/database/gorm/models"
	uuid "github.com/satori/go.uuid"
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
	var contact models.ContactModel
	err := r.DB.Where(&models.ContactModel{Value: value, IsActive: true}).Take(&contact).Error
	if err != nil {
		return nil, err
	}
	return &entities.ContactEntity{
		ID:         uuid.UUID(contact.ID),
		Status:     contact.Status,
		Kind:       contact.Kind,
		Value:      contact.Value,
		IsActive:   contact.IsActive,
		CreatedAt:  contact.CreatedAt,
		VeriedAt:   contact.VeriedAt,
		AcceptedAt: contact.AcceptedAt,
		UpdatedAt:  contact.UpdatedAt,
		UserID:     uuid.UUID(contact.UserID),
	}, nil
}
