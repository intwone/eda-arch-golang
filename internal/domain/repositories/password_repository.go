package repositories

import "github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"

type PasswordRepositoryInterface interface {
	Save(password entities.PasswordEntity) (*entities.PasswordEntity, error)
}