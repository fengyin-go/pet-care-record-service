package flow15_test

import (
	"errors"
	"pet-care/internal/state15"
	flow15 "pet-care/o-record-conflict"
	"testing"
)

func TestO(t *testing.T) {
	source := state15.NewSource(&state15.Rejected{Reason: "record-conflict"}, nil)
	err := flow15.Forward(source, 2)
	var rejected *state15.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent record-conflict response was retried or lost its typed rejection")
	}
}
