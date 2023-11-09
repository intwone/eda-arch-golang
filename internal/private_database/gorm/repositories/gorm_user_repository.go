package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	"github.com/intwone/eda-arch-golang/internal/private_database/mappers"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GORMUserRepository struct {
	DB *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{
		DB: db,
	}
}

func (r *GORMUserRepository) Create(user userEntities.UserEntity) (*userEntities.UserEntity, error) {
	userModel := mappers.UserMapperDomainToGORM(user)
	if err := r.DB.Create(userModel).Error; err != nil {
		return nil, err
	}
	userDomain := mappers.UserMapperGORMToDomain(userModel)
	return &userDomain, nil
}

func (r *GORMUserRepository) Update(user userEntities.UserEntity) (*userEntities.UserEntity, error) {
	userModel := mappers.UserMapperDomainToGORM(user)
	if err := r.DB.Save(userModel).Error; err != nil {
		return nil, err
	}
	userDomain := mappers.UserMapperGORMToDomain(userModel)
	return &userDomain, nil
}

func (r *GORMUserRepository) FindById(id uuid.UUID) (*userEntities.UserEntity, error) {
	var userModel models.UserModel
	if err := r.DB.Where(&models.UserModel{ID: id}).Take(&userModel).Error; err != nil {
		return nil, err
	}
	userDomain := mappers.UserMapperGORMToDomain(userModel)
	return &userDomain, nil
}
