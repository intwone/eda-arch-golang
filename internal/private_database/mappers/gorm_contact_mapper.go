package mappers

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	contactEntities "github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
	contactValueObjects "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
)

func ContactMapperDomainToGORM(contact contactEntities.ContactEntity) models.ContactModel {
	return models.ContactModel{
		ID:         contact.GetID(),
		Status:     contact.GetStatus(),
		Kind:       contact.GetKind(),
		Value:      contact.GetValue().Value,
		IsActive:   contact.GetIsActive(),
		CreatedAt:  contact.GetCreatedAt(),
		VerifiedAt: contact.GetVerifiedAt(),
		AcceptedAt: contact.GetAcceptedAt(),
		UpdatedAt:  contact.GetUpdatedAt(),
		UserID:     contact.GetUserID(),
	}
}

func ContactMapperGORMToDomain(contact models.ContactModel) *contactEntities.ContactEntity {
	email, err := contactValueObjects.NewEmail(contact.Value)
	if err != nil {
		return nil
	}
	return &contactEntities.ContactEntity{
		ID:         contact.ID,
		Status:     contact.Status,
		Kind:       contact.Kind,
		Value:      *email,
		IsActive:   contact.IsActive,
		CreatedAt:  contact.CreatedAt,
		VerifiedAt: contact.VerifiedAt,
		AcceptedAt: contact.AcceptedAt,
		UpdatedAt:  contact.UpdatedAt,
		UserID:     contact.UserID,
	}
}
