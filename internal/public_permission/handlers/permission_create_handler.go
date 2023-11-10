package handlers

import (
	"sync"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	permissionEntity "github.com/intwone/eda-arch-golang/internal/public_permission/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type PermissionCreateHandler struct {
	PermissionRepository repositories.PermissionRepositoryInterface
}

func NewPermissionCreateHandler(pr repositories.PermissionRepositoryInterface) *PermissionCreateHandler {
	return &PermissionCreateHandler{
		PermissionRepository: pr,
	}
}

func (h *PermissionCreateHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.UserCreatedEventName {
		payload := event.GetPayload().(*domainEvents.UserCreatedEvent)
		userID := payload.GetUserID()
		companyID := payload.GetCompanyID()
		permission := permissionEntity.NewPermissionEntity(permissionEntity.PermissionUser, userID, userID, companyID)
		h.PermissionRepository.Create(*permission)
	}
	wg.Done()
}
