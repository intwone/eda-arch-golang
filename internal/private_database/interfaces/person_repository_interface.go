package interfaces

import (
	personEntities "github.com/intwone/eda-arch-golang/internal/public_person/domain/entities"
	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"

	uuid "github.com/satori/go.uuid"
)

type PersonRepositoryInterface interface {
	Create(person personEntities.PersonEntity) (*personEntities.PersonEntity, error)
	UpdateManyIsActiveByUserID(userID uuid.UUID) error
	FindFirstActiveByCpf(cpf personValueObject.Cpf) (*personEntities.PersonEntity, error)
}
