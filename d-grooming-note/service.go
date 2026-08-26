package flow04

import (
	"context"
	"pet-care/internal/state04"
)

func Route(ctx context.Context, sink *state04.Sink, signal string) error {
	return sink.Deliver(ctx, signal)
}
