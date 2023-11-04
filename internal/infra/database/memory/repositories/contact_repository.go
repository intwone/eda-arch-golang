package repositories

import (
	"errors"

	"github.com/intwone/eda-arch-golang/internal/domain/auth/entities"
	uuid "github.com/satori/go.uuid"
)

type InMemoryContactRepository struct {
	contacts []entities.ContactEntity
}

func NewInMemoryContactRepository() *InMemoryContactRepository {
	return &InMemoryContactRepository{
		contacts: []entities.ContactEntity{},
	}
}

func (r *InMemoryContactRepository) FindFirstActiveByValue(value string) (*entities.ContactEntity, error) {
	contact := entities.NewContactEntity("cassio@gmail.com", "email", uuid.NewV4())
	r.contacts = append(r.contacts, *contact)

	for _, contact := range r.contacts {
		if contact.Value == value {
			return &contact, nil
		}
	}

	return nil, errors.New("not found")
}
