package handlers

import (
	"sync"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type PersonCreateHandler struct {
	PersonRepository repositories.PersonRepositoryInterface
}

func NewContactVerifiedHandler(pr repositories.PersonRepositoryInterface) *PersonCreateHandler {
	return &PersonCreateHandler{
		PersonRepository: pr,
	}
}

func (h *PersonCreateHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.UserCreatedEventName {

	}
	wg.Done()
}
