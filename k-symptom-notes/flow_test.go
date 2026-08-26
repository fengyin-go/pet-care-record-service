package flow11_test

import (
	"pet-care/internal/state11"
	flow11 "pet-care/k-symptom-notes"
	"testing"
)

func TestK(t *testing.T) {
	input := []string{"symptom"}
	store := &state11.Store{}
	flow11.Capture(store, input)
	input[0] = "later-input"
	first := flow11.Read(store)
	first[0] = "later-read"
	second := flow11.Read(store)
	if len(second) != 1 || second[0] != "symptom" {
		t.Fatalf("captured symptom-notes values changed after later ownership mutations")
	}
}
