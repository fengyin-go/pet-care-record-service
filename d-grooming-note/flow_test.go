package flow04_test

import (
	"context"
	"errors"
	flow04 "pet-care/d-grooming-note"
	"pet-care/internal/state04"
	"testing"
)

func TestD(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state04.Sink{}
	err := flow04.Route(ctx, sink, "grooming")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled grooming-note signal was delivered without a cancellation result")
	}
}
