package flow28_test

import (
	"errors"
	flow28 "pet-care/3-vaccine-segments"
	"pet-care/internal/state28"
	"testing"
)

func Test3(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state28.NewTracker(1)
	err := flow28.Process(tracker, frames)
	if !errors.Is(err, flow28.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("vaccine-segments batch lost its item error or retained an open resource")
	}
}
