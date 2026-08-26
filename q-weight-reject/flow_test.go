package flow17_test

import (
	"errors"
	"pet-care/internal/state17"
	flow17 "pet-care/q-weight-reject"
	"testing"
)

func TestQ(t *testing.T) {
	source := state17.NewSource(&state17.Rejected{Reason: "weight-reject"}, nil)
	err := flow17.Forward(source, 2)
	var rejected *state17.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent weight-reject response was retried or lost its typed rejection")
	}
}
