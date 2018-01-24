package Test

import (
	"testing"
	bus "Golang-CQRS/Bus"
	)

func TestNew(t *testing.T) {
	eventBus := bus.New()

	if eventBus == nil {
		t.Fail()
	}
}


