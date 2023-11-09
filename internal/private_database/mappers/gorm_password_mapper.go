package mappers

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	passwordEntities "github.com/intwone/eda-arch-golang/internal/public_password/domain/entities"
)

func PasswordMapperDomainToGORM(password passwordEntities.PasswordEntity) models.PasswordModel {
	return models.PasswordModel{
		ID:        password.GetID(),
		Kind:      password.GetKind(),
		Hash:      password.GetHash(),
		IsActive:  password.GetIsActive(),
		CreatedAt: password.GetCreatedAt(),
		ContactID: password.GetContactID(),
	}
}

func PasswordMapperGORMToDomain(password models.PasswordModel) passwordEntities.PasswordEntity {
	return passwordEntities.PasswordEntity{
		ID:        password.ID,
		Kind:      password.Kind,
		Hash:      password.Hash,
		IsActive:  password.IsActive,
		CreatedAt: password.CreatedAt,
		ContactID: password.ContactID,
	}
}
