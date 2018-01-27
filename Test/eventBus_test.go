package Test

import (
	"testing"
	bus "Golang-CQRS/Bus"
	"time"
	"github.com/pkg/errors"
	"fmt"
)

const EventFake1 = "_EventFake1"
const EventFake2 = "_EventFake2"

type FakeHandler1 struct {
	_name, _event string
	_isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired bool
	_isPanicOnEvent, _isPanicOnOnSubscribe, _isPanicOnOnUnsubscribe, _isPanicOnExecute bool
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
		panic(errors.New( h._event  + ":Panic in Execute"))
	}

	h._isExecuteFired = true
	fmt.Printf("Executed: %s : %s", h._name, h._event)
	fmt.Println()

	return nil
}
func (h *FakeHandler1) OnSubscribe() {
	if h._isPanicOnOnSubscribe {
		panic(errors.New( h._event  + ":Panic in OnSubscribe"))
	}

	h._isOnSubscribeFired = true
}
func (h *FakeHandler1) OnUnsubscribe() {
	if h._isPanicOnOnUnsubscribe {
		panic(errors.New( h._event  + ":Panic in OnUnsubscribe"))
	}

	h._isOnUnsubscribeFired = true
}

type FakeHandler2 struct {
	_name, _event                                                      					string
	_isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired 						bool
	_isPanicOnEvent, _isPanicOnOnSubscribe, _isPanicOnOnUnsubscribe, _isPanicOnExecute 	bool
}

func (h *FakeHandler2) Event() string {
	if h._isPanicOnEvent {
		panic(errors.New( h._event  + ":Panic in Event"))
	}

	if len(h._event) > 0 {
		return h._event
	}

	return ""
}
func (h *FakeHandler2) Execute(... interface{}) error {
	if h._isPanicOnExecute{
		panic(errors.New( h._event  + ":Panic in Execute"))
	}

	h._isExecuteFired = true

	fmt.Printf("Executed: %s : %s", h._name, h._event)
	fmt.Println()

	return nil
}
func (h *FakeHandler2) OnSubscribe() {
	if h._isPanicOnOnSubscribe{
		panic(errors.New( h._event  + ":Panic in OnSubscribe"))
	}

	h._isOnSubscribeFired = true
}
func (h *FakeHandler2) OnUnsubscribe() {
	if h._isPanicOnOnUnsubscribe{
		panic(errors.New( h._event  + ":Panic in OnUnsubscribe"))
	}
	h._isOnUnsubscribeFired = true
}

func Test_Should_not_panic_when_create_instance_of_EventBus(t *testing.T) {
	bus := bus.New()

	if bus == nil {
		t.Fail()
	}
}

func Test_Should_not_panic_when_Subscribe(t *testing.T)  {
	eventBus := bus.New()

	h := new(FakeHandler1)

	eventBus.Subscribe(h)

	if !h._isOnSubscribeFired  {
		t.Fail()
	}
}

func Test_Should_not_accept_nil_when_Subscribe(t *testing.T) {
	eventBus := bus.New()

	err :=  eventBus.Subscribe(nil)

	if err.Error() != "Handler can not be nil." {
		t.Errorf("Error Message should be: %s", "Handler can not be nil..." )
	}
}

func Test_Should_not_accept_handler_with_empty_event_when_Subscribe(t *testing.T) {
	eventBus := bus.New()

	h := new(FakeHandler2)

	err :=  eventBus.Subscribe(h)

	if err.Error() != "Handlers with empty Event are not allowed." {
		t.Errorf("Error Message should be: %s", "Handlers with empty Event are not allowed." )
	}
}


func Test_Should_not_panic_when_Publish(t *testing.T)  {
	eventBus := bus.New()

	h := new(FakeHandler1)

	eventBus.Subscribe(h)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	if !h._isExecuteFired  {
		t.Fail()
	}
}

func Test_Should_call_handler_by_event(t *testing.T)  {
	eventBus := bus.New()

	h1 := &FakeHandler1{_event:EventFake1}
	h2 := &FakeHandler2{_event:EventFake2}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should NOT be call for FakeHandler1
	if h1._isExecuteFired  {
		t.Error("Test1: Execute should NOT be call for FakeHandler1")
	}

	//Execute should be call for FakeHandler2
	if !h2._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler2")
	}

	//Clear state
	h1._isExecuteFired = false
	h2._isExecuteFired = false

	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test 2: Execute should be call for FakeHandler1")
	}

	//Execute should NOT be call for FakeHandler2
	if h2._isExecuteFired  {
		t.Error("Test 2: Execute should NOT be call for FakeHandler2")
	}

}

func Test_Should_call_all_handlers_with_same_event(t *testing.T)  {
	eventBus := bus.New()

	h1 := &FakeHandler1{_name: "Handler1", _event:EventFake1}
	h2 := &FakeHandler2{_name: "Handler2", _event:EventFake1}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should  be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//Execute should be call for FakeHandler2
	if !h2._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler2")
	}
}

func Test_Should_not_panic_when_Unsubscribe(t *testing.T)  {
	eventBus := bus.New()

	h := new(FakeHandler1)

	eventBus.Subscribe(h)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Unsubscribe(h.Event())

	if !h._isOnSubscribeFired  {
		t.Fail()
	}
}


func Test_Should_fail_when_one_of_handler_panic_on_Event(t *testing.T) {
	eventBus := bus.New()

	h1 := &FakeHandler1{_event:EventFake1}
	h2 := &FakeHandler2{_event:EventFake2, _isPanicOnEvent: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should NOT be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test1: Execute should NOT be call for FakeHandler1")
	}
}
