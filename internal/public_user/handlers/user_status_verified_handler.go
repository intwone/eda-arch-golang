package handlers

import (
	"sync"
	"time"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_auth/events"
	userEntities "github.com/intwone/eda-arch-golang/internal/public_user/domain/entities"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type UserStatusVerifiedHandler struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserStatusVerifiedHandler(ur repositories.UserRepositoryInterface) *UserStatusVerifiedHandler {
	return &UserStatusVerifiedHandler{
		UserRepository: ur,
	}
}

func (h *UserStatusVerifiedHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.AuthenticatedEventName {
		payload := event.GetPayload().(*domainEvents.AuthenticatedEvent)
		user := payload.GetUser()
		now := time.Now()
		user.SetStatus(userEntities.UserVerified)
		user.SetVerified(now)
		h.UserRepository.Update(user)
	}
	wg.Done()
}
