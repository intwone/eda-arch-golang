package handlers

import (
	"sync"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	"github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type ContactCreateHandler struct {
	ContactRepository repositories.ContactRepositoryInterface
}

func NewContactCreateHandler(cr repositories.ContactRepositoryInterface) *ContactCreateHandler {
	return &ContactCreateHandler{
		ContactRepository: cr,
	}
}

func (h *ContactCreateHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.UserCreatedEventName {
		payload := event.GetPayload().(*domainEvents.UserCreatedEvent)
		email := payload.GetEmail()
		userID := payload.GetUserID()
		contact := entities.NewContactEntity(email, entities.ContactEmail, userID)
		h.ContactRepository.Upsert(*contact)
	}
	wg.Done()
}
