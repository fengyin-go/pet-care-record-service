package flow05

import (
	"context"
	"pet-care/internal/state05"
)

func Route(ctx context.Context, sink *state05.Sink, signal string) error {
	return sink.Deliver(ctx, signal)
}
