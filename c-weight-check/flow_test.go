package flow03_test

import (
	"context"
	"errors"
	flow03 "pet-care/c-weight-check"
	"pet-care/internal/state03"
	"testing"
)

func TestC(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state03.Sink{}
	err := flow03.Route(ctx, sink, "weight")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled weight-check signal was delivered without a cancellation result")
	}
}
