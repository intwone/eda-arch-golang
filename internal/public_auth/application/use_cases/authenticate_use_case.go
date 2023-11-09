package usecases

import (
	cryptography "github.com/intwone/eda-arch-golang/internal/private_cryptography/interfaces"
	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	hasher "github.com/intwone/eda-arch-golang/internal/private_hasher/interfaces"
	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_auth/events"
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type AuthenticateInput struct {
	Email    contactValueObject.Email
	Password string
}

type AuthenticateOutput struct {
	Token string
}

type AuthenticateUseCaseInterface interface {
	Execute(input AuthenticateInput) (*AuthenticateOutput, error)
}

type AuthenticateUseCase struct {
	EventDispatcher    events.EventDispatcherInterface
	ContactRepository  repositories.ContactRepositoryInterface
	PasswordRepository repositories.PasswordRepositoryInterface
	UserRepository     repositories.UserRepositoryInterface
	Cryptography       cryptography.CryptographyInterface
	Hasher             hasher.HasherInterface
}

func NewAuthenticateUseCase(
	ed events.EventDispatcherInterface,
	cr repositories.ContactRepositoryInterface,
	pr repositories.PasswordRepositoryInterface,
	ur repositories.UserRepositoryInterface,
	crypt cryptography.CryptographyInterface,
	hasher hasher.HasherInterface,
) *AuthenticateUseCase {
	uc := AuthenticateUseCase{
		EventDispatcher:    ed,
		ContactRepository:  cr,
		PasswordRepository: pr,
		UserRepository:     ur,
		Cryptography:       crypt,
		Hasher:             hasher,
	}
	return &uc
}

func (uc *AuthenticateUseCase) Execute(input AuthenticateInput) (*AuthenticateOutput, error) {
	contact, crErr := uc.ContactRepository.FindFirstActiveByValue(input.Email)
	if crErr != nil {
		return nil, crErr
	}
	user, urErr := uc.UserRepository.FindById(contact.UserID)
	if urErr != nil {
		return nil, urErr
	}
	password, prErr := uc.PasswordRepository.FindFirstActiveByContactID(contact.GetID())
	if prErr != nil {
		return nil, prErr
	}
	match := uc.Hasher.Compare(input.Password, password.GetHash())
	if !match {
		return nil, err.UnauthorizedError
	}
	token, cryptErr := uc.Cryptography.Encrypt(user.GetID().String())
	if cryptErr != nil {
		return nil, cryptErr
	}
	payload := domainEvents.NewAuthenticatedEvent(*user, *contact)
	event := events.NewEvent(domainEvents.AuthenticatedEventName, payload)
	uc.EventDispatcher.Dispatch(*event)
	return &AuthenticateOutput{Token: *token}, nil
}
