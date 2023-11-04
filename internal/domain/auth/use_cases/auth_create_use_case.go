package usecases

import (
	"github.com/intwone/eda-arch-golang/internal/domain/auth/entities"
	"github.com/intwone/eda-arch-golang/internal/domain/auth/repositories"
	domainEvents "github.com/intwone/eda-arch-golang/internal/domain/password/events"
	"github.com/intwone/eda-arch-golang/internal/infra/hasher"
	"github.com/intwone/eda-arch-golang/internal/utils"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type AuthCreateInput struct {
	Email string
}

type AuthCreateUseCaseInterface interface {
	Execute(input AuthCreateInput) bool
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

func (uc *AuthCreateUseCase) Execute(input AuthCreateInput) bool {
	contact, crErr := uc.ContactRepository.FindFirstActiveByValue(input.Email)
	if crErr != nil {
		return false
	}
	unhash := utils.GenerateCodeUtil()
	hash, hErr := uc.Hasher.Hash(unhash)
	if hErr != nil {
		return false
	}
	password := entities.NewPasswordEntity(entities.PasswordKind(contact.Kind), *hash, contact.ID)
	_, prErr := uc.PasswordRepository.Save(*password)
	if prErr != nil {
		return false
	}
	payload := domainEvents.NewPasswordCreatedEvent(*password, unhash)
	event := events.NewEvent(domainEvents.PasswordCreatedEventName, payload)
	uc.EventDispatcher.Dispatch(*event)
	return true
}
