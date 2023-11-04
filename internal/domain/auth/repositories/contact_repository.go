package repositories

import "github.com/intwone/eda-arch-golang/internal/domain/auth/entities"

type ContactRepositoryInterface interface {
	FindFirstActiveByValue(value string) (*entities.ContactEntity, error)
}
