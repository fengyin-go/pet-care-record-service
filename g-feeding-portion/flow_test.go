package flow07_test

import (
	flow07 "pet-care/g-feeding-portion"
	"pet-care/internal/state07"
	"testing"
)

func TestG(t *testing.T) {
	input := []string{"portion"}
	store := &state07.Store{}
	flow07.Capture(store, input)
	input[0] = "later-input"
	first := flow07.Read(store)
	first[0] = "later-read"
	second := flow07.Read(store)
	if len(second) != 1 || second[0] != "portion" {
		t.Fatalf("captured feeding-portion values changed after later ownership mutations")
	}
}
