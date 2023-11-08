package interfaces

import (
	contactEntities "github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
)

type ContactRepositoryInterface interface {
	Update(contact contactEntities.ContactEntity) (*contactEntities.ContactEntity, error)
	FindFirstActiveByValue(value string) (*contactEntities.ContactEntity, error)
}
