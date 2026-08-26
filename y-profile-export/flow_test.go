package flow25_test

import (
	"errors"
	"pet-care/internal/state25"
	flow25 "pet-care/y-profile-export"
	"testing"
)

func TestY(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state25.NewTracker(1)
	err := flow25.Process(tracker, frames)
	if !errors.Is(err, flow25.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("profile-export batch lost its item error or retained an open resource")
	}
}
