package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	"github.com/intwone/eda-arch-golang/internal/private_database/mappers"
	contactEntities "github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GORMContactRepository struct {
	DB *gorm.DB
}

func NewGORMContactRepository(db *gorm.DB) *GORMContactRepository {
	return &GORMContactRepository{
		DB: db,
	}
}

func (r *GORMContactRepository) Upsert(contact contactEntities.ContactEntity) (*contactEntities.ContactEntity, error) {
	contactModel := mappers.ContactMapperDomainToGORM(contact)
	if err := r.UpdateManyIsActiveByUserID(contact.GetUserID()); err != nil {
		return nil, err
	}
	if err := r.DB.Create(contactModel).Error; err != nil {
		return nil, err
	}
	contactDomain := mappers.ContactMapperGORMToDomain(contactModel)
	return contactDomain, nil
}

func (r *GORMContactRepository) Update(contact contactEntities.ContactEntity) (*contactEntities.ContactEntity, error) {
	contactModel := mappers.ContactMapperDomainToGORM(contact)
	if err := r.DB.Save(contactModel).Error; err != nil {
		return nil, err
	}
	contactDomain := mappers.ContactMapperGORMToDomain(contactModel)
	return contactDomain, nil
}

func (r *GORMContactRepository) UpdateManyIsActiveByUserID(userID uuid.UUID) error {
	return r.DB.Model(models.ContactModel{}).Where("is_active = ? and user_id = ?", true, userID).Update("is_active", false).Error
}

func (r *GORMContactRepository) FindFirstActiveByValue(value contactValueObject.Email) (*contactEntities.ContactEntity, error) {
	var contactModel models.ContactModel
	if err := r.DB.Where(&models.ContactModel{IsActive: true, Value: value.Value}).First(&contactModel).Error; err != nil {
		return nil, err
	}
	contactDomain := mappers.ContactMapperGORMToDomain(contactModel)
	return contactDomain, nil
}
