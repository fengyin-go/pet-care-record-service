package flow24_test

import (
	"pet-care/internal/state24"
	flow24 "pet-care/x-record-validator"
	"testing"
)

func callDisabledGate(gate state24.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow24.Accept(gate, payload)
	return err, false
}

func TestX(t *testing.T) {
	gate := state24.NewGate(false)
	err, panicked := callDisabledGate(gate, "record-validate")
	if panicked || err != nil {
		t.Fatalf("disabled record-validator gate rejected or panicked on an accepted signal")
	}
}
