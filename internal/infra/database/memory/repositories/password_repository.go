package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/domain/auth/entities"
	uuid "github.com/satori/go.uuid"
)

type InMemoryPasswordRepository struct {
	password map[int]entities.PasswordEntity
}

func NewInMemoryPasswordRepository() *InMemoryPasswordRepository {
	return &InMemoryPasswordRepository{
		password: map[int]entities.PasswordEntity{},
	}
}

func (r *InMemoryPasswordRepository) Save(password entities.PasswordEntity) (*entities.PasswordEntity, error) {
	pass := entities.NewPasswordEntity(entities.PasswordEmail, "aaa", uuid.NewV4())
	return pass, nil
}
