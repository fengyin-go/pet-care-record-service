package flow06_test

import (
	"context"
	"errors"
	flow06 "pet-care/f-profile-refresh"
	"pet-care/internal/state06"
	"testing"
)

func TestF(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state06.Sink{}
	err := flow06.Route(ctx, sink, "profile")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled profile-refresh signal was delivered without a cancellation result")
	}
}
