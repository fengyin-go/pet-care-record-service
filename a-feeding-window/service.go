package flow01

import (
	"context"
	"pet-care/internal/state01"
)

func Route(ctx context.Context, sink *state01.Sink, signal string) error {
	return sink.Deliver(ctx, signal)
}
