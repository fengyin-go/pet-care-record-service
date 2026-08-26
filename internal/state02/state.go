package state02

import (
	"context"
	"sync/atomic"
)

type Sink struct{ calls atomic.Int32 }

func (s *Sink) Deliver(ctx context.Context, signal string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.calls.Add(1)
	return nil
}

func (s *Sink) Calls() int { return int(s.calls.Load()) }
