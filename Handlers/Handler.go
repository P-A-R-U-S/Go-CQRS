package Handlers

type Handler interface {
	Name() string
	Execute(... interface{}) error
	OnSubscribe()
	OnUnsubscribe()
}