package flow06

import (
	"context"
	"pet-care/internal/state06"
)

func Route(ctx context.Context, sink *state06.Sink, signal string) error {
	return sink.Deliver(context.Background(), signal)
}
