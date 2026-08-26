package flow12_test

import (
	"pet-care/internal/state12"
	flow12 "pet-care/l-record-window"
	"testing"
)

func TestL(t *testing.T) {
	input := []string{"record-window"}
	store := &state12.Store{}
	flow12.Capture(store, input)
	input[0] = "later-input"
	first := flow12.Read(store)
	first[0] = "later-read"
	second := flow12.Read(store)
	if len(second) != 1 || second[0] != "record-window" {
		t.Fatalf("captured record-window values changed after later ownership mutations")
	}
}
