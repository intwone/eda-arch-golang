package mappers

import (
	"github.com/intwone/eda-arch-golang/internal/private_database/gorm/models"
	personEntities "github.com/intwone/eda-arch-golang/internal/public_person/domain/entities"
	personValueObjects "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
)

func PersonMapperDomainToGORM(person personEntities.PersonEntity) models.PersonModel {
	return models.PersonModel{
		ID:                 person.GetID(),
		Status:             person.GetStatus(),
		Name:               person.GetName(),
		Cpf:                person.GetCpf().Value,
		LegalName:          person.GetLegalName(),
		DefaultProbability: person.GetDefaultProbability(),
		DefaultScore:       person.GetDefaultScore(),
		DefaultDebits:      person.GetDefaultDebits(),
		DefaultProtests:    person.GetDefaultProtests(),
		IsActive:           person.GetIsActive(),
		CreatedAt:          person.GetCreatedAt(),
		AcceptedAt:         person.GetAcceptedAt(),
		RejectedAt:         person.GetRejectedAt(),
		FailedAt:           person.GetFailedAt(),
		UpdatedAt:          person.GetUpdatedAt(),
		UserID:             person.GetUserID(),
	}
}

func PersonMapperGORMToDomain(person models.PersonModel) *personEntities.PersonEntity {
	cpf, err := personValueObjects.NewCpf(person.Cpf)
	if err != nil {
		return nil
	}
	return &personEntities.PersonEntity{
		ID:                 person.ID,
		Status:             person.Status,
		Name:               person.Name,
		Cpf:                *cpf,
		LegalName:          person.LegalName,
		DefaultProbability: person.DefaultProbability,
		DefaultScore:       person.DefaultScore,
		DefaultDebits:      person.DefaultDebits,
		DefaultProtests:    person.DefaultProtests,
		IsActive:           person.IsActive,
		CreatedAt:          person.CreatedAt,
		AcceptedAt:         person.AcceptedAt,
		RejectedAt:         person.RejectedAt,
		FailedAt:           person.FailedAt,
		UpdatedAt:          person.UpdatedAt,
		UserID:             person.UserID,
	}
}
