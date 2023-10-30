package events

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEventHandler struct {
	ID int
}

func NewTestEventHandler(id int) *TestEventHandler {
	eventHandlerTest := TestEventHandler{
		ID: id,
	}

	return &eventHandlerTest
}

func (h *TestEventHandler) Handle(event Event, wg *sync.WaitGroup) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event1          *Event
	event2          *Event
	handler1        *TestEventHandler
	handler2        *TestEventHandler
	handler3        *TestEventHandler
	eventDispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.event1 = NewEvent("event1", "any")
	s.event2 = NewEvent("event2", "any")
	s.handler1 = NewTestEventHandler(1)
	s.handler2 = NewTestEventHandler(2)
	s.handler3 = NewTestEventHandler(3)
	s.eventDispatcher = NewEventDispatcher()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	s.Run("should be able to register a handler", func() {
		err := s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
		s.Nil(err)
		s.Equal(1, len(s.eventDispatcher.handlers[s.event1.GetName()]))

		s.eventDispatcher = NewEventDispatcher()
	})

	s.Run("should be able to register one or more handlers", func() {
		err := s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
		s.Nil(err)

		err = s.eventDispatcher.Register(s.event1.GetName(), s.handler2)
		s.Nil(err)

		s.Equal(2, len(s.eventDispatcher.handlers[s.event1.GetName()]))

		s.eventDispatcher = NewEventDispatcher()
	})

	s.Run("should not be able to register two handlers with the same name", func() {
		err := s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
		s.Nil(err)

		err = s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
		s.Equal(1, len(s.eventDispatcher.handlers[s.event1.GetName()]))
		s.Equal(err, HandlerAlreadRegisteredError)

		s.eventDispatcher = NewEventDispatcher()
	})
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	s.Run("should be able to clean all handlers", func() {
		err := s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
		s.Nil(err)

		err = s.eventDispatcher.Register(s.event1.GetName(), s.handler2)
		s.Nil(err)

		err = s.eventDispatcher.Register(s.event2.GetName(), s.handler3)
		s.Nil(err)

		s.eventDispatcher.Clear()
		s.Equal(0, len(s.eventDispatcher.handlers))
	})
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	s.Run("should be able to clean all handlers", func() {
		err := s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
		s.Nil(err)

		err = s.eventDispatcher.Register(s.event1.GetName(), s.handler2)
		s.Nil(err)

		s.True(s.eventDispatcher.Has(s.event1.GetName(), s.handler1))
		s.True(s.eventDispatcher.Has(s.event1.GetName(), s.handler2))
		s.False(s.eventDispatcher.Has(s.event1.GetName(), s.handler3))
	})
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event Event, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", *s.event1).Return(nil)
	s.eventDispatcher.Register(s.event1.GetName(), eventHandler)

	s.eventDispatcher.Dispatch(*s.event1)
	eventHandler.AssertExpectations(s.T())
	eventHandler.AssertNumberOfCalls(s.T(), "Handle", 1)
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := s.eventDispatcher.Register(s.event1.GetName(), s.handler1)
	s.Nil(err)

	err = s.eventDispatcher.Register(s.event1.GetName(), s.handler2)
	s.Nil(err)

	err = s.eventDispatcher.Register(s.event2.GetName(), s.handler3)
	s.Nil(err)

	s.eventDispatcher.Remove(s.event1.GetName(), s.handler1)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event1.GetName()]))
	s.Equal(s.handler2, s.eventDispatcher.handlers[s.event1.GetName()][0])

	s.eventDispatcher.Remove(s.event1.GetName(), s.handler2)
	s.Equal(0, len(s.eventDispatcher.handlers[s.event1.GetName()]))
}
