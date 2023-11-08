package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GORMUserRepository struct {
	DB *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{
		DB: db,
	}
}

func (r *GORMUserRepository) Update(user userEntities.UserEntity) (*userEntities.UserEntity, error) {
	u := models.UserModel{
		ID:           uuid.UUID(user.ID),
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
	err := r.DB.Save(u).Error
	if err != nil {
		return nil, err
	}
	return &userEntities.UserEntity{
		ID:           uuid.UUID(user.ID),
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
	}, nil
}

func (r *GORMUserRepository) FindById(id uuid.UUID) (*userEntities.UserEntity, error) {
	var user models.UserModel
	err := r.DB.Where(&models.UserModel{ID: id}).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &userEntities.UserEntity{
		ID:           uuid.UUID(user.ID),
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
	}, nil
}
