package flow01_test

import (
	"context"
	"errors"
	flow01 "pet-care/a-feeding-window"
	"pet-care/internal/state01"
	"testing"
)

func TestA(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state01.Sink{}
	err := flow01.Route(ctx, sink, "feeding")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled feeding-window signal was delivered without a cancellation result")
	}
}
