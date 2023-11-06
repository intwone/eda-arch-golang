package messenger

import (
	"fmt"
	"sync"

	domainEvents "github.com/intwone/eda-arch-golang/internal/domain/password/events"
	"github.com/intwone/eda-arch-golang/pkg/events"
)

type PasswordCreatedEmailDispatchHandler struct{}

func NewPasswordCreatedEmailDispatchHandler() *PasswordCreatedEmailDispatchHandler {
	return &PasswordCreatedEmailDispatchHandler{}
}

func (h *PasswordCreatedEmailDispatchHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	if event.GetName() == domainEvents.PasswordCreatedEventName {
		payload := event.GetPayload().(*domainEvents.PasswordCreatedEvent)
		fmt.Println(payload.GetUnhash())
	}
	wg.Done()
}
