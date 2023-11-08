package handlers

import (
	"sync"
	"time"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_auth/events"
	"github.com/intwone/eda-arch-golang/internal/public_contact/domain/entities"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type ContactStatusVerifiedHandler struct {
	ContactRepository repositories.ContactRepositoryInterface
}

func NewContactVerifiedHandler(cr repositories.ContactRepositoryInterface) *ContactStatusVerifiedHandler {
	return &ContactStatusVerifiedHandler{
		ContactRepository: cr,
	}
}

func (h *ContactStatusVerifiedHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.AuthenticatedEventName {
		payload := event.GetPayload().(*domainEvents.AuthenticatedEvent)
		contact := payload.GetContactID()
		now := time.Now()
		contact.SetStatus(entities.ContactVerified)
		contact.SetVerified(now)
		h.ContactRepository.Update(contact)
	}
	wg.Done()
}
