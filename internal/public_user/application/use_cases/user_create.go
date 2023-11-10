package usecases

import (
	"time"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
	contactValueObject "github.com/intwone/eda-arch-golang/internal/public_contact/domain/value_objects"
	personValueObject "github.com/intwone/eda-arch-golang/internal/public_person/domain/value_objects"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
	uuid "github.com/satori/go.uuid"
)

type UserCreateInput struct {
	Name      string
	Email     contactValueObject.Email
	Cpf       personValueObject.Cpf
	Birthdate string
}

type UserCreateOutput struct {
	User userEntities.UserEntity
}

type UserCreateUseCaseInterface interface {
	Execute(input UserCreateInput) (*UserCreateOutput, error)
}

type UserCreateUseCase struct {
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
) *UserCreateUseCase {
	uc := UserCreateUseCase{
		EventDispatcher:   ed,
		UserRepository:    ur,
		ContactRepository: cr,
		PersonRepository:  pr,
	}
	return &uc
}

func (uc *UserCreateUseCase) Execute(input UserCreateInput) (*UserCreateOutput, error) {
	contact, _ := uc.ContactRepository.FindFirstActiveByValue(input.Email)
	if contact != nil {
		return nil, err.NewResourceAlreadyTakenError("contact")
	}
	person, _ := uc.PersonRepository.FindFirstActiveByCpf(input.Cpf)
	if person != nil {
		return nil, err.NewResourceAlreadyTakenError("person")
	}
	user := userEntities.NewUserEntity()
	_, urErr := uc.UserRepository.Create(*user)
	if urErr != nil {
		return nil, urErr
	}
	birthdate, pErr := time.Parse("2006/01/02", input.Birthdate)
	if pErr != nil {
		return nil, pErr
	}
	companyID, _ := uuid.FromString("df708e7a-f455-43c4-89b0-d5c50e6f04be") // my-company
	payload := domainEvents.NewUserCreatedEvent(input.Email, input.Cpf, birthdate, input.Name, user.GetID(), companyID)
	event := events.NewEvent(domainEvents.UserCreatedEventName, payload)
	uc.EventDispatcher.Dispatch(*event)
	return &UserCreateOutput{User: *user}, nil
}
