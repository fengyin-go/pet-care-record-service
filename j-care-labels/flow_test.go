package flow10_test

import (
	"pet-care/internal/state10"
	flow10 "pet-care/j-care-labels"
	"testing"
)

func TestJ(t *testing.T) {
	input := []string{"care-label"}
	store := &state10.Store{}
	flow10.Capture(store, input)
	input[0] = "later-input"
	first := flow10.Read(store)
	first[0] = "later-read"
	second := flow10.Read(store)
	if len(second) != 1 || second[0] != "care-label" {
		t.Fatalf("captured care-labels values changed after later ownership mutations")
	}
}
