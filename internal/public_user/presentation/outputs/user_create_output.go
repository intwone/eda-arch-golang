package outputs

import (
	"github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	uuid "github.com/satori/go.uuid"
)

type CreateUserOutput struct {
	ID     uuid.UUID           `json:"id"`
	Status entities.UserStatus `json:"status"`
}

func CreateUserOutputMapper(user entities.UserEntity) CreateUserOutput {
	return CreateUserOutput{
		ID:     user.GetID(),
		Status: user.GetStatus(),
	}
}
