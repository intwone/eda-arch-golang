package mappers

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	accountEntities "github.com/intwone/eda-arch-golang/internal/public_account/domain/entities"
)

func AccountMapperDomainToGORM(account accountEntities.AccountEntity) models.AccountModel {
	return models.AccountModel{
		ID:        account.GetID(),
		Kind:      account.GetKind(),
		CreatedAt: account.GetCreatedAt(),
		UpdatedAt: account.GetUpdatedAt(),
		CohortAt:  account.GetCohortAt(),
	}
}

func AccountMapperGORMToDomain(account models.AccountModel) *accountEntities.AccountEntity {
	return &accountEntities.AccountEntity{
		ID:        account.ID,
		Kind:      account.Kind,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		CohortAt:  account.CohortAt,
	}
}
