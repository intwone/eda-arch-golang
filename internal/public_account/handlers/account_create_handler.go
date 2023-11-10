package handlers

import (
	"sync"

	repositories "github.com/intwone/eda-arch-golang/internal/private_database/interfaces"
	accountEntity "github.com/intwone/eda-arch-golang/internal/public_account/domain/entities"
	domainEvents "github.com/intwone/eda-arch-golang/internal/public_user/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type AccountCreateHandler struct {
	AccountRepository repositories.AccountRepositoryInterface
}

func NewAccountCreateHandler(ar repositories.AccountRepositoryInterface) *AccountCreateHandler {
	return &AccountCreateHandler{
		AccountRepository: ar,
	}
}

func (h *AccountCreateHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.UserCreatedEventName {
		account := accountEntity.NewAccountEntity(accountEntity.AccountB2B)
		h.AccountRepository.Create(*account)
	}
	wg.Done()
}
