package flow13_test

import (
	"errors"
	"pet-care/internal/state13"
	flow13 "pet-care/m-vaccine-sync"
	"testing"
)

func TestM(t *testing.T) {
	source := state13.NewSource(&state13.Rejected{Reason: "vaccine-sync"}, nil)
	err := flow13.Forward(source, 2)
	var rejected *state13.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent vaccine-sync response was retried or lost its typed rejection")
	}
}
