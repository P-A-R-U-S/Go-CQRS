package Handlers

type Handler interface {
	Event() string
	Execute(... interface{}) error
	OnSubscribe()
	OnUnsubscribe()
}