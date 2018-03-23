package Handlers

// Handler is base interface for handlers
type Handler interface {
	Event() string
	Execute(args ...interface{}) error
	OnSubscribe()
	OnUnsubscribe()
}
