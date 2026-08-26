package flow23_test

import (
	"pet-care/internal/state23"
	flow23 "pet-care/w-owner-selector"
	"testing"
)

func callDisabledGate(gate state23.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow23.Accept(gate, payload)
	return err, false
}

func TestW(t *testing.T) {
	gate := state23.NewGate(false)
	err, panicked := callDisabledGate(gate, "owner-select")
	if panicked || err != nil {
		t.Fatalf("disabled owner-selector gate rejected or panicked on an accepted signal")
	}
}
