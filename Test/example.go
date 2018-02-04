package Test

import (
	bus "Golang-CQRS/Bus"
	"fmt"
)

const ExampleEvent = "_EventExample"

type ExampleHandler1 struct {
}

func (h *ExampleHandler1) Event() string {

	return ExampleEvent
}
func (h *ExampleHandler1) Execute(... interface{}) error {

	fmt.Println("Run Execute...")

	return nil
}
func (h *ExampleHandler1) OnSubscribe() {
	fmt.Println("Run OnSubscribe...")
}
func (h *ExampleHandler1) OnUnsubscribe() {
	fmt.Println("Run OnUnsubscribe...")
}

func main()  {

	eventBus := bus.New()

	h := &ExampleHandler1{}

	eventBus.Subscribe(h)

	eventBus.Publish(ExampleEvent, 1, 2, "Test Message", 4.5)
	eventBus.Publish(ExampleEvent, 11, 22, "Test Message2", 14.5)

	eventBus.Unsubscribe(ExampleEvent)
}
