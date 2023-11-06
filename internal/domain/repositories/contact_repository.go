package repositories

import "github.com/intwone/eda-arch-golang/internal/domain/modules/auth/entities"

type ContactRepositoryInterface interface {
	FindFirstActiveByValue(value string) (*entities.ContactEntity, error)
}
