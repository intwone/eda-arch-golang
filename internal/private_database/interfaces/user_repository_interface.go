package interfaces

import (
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	uuid "github.com/satori/go.uuid"
)

type UserRepositoryInterface interface {
	Update(user userEntities.UserEntity) (*userEntities.UserEntity, error)
	FindById(id uuid.UUID) (*userEntities.UserEntity, error)
}
