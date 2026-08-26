package flow16_test

import (
	"errors"
	"pet-care/internal/state16"
	flow16 "pet-care/p-profile-reject"
	"testing"
)

func TestP(t *testing.T) {
	source := state16.NewSource(&state16.Rejected{Reason: "profile-reject"}, nil)
	err := flow16.Forward(source, 2)
	var rejected *state16.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent profile-reject response was retried or lost its typed rejection")
	}
}
