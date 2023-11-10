package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	"github.com/intwone/eda-arch-golang/internal/private_database/mappers"
	passwordEntities "github.com/intwone/eda-arch-golang/internal/public_password/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GORMPasswordRepository struct {
	DB *gorm.DB
}

func NewGORMPasswordRepository(db *gorm.DB) *GORMPasswordRepository {
	return &GORMPasswordRepository{
		DB: db,
	}
}

func (r *GORMPasswordRepository) Upsert(password passwordEntities.PasswordEntity) (*passwordEntities.PasswordEntity, error) {
	passwordModel := mappers.PasswordMapperDomainToGORM(password)
	if err := r.UpdateManyIsActiveByContactID(password.GetContactID()); err != nil {
		return nil, err
	}
	if err := r.DB.Create(passwordModel).Error; err != nil {
		return nil, err
	}
	return &password, nil
}

func (r *GORMPasswordRepository) UpdateManyIsActiveByContactID(contactID uuid.UUID) error {
	return r.DB.Model(models.PasswordModel{}).Where("is_active = ? and contact_id = ?", true, contactID).Update("is_active", false).Error
}

func (r *GORMPasswordRepository) FindFirstActiveByContactID(contactID uuid.UUID) (*passwordEntities.PasswordEntity, error) {
	var passwordModel models.PasswordModel
	if err := r.DB.Where(&models.PasswordModel{ContactID: contactID, IsActive: true}).Take(&passwordModel).Error; err != nil {
		return nil, err
	}
	passwordDomain := mappers.PasswordMapperGORMToDomain(passwordModel)
	return &passwordDomain, nil
}
