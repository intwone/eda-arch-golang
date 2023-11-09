package mappers

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
)

func UserMapperDomainToGORM(user userEntities.UserEntity) models.UserModel {
	return models.UserModel{
		ID:           user.GetID(),
		Status:       user.GetStatus(),
		CreatedAt:    user.GetCreatedAt(),
		VerifiedAt:   user.GetVerifiedAt(),
		PendingAt:    user.GetPendingAt(),
		WaitlistedAt: user.GetWaitlistedAt(),
		SuspendedAt:  user.GetSuspendedAt(),
		MemberAt:     user.GetMemberAt(),
		ChurnedAt:    user.GetChurnedAt(),
		ArchivedAt:   user.GetArchivedAt(),
		SettledAt:    user.GetSettledAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
}

func UserMapperGORMToDomain(user models.UserModel) userEntities.UserEntity {
	return userEntities.UserEntity{
		ID:           user.ID,
		Status:       user.Status,
		CreatedAt:    user.CreatedAt,
		VerifiedAt:   user.VerifiedAt,
		PendingAt:    user.PendingAt,
		WaitlistedAt: user.WaitlistedAt,
		SuspendedAt:  user.SuspendedAt,
		MemberAt:     user.MemberAt,
		ChurnedAt:    user.ChurnedAt,
		ArchivedAt:   user.ArchivedAt,
		SettledAt:    user.SettledAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
