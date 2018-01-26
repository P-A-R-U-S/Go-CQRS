package Bus

import (
	"fmt"
	"reflect"
	"sync"
	"github.com/pkg/errors"
	handlers "Golang-CQRS/Handlers"
)

// EventBus implements publish/subscribe pattern: https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern
type EventBus interface {
	Publish(eventName string, args ...interface{})
	Subscribe(handler handlers.Handler) error
	Unsubscribe(eventName string) error
}

type handlersMap map[string][] handlers.Handler

type eventBus struct {
	mtx      sync.RWMutex
	handlers handlersMap
}

// Execute appropriate handlers
func (b *eventBus) Publish(eventName string, args ...interface{}) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	if hs, ok := b.handlers[eventName]; ok {
		rArgs := buildHandlerArgs(args)

		for _, h := range hs {
			go h.Execute(rArgs)
		}
	}
}

// Subscribe Handler
func (b *eventBus) Subscribe(h handlers.Handler) error {
	if h == nil {
		return errors.New("Handler can not be nil.")
	}

	if len(h.Event()) == 0 {
		return errors.New("Handlers with empty Event are not allowed.")
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()

	h.OnSubscribe()
	b.handlers[h.Event()] = append(b.handlers[h.Event()], h)

	return nil
}

// Unsubscribe Handler
func (b *eventBus) Unsubscribe(eventName string) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if _, ok := b.handlers[eventName]; ok {

		for i, h := range b.handlers[eventName] {
			if h != nil {
				h.OnUnsubscribe()
				b.handlers[eventName] = append(b.handlers[eventName][:i], b.handlers[eventName][i+1:]...)
			}
		}

		return nil
	}

	return fmt.Errorf("event are not %s exist", eventName)
}

func buildHandlerArgs(args []interface{}) []reflect.Value {
	reflectedArgs := make([]reflect.Value, 0)

	for _, arg := range args {
		reflectedArgs = append(reflectedArgs, reflect.ValueOf(arg))
	}

	return reflectedArgs
}

// New creates new EventBus
func New() EventBus {
	return &eventBus{
		handlers: make(handlersMap),
	}
}