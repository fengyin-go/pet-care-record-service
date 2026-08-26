package flow03

import (
	"context"
	"pet-care/internal/state03"
)

func Route(ctx context.Context, sink *state03.Sink, signal string) error {
	return sink.Deliver(context.Background(), signal)
}
