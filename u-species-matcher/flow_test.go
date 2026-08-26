package flow21_test

import (
	"pet-care/internal/state21"
	flow21 "pet-care/u-species-matcher"
	"testing"
)

func callDisabledGate(gate state21.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow21.Accept(gate, payload)
	return err, false
}

func TestU(t *testing.T) {
	gate := state21.NewGate(false)
	err, panicked := callDisabledGate(gate, "species-match")
	if panicked || err != nil {
		t.Fatalf("disabled species-matcher gate rejected or panicked on an accepted signal")
	}
}
