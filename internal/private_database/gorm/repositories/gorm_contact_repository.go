package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	contactEntities "github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
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

func (r *GORMContactRepository) Update(contact contactEntities.ContactEntity) (*contactEntities.ContactEntity, error) {
	c := models.ContactModel{
		ID:         contact.ID,
		Status:     contact.Status,
		Kind:       contact.Kind,
		Value:      contact.Value,
		IsActive:   contact.IsActive,
		CreatedAt:  contact.CreatedAt,
		VerifiedAt: contact.VerifiedAt,
		AcceptedAt: contact.AcceptedAt,
		UpdatedAt:  contact.UpdatedAt,
		UserID:     contact.UserID,
	}
	err := r.DB.Save(c).Error
	if err != nil {
		return nil, err
	}
	return &contactEntities.ContactEntity{
		ID:         uuid.UUID(contact.ID),
		Status:     contact.Status,
		Kind:       contact.Kind,
		Value:      contact.Value,
		IsActive:   contact.IsActive,
		CreatedAt:  contact.CreatedAt,
		VerifiedAt: contact.VerifiedAt,
		AcceptedAt: contact.AcceptedAt,
		UpdatedAt:  contact.UpdatedAt,
		UserID:     uuid.UUID(contact.UserID),
	}, nil
}

func (r *GORMContactRepository) FindFirstActiveByValue(value string) (*contactEntities.ContactEntity, error) {
	var contact models.ContactModel
	err := r.DB.Where(&models.ContactModel{IsActive: true, Value: value}).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contactEntities.ContactEntity{
		ID:         uuid.UUID(contact.ID),
		Status:     contact.Status,
		Kind:       contact.Kind,
		Value:      contact.Value,
		IsActive:   contact.IsActive,
		CreatedAt:  contact.CreatedAt,
		VerifiedAt: contact.VerifiedAt,
		AcceptedAt: contact.AcceptedAt,
		UpdatedAt:  contact.UpdatedAt,
		UserID:     uuid.UUID(contact.UserID),
	}, nil
}
