package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/mappers"
	accountEntities "github.com/intwone/eda-arch-golang/internal/public_account/domain/entities"
	"gorm.io/gorm"
)

type GORMAccountRepository struct {
	DB *gorm.DB
}

func NewGORMAccountRepository(db *gorm.DB) *GORMAccountRepository {
	return &GORMAccountRepository{
		DB: db,
	}
}

func (r *GORMAccountRepository) Create(account accountEntities.AccountEntity) (*accountEntities.AccountEntity, error) {
	accountModel := mappers.AccountMapperDomainToGORM(account)
	if err := r.DB.Create(accountModel).Error; err != nil {
		return nil, err
	}
	accountDomain := mappers.AccountMapperGORMToDomain(accountModel)
	return accountDomain, nil
}
