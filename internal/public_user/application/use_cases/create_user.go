package usecases

import (
	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type CreateUserInput struct {
	Name  string
	Email contactValueObject.Email
	Cpf   personValueObject.Cpf
}

type CreateUserOutput struct {
	User userEntities.UserEntity
}

type CreateUserUseCaseInterface interface {
	Execute(input CreateUserInput) (*CreateUserOutput, error)
}

type CreateUserUseCase struct {
	EventDispatcher   events.EventDispatcherInterface
	UserRepository    repositories.UserRepositoryInterface
	ContactRepository repositories.ContactRepositoryInterface
	PersonRepository  repositories.PersonRepositoryInterface
}

func NewCreateUserUseCase(
	ed events.EventDispatcherInterface,
	ur repositories.UserRepositoryInterface,
	cr repositories.ContactRepositoryInterface,
	pr repositories.PersonRepositoryInterface,
) *CreateUserUseCase {
	uc := CreateUserUseCase{
		EventDispatcher:   ed,
		UserRepository:    ur,
		ContactRepository: cr,
		PersonRepository:  pr,
	}
	return &uc
}

func (uc *CreateUserUseCase) Execute(input CreateUserInput) (*CreateUserOutput, error) {
	contact, crErr := uc.ContactRepository.FindFirstActiveByValue(input.Email)
	if crErr != nil {
		return nil, crErr
	}
	if contact != nil {
		return nil, err.NewResourceAlreadyTakenError("contact")
	}
	person, prErr := uc.PersonRepository.FindFirstActiveByCpf(input.Cpf)
	if prErr != nil {
		return nil, prErr
	}
	if person != nil {
		return nil, err.NewResourceAlreadyTakenError("person")
	}
	user := userEntities.NewUserEntity()
	_, urErr := uc.UserRepository.Create(*user)
	if urErr != nil {
		return nil, urErr
	}
	payload := domainEvents.NewUserCreatedEvent(user.GetID(), input.Email, input.Cpf)
	event := events.NewEvent(domainEvents.UserCreatedEventName, payload)
	uc.EventDispatcher.Dispatch(*event)
	return &CreateUserOutput{User: *user}, nil
}
