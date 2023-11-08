package interfaces

import (
	passwordEntities "github.com/intwone/eda-arch-golang/internal/public_password/domain/entities"
	uuid "github.com/satori/go.uuid"
)

type PasswordRepositoryInterface interface {
	Create(password passwordEntities.PasswordEntity) (*passwordEntities.PasswordEntity, error)
	FindFirstActiveByContactID(userID uuid.UUID) (*passwordEntities.PasswordEntity, error)
}
