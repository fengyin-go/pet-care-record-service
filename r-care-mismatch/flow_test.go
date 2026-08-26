package flow18_test

import (
	"errors"
	"pet-care/internal/state18"
	flow18 "pet-care/r-care-mismatch"
	"testing"
)

func TestR(t *testing.T) {
	source := state18.NewSource(&state18.Rejected{Reason: "care-mismatch"}, nil)
	err := flow18.Forward(source, 2)
	var rejected *state18.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent care-mismatch response was retried or lost its typed rejection")
	}
}
