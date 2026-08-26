package flow14_test

import (
	"errors"
	"pet-care/internal/state14"
	flow14 "pet-care/n-feeding-limit"
	"testing"
)

func TestN(t *testing.T) {
	source := state14.NewSource(&state14.Rejected{Reason: "feeding-limit"}, nil)
	err := flow14.Forward(source, 2)
	var rejected *state14.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent feeding-limit response was retried or lost its typed rejection")
	}
}
