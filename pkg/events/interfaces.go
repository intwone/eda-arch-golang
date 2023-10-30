package events

import (
	"sync"
	"time"
)

type Event struct {
	name      string
	payload   interface{}
	timestamp time.Time
}

func NewEvent(name string, payload interface{}) *Event {
	event := Event{
		name:      name,
		payload:   payload,
		timestamp: time.Now(),
	}

	return &event
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) GetPayload() interface{} {
	return e.payload
}

func (e *Event) GetTimestamp() time.Time {
	return e.timestamp
}

type EventHandlerInterface interface {
	Handle(event Event, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event Event) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
