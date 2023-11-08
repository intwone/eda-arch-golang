package usecases

import (
	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	hasher "github.com/intwone/eda-arch-golang/internal/private_hasher/interfaces"
	passwordEntities "github.com/intwone/eda-arch-golang/internal/public_password/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_password/events"
	"github.com/intwone/eda-arch-golang/internal/utils"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type AuthCreateInput struct {
	Email string
}

type AuthCreateUseCaseInterface interface {
	Execute(input AuthCreateInput) error
}

type AuthCreateUseCase struct {
	EventDispatcher    events.EventDispatcherInterface
	ContactRepository  repositories.ContactRepositoryInterface
	PasswordRepository repositories.PasswordRepositoryInterface
	Hasher             hasher.HasherInterface
}

func NewAuthCreateUseCase(ed events.EventDispatcherInterface, cr repositories.ContactRepositoryInterface, pr repositories.PasswordRepositoryInterface, h hasher.HasherInterface) *AuthCreateUseCase {
	uc := AuthCreateUseCase{
		EventDispatcher:    ed,
		ContactRepository:  cr,
		PasswordRepository: pr,
		Hasher:             h,
	}
	return &uc
}

func (uc *AuthCreateUseCase) Execute(input AuthCreateInput) error {
	contact, crErr := uc.ContactRepository.FindFirstActiveByValue(input.Email)
	if crErr != nil {
		return crErr
	}
	unhash := utils.GenerateCodeUtil()
	hash, hErr := uc.Hasher.Hash(unhash)
	if hErr != nil {
		return hErr
	}
	password := passwordEntities.NewPasswordEntity(passwordEntities.PasswordKind(contact.Kind), *hash, contact.GetID())
	_, prErr := uc.PasswordRepository.Create(*password)
	if prErr != nil {
		return prErr
	}
	payload := domainEvents.NewPasswordCreatedEvent(*password, unhash)
	event := events.NewEvent(domainEvents.PasswordCreatedEventName, payload)
	uc.EventDispatcher.Dispatch(*event)
	return nil
}
