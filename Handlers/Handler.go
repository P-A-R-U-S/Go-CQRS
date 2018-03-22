package Handlers

// Base interface for Handlers
type Handler interface {
	Event() string
	Execute(args ... interface{}) error
	OnSubscribe()
	OnUnsubscribe()
}