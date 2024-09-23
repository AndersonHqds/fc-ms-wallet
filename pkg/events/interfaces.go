package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(event string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(event string, handler EventHandlerInterface) error
	Has(event string, handler EventHandlerInterface) bool
	Clear()
}
