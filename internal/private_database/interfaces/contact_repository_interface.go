package interfaces

import (
	contactEntities "github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	uuid "github.com/satori/go.uuid"
)

type ContactRepositoryInterface interface {
	Create(contact contactEntities.ContactEntity) (*contactEntities.ContactEntity, error)
	Update(contact contactEntities.ContactEntity) (*contactEntities.ContactEntity, error)
	UpdateManyIsActiveByUserID(userID uuid.UUID) error
	FindFirstActiveByValue(value contactValueObject.Email) (*contactEntities.ContactEntity, error)
}
