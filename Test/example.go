package Test

import (
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