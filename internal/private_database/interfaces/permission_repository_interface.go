package interfaces

import (
	permissionEntities "github.com/intwone/eda-arch-golang/internal/public_permission/domain/entities"
)

type PermissionRepositoryInterface interface {
	Create(permission permissionEntities.PermissionEntity) (*permissionEntities.PermissionEntity, error)
}
