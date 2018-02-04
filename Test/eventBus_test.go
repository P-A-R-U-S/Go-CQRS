package Test

import (
	"testing"
	bus "Golang-CQRS/Bus"
	"time"
)

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


func Test_Should_not_fail_when_one_of_handler_panic_on_Event(t *testing.T) {
	eventBus := bus.New()

	h1 := &FakeHandler1{_event:EventFake1}
	h2 := &FakeHandler2{_event:EventFake2, _isPanicOnEvent: true}


	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(EventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should NOT be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//Execute should NOT be call for FakeHandler1
	if h2._isExecuteFired  {
		t.Error("Test1: Execute should NOT be call for FakeHandler2")
	}
}

func Test_Should_not_fail_when_one_of_handler_panic_on_OnSubscribe(t *testing.T) {
	eventBus := bus.New()

	h1 := &FakeHandler1{_event:EventFake1}
	h2 := &FakeHandler2{_event:EventFake2, _isPanicOnOnSubscribe: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)

	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(EventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//OnSubscribe should NOT be call for FakeHandler1
	if !h1._isOnSubscribeFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//Execute should NOT be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//OnSubscribe should be call for FakeHandler2
	if h2._isOnSubscribeFired  {
		t.Error("Test1: Execute should be call for FakeHandler2")
	}

	//Execute should NOT be call for FakeHandler2
	if h2._isExecuteFired  {
		t.Error("Test1: Execute should NOT be call for FakeHandler2")
	}
}

func Test_Should_not_fail_when_one_of_handler_panic_on_Execute(t *testing.T) {
	eventBus := bus.New()

	h1 := &FakeHandler1{_event:EventFake1}
	h2 := &FakeHandler2{_event:EventFake2, _isPanicOnExecute: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(EventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//OnSubscribe should NOT be call for FakeHandler1
	if !h1._isOnSubscribeFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//Execute should NOT be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//OnSubscribe should be call for FakeHandler2
	if !h2._isOnSubscribeFired  {
		t.Error("Test1: Execute should be call for FakeHandler2")
	}

	//Execute should NOT be call for FakeHandler2
	if h2._isExecuteFired  {
		t.Error("Test1: Execute should NOT be call for FakeHandler2")
	}
}

func Test_Should_not_fail_when_one_of_handler_panic_on_Unsubscribe(t *testing.T) {
	eventBus := bus.New()

	h1 := &FakeHandler1{_event:EventFake1}
	h2 := &FakeHandler2{_event:EventFake2, _isPanicOnOnUnsubscribe: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(EventFake2, 1,2,3,4,5,67,8,9)
	eventBus.Unsubscribe(EventFake1)
	eventBus.Unsubscribe(EventFake2)


	//Wait go routine complete
	time.Sleep(time.Second)

	//OnSubscribe should NOT be call for FakeHandler1
	if !h1._isOnSubscribeFired  {
		t.Error("Test1: OnSubscribe should be call for FakeHandler1")
	}

	//Execute should NOT be call for FakeHandler1
	if !h1._isExecuteFired  {
		t.Error("Test1: Execute should be call for FakeHandler1")
	}

	//Unsubscribe should NOT be call for FakeHandler1
	if !h1._isOnUnsubscribeFired  {
		t.Error("Test1: Unsubscribe should be call for FakeHandler1")
	}

	//OnSubscribe should be call for FakeHandler2
	if !h2._isOnSubscribeFired  {
		t.Error("Test1: OnSubscribe should be call for FakeHandler2")
	}

	//Execute should be call for FakeHandler2
	if !h2._isExecuteFired  {
		t.Error("Test1: Execute should  be call for FakeHandler2")
	}

	//Unsubscribe should NOT be call for FakeHandler2
	if !h2._isExecuteFired  {
		t.Error("Test1: Unsubscribe should  be call for FakeHandler2")
	}
}

func Test_Should_not_leak_args_changes_to_another_handler(t *testing.T){

	eventBus := bus.New()

	h1 := &FakeHandler1{_name: "FakeHandler1",  _event:EventFake1, _isAfterExecuteSleep:true, _delay: time.Second}
	h2 := &FakeHandler2{_name: "FakeHandler2",	_event:EventFake1, _isBeforeExecuteSleep:true, _delay: time.Second}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(EventFake1, 1, 2, 3)

	time.Sleep(time.Second * 2)

	if !(h1._argsChanges[0].(int) == 1001) {
		t.Fail()
	}
	if !(h1._argsChanges[1].(int) == 1002) {
		t.Fail()
	}
	if !(h1._argsChanges[2].(int) == 1003) {
		t.Fail()
	}

	if !(h2._argsChanges[0].(int) == 2001) {
		t.Fail()
	}
	if !(h2._argsChanges[1].(int) == 2002) {
		t.Fail()
	}
	if !(h2._argsChanges[2].(int) == 2003) {
		t.Fail()
	}
	time.Sleep(time.Second * 5)
}