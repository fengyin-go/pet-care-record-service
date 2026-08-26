package flow19_test

import (
	"pet-care/internal/state19"
	flow19 "pet-care/s-feeding-rule"
	"testing"
)

func callDisabledGate(gate state19.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow19.Accept(gate, payload)
	return err, false
}

func TestS(t *testing.T) {
	gate := state19.NewGate(false)
	err, panicked := callDisabledGate(gate, "feeding-rule")
	if panicked || err != nil {
		t.Fatalf("disabled feeding-rule gate rejected or panicked on an accepted signal")
	}
}
