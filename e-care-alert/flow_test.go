package flow05_test

import (
	"context"
	"errors"
	flow05 "pet-care/e-care-alert"
	"pet-care/internal/state05"
	"testing"
)

func TestE(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state05.Sink{}
	err := flow05.Route(ctx, sink, "care")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled care-alert signal was delivered without a cancellation result")
	}
}
