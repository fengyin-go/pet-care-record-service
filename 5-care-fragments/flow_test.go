package flow30_test

import (
	"errors"
	flow30 "pet-care/5-care-fragments"
	"pet-care/internal/state30"
	"testing"
)

func Test5(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state30.NewTracker(1)
	err := flow30.Process(tracker, frames)
	if !errors.Is(err, flow30.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("care-fragments batch lost its item error or retained an open resource")
	}
}
