package flow09_test

import (
	flow09 "pet-care/i-vaccine-dates"
	"pet-care/internal/state09"
	"testing"
)

func TestI(t *testing.T) {
	input := []string{"vaccine-date"}
	store := &state09.Store{}
	flow09.Capture(store, input)
	input[0] = "later-input"
	first := flow09.Read(store)
	first[0] = "later-read"
	second := flow09.Read(store)
	if len(second) != 1 || second[0] != "vaccine-date" {
		t.Fatalf("captured vaccine-dates values changed after later ownership mutations")
	}
}
