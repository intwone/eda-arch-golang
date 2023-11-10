package interfaces

import (
	accountEntities "github.com/intwone/eda-arch-golang/internal/public_account/domain/entities"
)

type AccountRepositoryInterface interface {
	Create(account accountEntities.AccountEntity) (*accountEntities.AccountEntity, error)
}
