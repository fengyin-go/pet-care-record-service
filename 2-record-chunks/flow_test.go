package flow27_test

import (
	"errors"
	flow27 "pet-care/2-record-chunks"
	"pet-care/internal/state27"
	"testing"
)

func Test2(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state27.NewTracker(1)
	err := flow27.Process(tracker, frames)
	if !errors.Is(err, flow27.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("record-chunks batch lost its item error or retained an open resource")
	}
}
