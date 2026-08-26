package flow29_test

import (
	"errors"
	flow29 "pet-care/4-weight-blocks"
	"pet-care/internal/state29"
	"testing"
)

func Test4(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state29.NewTracker(1)
	err := flow29.Process(tracker, frames)
	if !errors.Is(err, flow29.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("weight-blocks batch lost its item error or retained an open resource")
	}
}
