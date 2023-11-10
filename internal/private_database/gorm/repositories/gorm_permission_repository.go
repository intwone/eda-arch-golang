package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/mappers"
	permissionEntities "github.com/intwone/eda-arch-golang/internal/public_permission/domain/entities"
	"gorm.io/gorm"
)

type GORMPermissionRepository struct {
	DB *gorm.DB
}

func NewGORMPermissionRepository(db *gorm.DB) *GORMPermissionRepository {
	return &GORMPermissionRepository{
		DB: db,
	}
}

func (r *GORMPermissionRepository) Create(permission permissionEntities.PermissionEntity) (*permissionEntities.PermissionEntity, error) {
	permissionModel := mappers.PermissionMapperDomainToGORM(permission)
	if err := r.DB.Create(permissionModel).Error; err != nil {
		return nil, err
	}
	permissionDomain := mappers.PermissionMapperGORMToDomain(permissionModel)
	return permissionDomain, nil
}
