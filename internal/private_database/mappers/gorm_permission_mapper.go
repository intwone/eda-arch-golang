package mappers

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	permissionEntities "github.com/intwone/eda-arch-golang/internal/public_permission/domain/entities"
)

func PermissionMapperDomainToGORM(permission permissionEntities.PermissionEntity) models.PermissionModel {
	return models.PermissionModel{
		ID:         permission.GetID(),
		Status:     permission.GetStatus(),
		Role:       permission.GetRole(),
		CreatedAt:  permission.GetCreatedAt(),
		ArchivedAt: permission.GetArchivedAt(),
		UpdatedAt:  permission.GetUpdatedAt(),
		UserID:     permission.GetUserID(),
		CompanyID:  permission.GetCompanyID(),
		AccountID:  permission.GetAccountID(),
	}
}

func PermissionMapperGORMToDomain(permission models.PermissionModel) *permissionEntities.PermissionEntity {
	return &permissionEntities.PermissionEntity{
		ID:         permission.ID,
		Status:     permission.Status,
		Role:       permission.Role,
		CreatedAt:  permission.CreatedAt,
		ArchivedAt: permission.ArchivedAt,
		UpdatedAt:  permission.UpdatedAt,
		UserID:     permission.UserID,
		CompanyID:  permission.CompanyID,
		AccountID:  permission.AccountID,
	}
}
