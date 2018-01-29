package Test

import (
	"fmt"
	"errors"
	"time"
)

const EventFake1 = "_EventFake1"
const EventFake2 = "_EventFake2"

type FakeHandler1 struct {
	_name, _event string
	_isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired bool
	_isPanicOnEvent, _isPanicOnOnSubscribe, _isPanicOnOnUnsubscribe, _isPanicOnExecute bool
	_isDisableMessage bool
}

func (h *FakeHandler1) Event() string {
	if h._isPanicOnEvent {
		panic(errors.New( h._event  + ":Panic in Event"))
	}

	if len(h._event) > 0 {
		return h._event
	}


	return EventFake1
}
func (h *FakeHandler1) Execute(... interface{}) error {
	if h._isPanicOnExecute {
		panic(errors.New( h.Event()  + ":Panic in Execute"))
	}

	h._isExecuteFired = true

	if !h._isDisableMessage {
		fmt.Printf("Executed: %s : %s", h._name, h.Event())
		fmt.Println()
	}
	time.Sleep(time.Microsecond * 500)

	return nil
}
func (h *FakeHandler1) OnSubscribe() {
	if h._isPanicOnOnSubscribe {
		panic(errors.New( h.Event()  + ":Panic in OnSubscribe"))
	}

	h._isOnSubscribeFired = true
}
func (h *FakeHandler1) OnUnsubscribe() {
	if h._isPanicOnOnUnsubscribe {
		panic(errors.New( h.Event()  + ":Panic in OnUnsubscribe"))
	}

	h._isOnUnsubscribeFired = true
}

type FakeHandler2 struct {
	_name, _event                                                      					string
	_isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired 						bool
	_isPanicOnEvent, _isPanicOnOnSubscribe, _isPanicOnOnUnsubscribe, _isPanicOnExecute 	bool
	_isDisableMessage bool
}

func (h *FakeHandler2) Event() string {
	if h._isPanicOnEvent {
		panic(errors.New( h._event + ":Panic in Event"))
	}

	if len(h._event) > 0 {
		return h._event
	}

	return ""
}
func (h *FakeHandler2) Execute(... interface{}) error {
	if h._isPanicOnExecute{
		panic(errors.New( h.Event()  + ":Panic in Execute"))
	}

	h._isExecuteFired = true

	if !h._isDisableMessage {
		fmt.Printf("Executed: %s : %s", h._name, h.Event())
		fmt.Println()
	}
	time.Sleep(time.Microsecond * 500)
	return nil
}
func (h *FakeHandler2) OnSubscribe() {
	if h._isPanicOnOnSubscribe{
		panic(errors.New( h.Event()  + ":Panic in OnSubscribe"))
	}

	h._isOnSubscribeFired = true
}
func (h *FakeHandler2) OnUnsubscribe() {
	if h._isPanicOnOnUnsubscribe{
		panic(errors.New( h.Event()  + ":Panic in OnUnsubscribe"))
	}
	h._isOnUnsubscribeFired = true
}
