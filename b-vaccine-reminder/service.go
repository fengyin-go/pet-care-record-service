package flow02

import (
	"context"
	"pet-care/internal/state02"
)

func Route(ctx context.Context, sink *state02.Sink, signal string) error {
	return sink.Deliver(ctx, signal)
}
