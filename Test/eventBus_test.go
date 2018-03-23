package Test

import (
	"testing"
	bus "Golang-CQRS/Bus"
	"time"
)

func Test_Should_not_panic_when_create_instance_of_EventBus(t *testing.T) {
	eventBus := bus.New()

	if eventBus == nil {
		t.Fail()
	}
}

func Test_Should_not_panic_when_Subscribe(t *testing.T)  {
	eventBus := bus.New()

	h := new(fakeHandler1)

	eventBus.Subscribe(h)

	if !h.isOnSubscribeFired {
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

	h := new(fakeHandler2)

	err :=  eventBus.Subscribe(h)

	if err.Error() != "Handlers with empty Event are not allowed." {
		t.Errorf("Error Message should be: %s", "Handlers with empty Event are not allowed." )
	}
}


func Test_Should_not_panic_when_Publish(t *testing.T)  {
	eventBus := bus.New()

	h := new(fakeHandler1)

	eventBus.Subscribe(h)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	if !h.isExecuteFired {
		t.Fail()
	}
}

func Test_Should_call_handler_by_event(t *testing.T)  {
	eventBus := bus.New()

	h1 := &fakeHandler1{event: eventFake1}
	h2 := &fakeHandler2{event: eventFake2}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should NOT be call for fakeHandler1
	if h1.isExecuteFired {
		t.Error("Test1: Execute should NOT be call for fakeHandler1")
	}

	//Execute should be call for fakeHandler2
	if !h2.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler2")
	}

	//Clear state
	h1.isExecuteFired = false
	h2.isExecuteFired = false

	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test 2: Execute should be call for fakeHandler1")
	}

	//Execute should NOT be call for fakeHandler2
	if h2.isExecuteFired {
		t.Error("Test 2: Execute should NOT be call for fakeHandler2")
	}

}

func Test_Should_call_all_handlers_with_same_event(t *testing.T)  {
	eventBus := bus.New()

	h1 := &fakeHandler1{name: "Handler1", event: eventFake1}
	h2 := &fakeHandler2{name: "Handler2", event: eventFake1}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should  be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//Execute should be call for fakeHandler2
	if !h2.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler2")
	}
}

func Test_Should_not_panic_when_Unsubscribe(t *testing.T)  {
	eventBus := bus.New()

	h := new(fakeHandler1)

	eventBus.Subscribe(h)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Unsubscribe(h.Event())

	if !h.isOnSubscribeFired {
		t.Fail()
	}
}


