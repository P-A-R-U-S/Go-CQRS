package Test

import (
	"fmt"
)

const exampleEvent = "_EventExample"

type exampleHandler1 struct {
}

func (h *exampleHandler1) Event() string {

	return exampleEvent
}

func (h *exampleHandler1) Execute(...interface{}) error {

	fmt.Println("Run Execute...")

	return nil
}

func (h *exampleHandler1) OnSubscribe() {
	fmt.Println("Run OnSubscribe...")
}

func (h *exampleHandler1) OnUnsubscribe() {
	fmt.Println("Run OnUnsubscribe...")
}
