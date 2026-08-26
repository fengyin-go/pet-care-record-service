package flow08_test

import (
	flow08 "pet-care/h-weight-samples"
	"pet-care/internal/state08"
	"testing"
)

func TestH(t *testing.T) {
	input := []string{"weight-sample"}
	store := &state08.Store{}
	flow08.Capture(store, input)
	input[0] = "later-input"
	first := flow08.Read(store)
	first[0] = "later-read"
	second := flow08.Read(store)
	if len(second) != 1 || second[0] != "weight-sample" {
		t.Fatalf("captured weight-samples values changed after later ownership mutations")
	}
}
