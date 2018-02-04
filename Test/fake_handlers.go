package Test

import (
	"fmt"
	"errors"
	"time"
	"reflect"
)

const EventFake1 = "_EventFake1"
const EventFake2 = "_EventFake2"

type FakeHandler1 struct {
	_name, _event string
	_isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired bool
	_isPanicOnEvent, _isPanicOnOnSubscribe, _isPanicOnOnUnsubscribe, _isPanicOnExecute bool
	_isDisableMessage, _isBeforeExecuteSleep, _isAfterExecuteSleep bool
	_delay time.Duration
	_argsChanges []interface{}
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
func (h *FakeHandler1) Execute(args ... interface{}) error {
	fmt.Printf("--> %s : %s Args before changes %d\n", h._name, h.Event(), args)

	if h._isBeforeExecuteSleep {
		time.Sleep(h._delay)
	}

	if h._isPanicOnExecute {
		panic(errors.New( h.Event()  + ":Panic in Execute"))
	}

	for i := 0; i < len(args); i++ {
		if _, ok := args[i].(int); ok {
			args[i] = args[i].(int) + 1000
		}
	}

	h._isExecuteFired = true


	time.Sleep(time.Microsecond * 500)

	if h._isAfterExecuteSleep {
		time.Sleep(h._delay)
	}


	if !h._isDisableMessage {
		fmt.Printf("Executed: %s : %s", h._name, h.Event())
		fmt.Println()
	}

	fmt.Printf("--> %s : %s Args after changes %d\n", h._name, h.Event(), args)
	h._argsChanges = make([]interface{}, len(args))
	for i, arg := range args {
		h._argsChanges[i] = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	}

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
	_isDisableMessage, _isBeforeExecuteSleep, _isAfterExecuteSleep bool
	_delay time.Duration
	_argsChanges []interface{}
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
func (h *FakeHandler2) Execute(args ... interface{}) error {
	fmt.Printf("--> %s : %s Args before changes %d\n", h._name, h.Event(), args)

	if !h._isDisableMessage {
		fmt.Printf("Executed: %s : %s", h._name, h.Event())
		fmt.Println()
	}

	if h._isBeforeExecuteSleep {
		time.Sleep(h._delay)
	}

	if h._isPanicOnExecute{
		panic(errors.New( h.Event()  + ":Panic in Execute"))
	}

	for i := 0; i < len(args); i++ {
		if _, ok := args[i].(int); ok {
			args[i] = args[i].(int) + 2000
		}
	}

	h._isExecuteFired = true

	time.Sleep(time.Microsecond * 500)

	if h._isAfterExecuteSleep {
		time.Sleep(h._delay)
	}

	fmt.Printf("--> %s : %s Args after changes %d\n", h._name, h.Event(), args)
	h._argsChanges = make([]interface{}, len(args))
	for i, arg := range args {
		h._argsChanges[i] = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	}

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