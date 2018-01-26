package Test

import (
	"testing"
	bus "Golang-CQRS/Bus"
	"time"
)

const EventFake1 = "_EventFake1"

type FakeHandler1 struct {
	_name string
	isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired bool
}

func (h *FakeHandler1) Event() string {
	return EventFake1
}
func (h *FakeHandler1) Execute(... interface{}) error {
	h._isExecuteFired = true
	return nil
}
func (h *FakeHandler1) OnSubscribe() {
	h.isOnSubscribeFired = true
}
func (h *FakeHandler1) OnUnsubscribe() {
	h._isOnUnsubscribeFired = true
}

func Test_Should_not_panic_when_create_instanse_of_EventBus(t *testing.T) {
	bus := bus.New()

	if bus == nil {
		t.Fail()
	}
}

func Test_Should_not_panic_when_Subscribe(t *testing.T)  {
	eventBus := bus.New()

	h := new(FakeHandler1)

	eventBus.Subscribe(h)

	if !h.isOnSubscribeFired  {
		t.Fail()
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

func Test_Should_not_panic_when_Unsubscribe(t *testing.T)  {
	eventBus := bus.New()

	h := new(FakeHandler1)

	eventBus.Subscribe(h)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Unsubscribe(h.Event())

	if !h.isOnSubscribeFired  {
		t.Fail()
	}
}
