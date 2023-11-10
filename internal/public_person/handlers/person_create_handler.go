package handlers

import (
	"sync"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	personEntity "github.com/intwone/eda-arch-golang/internal/public_person/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type PersonCreateHandler struct {
	PersonRepository repositories.PersonRepositoryInterface
}

func NewPersonCreateHandler(pr repositories.PersonRepositoryInterface) *PersonCreateHandler {
	return &PersonCreateHandler{
		PersonRepository: pr,
	}
}

func (h *PersonCreateHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.UserCreatedEventName {
		payload := event.GetPayload().(*domainEvents.UserCreatedEvent)
		name := payload.GetPersonName()
		cpf := payload.GetCpf()
		birthdate := payload.GetBirthdate()
		userID := payload.GetUserID()
		person, _ := personEntity.NewPersonEntity(name, cpf.Value, birthdate, userID)
		h.PersonRepository.Upsert(*person)
	}
	wg.Done()
}
