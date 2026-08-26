package flow02_test

import (
	"context"
	"errors"
	flow02 "pet-care/b-vaccine-reminder"
	"pet-care/internal/state02"
	"testing"
)

func TestB(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state02.Sink{}
	err := flow02.Route(ctx, sink, "vaccine")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled vaccine-reminder signal was delivered without a cancellation result")
	}
}
