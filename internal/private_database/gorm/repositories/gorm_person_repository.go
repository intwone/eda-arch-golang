package repositories

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	"github.com/intwone/eda-arch-golang/internal/private_database/mappers"
	personEntities "github.com/intwone/eda-arch-golang/internal/public_person/domain/entities"
	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GORMPersonRepository struct {
	DB *gorm.DB
}

func NewGORMPersonRepository(db *gorm.DB) *GORMPersonRepository {
	return &GORMPersonRepository{
		DB: db,
	}
}

func (r *GORMPersonRepository) Create(person personEntities.PersonEntity) (*personEntities.PersonEntity, error) {
	personModel := mappers.PersonMapperDomainToGORM(person)
	if err := r.UpdateManyIsActiveByUsertID(person.GetUserID()); err != nil {
		return nil, err
	}
	if err := r.DB.Create(personModel).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *GORMPersonRepository) UpdateManyIsActiveByUsertID(userID uuid.UUID) error {
	return r.DB.Model(models.PersonModel{}).Where("is_active = ? and user_id = ?", true, userID).Update("is_active", false).Error
}

func (r *GORMPersonRepository) FindFirstActiveByCpf(cpf personValueObject.Cpf) (*personEntities.PersonEntity, error) {
	var personModel models.PersonModel
	if err := r.DB.Where(&models.PersonModel{Cpf: cpf.Value, IsActive: true}).Error; err != nil {
		return nil, err
	}
	personDomain := mappers.PersonMapperGORMToDomain(personModel)
	return personDomain, nil
}