func Test_Should_not_fail_when_one_of_handler_panic_on_Event(t *testing.T) {
	eventBus := bus.New()

	h1 := &fakeHandler1{event: eventFake1}
	h2 := &fakeHandler2{event: eventFake2, isPanicOnEvent: true}


	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(eventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should NOT be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//Execute should NOT be call for fakeHandler1
	if h2.isExecuteFired {
		t.Error("Test1: Execute should NOT be call for fakeHandler2")
	}
}

func Test_Should_not_fail_when_one_of_handler_panic_inside_goroutine_on_Event(t *testing.T) {
	eventBus := bus.New()

	h1 := &fakeHandler1{event: eventFake1}
	h2 := &fakeHandler2{event: eventFake2, isPanicOnEvent: true, isPanicFromGoroutine: false}


	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(eventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//Execute should NOT be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//Execute should NOT be call for fakeHandler1
	if h2.isExecuteFired {
		t.Error("Test1: Execute should NOT be call for fakeHandler2")
	}
}


func Test_Should_not_fail_when_one_of_handler_panic_on_OnSubscribe(t *testing.T) {
	eventBus := bus.New()

	h1 := &fakeHandler1{event: eventFake1}
	h2 := &fakeHandler2{event: eventFake2, isPanicOnOnSubscribe: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)

	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(eventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//OnSubscribe should NOT be call for fakeHandler1
	if !h1.isOnSubscribeFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//Execute should NOT be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//OnSubscribe should be call for fakeHandler2
	if h2.isOnSubscribeFired {
		t.Error("Test1: Execute should be call for fakeHandler2")
	}

	//Execute should NOT be call for fakeHandler2
	if h2.isExecuteFired {
		t.Error("Test1: Execute should NOT be call for fakeHandler2")
	}
}

func Test_Should_not_fail_when_one_of_handler_panic_on_Execute(t *testing.T) {
	eventBus := bus.New()

	h1 := &fakeHandler1{event: eventFake1}
	h2 := &fakeHandler2{event: eventFake2, isPanicOnExecute: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(eventFake2, 1,2,3,4,5,67,8,9)

	//Wait go routine complete
	time.Sleep(time.Second)

	//OnSubscribe should NOT be call for fakeHandler1
	if !h1.isOnSubscribeFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//Execute should NOT be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//OnSubscribe should be call for fakeHandler2
	if !h2.isOnSubscribeFired {
		t.Error("Test1: Execute should be call for fakeHandler2")
	}

	//Execute should NOT be call for fakeHandler2
	if h2.isExecuteFired {
		t.Error("Test1: Execute should NOT be call for fakeHandler2")
	}
}

func Test_Should_not_fail_when_one_of_handler_panic_on_Unsubscribe(t *testing.T) {
	eventBus := bus.New()

	h1 := &fakeHandler1{event: eventFake1}
	h2 := &fakeHandler2{event: eventFake2, isPanicOnOnUnsubscribe: true}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake1, 1,2,3,4,5,67,8,9)
	eventBus.Publish(eventFake2, 1,2,3,4,5,67,8,9)
	eventBus.Unsubscribe(eventFake1)
	eventBus.Unsubscribe(eventFake2)


	//Wait go routine complete
	time.Sleep(time.Second)

	//OnSubscribe should NOT be call for fakeHandler1
	if !h1.isOnSubscribeFired {
		t.Error("Test1: OnSubscribe should be call for fakeHandler1")
	}

	//Execute should NOT be call for fakeHandler1
	if !h1.isExecuteFired {
		t.Error("Test1: Execute should be call for fakeHandler1")
	}

	//Unsubscribe should NOT be call for fakeHandler1
	if !h1.isOnUnsubscribeFired {
		t.Error("Test1: Unsubscribe should be call for fakeHandler1")
	}

	//OnSubscribe should be call for fakeHandler2
	if !h2.isOnSubscribeFired {
		t.Error("Test1: OnSubscribe should be call for fakeHandler2")
	}

	//Execute should be call for fakeHandler2
	if !h2.isExecuteFired {
		t.Error("Test1: Execute should  be call for fakeHandler2")
	}

	//Unsubscribe should NOT be call for fakeHandler2
	if !h2.isExecuteFired {
		t.Error("Test1: Unsubscribe should  be call for fakeHandler2")
	}
}


func Test_Should_not_leak_args_changes_to_another_handler(t *testing.T){

	eventBus := bus.New()

	h1 := &fakeHandler1{name: "fakeHandler1",  event: eventFake1, isAfterExecuteSleep:true, delay: time.Second}
	h2 := &fakeHandler2{name: "fakeHandler2",	event: eventFake1, isBeforeExecuteSleep:true, delay: time.Second}

	eventBus.Subscribe(h1)
	eventBus.Subscribe(h2)
	eventBus.Publish(eventFake1, 1, 2, 3)

	time.Sleep(time.Second * 3)

	if !(h1.argsChanges[0].(int) == 1001) {
		t.Fail()
	}
	if !(h1.argsChanges[1].(int) == 1002) {
		t.Fail()
	}
	if !(h1.argsChanges[2].(int) == 1003) {
		t.Fail()
	}

	if !(h2.argsChanges[0].(int) == 2001) {
		t.Fail()
	}
	if !(h2.argsChanges[1].(int) == 2002) {
		t.Fail()
	}
	if !(h2.argsChanges[2].(int) == 2003) {
		t.Fail()
	}
	time.Sleep(time.Second * 5)
}