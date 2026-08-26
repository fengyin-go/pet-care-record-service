package flow27

import (
	"errors"
	"pet-care/internal/state27"
)

var ErrBadItem = errors.New("invalid record-chunk")

func Process(tracker *state27.Tracker, frames []string) (err error) {
	for _, frame := range frames {
		resource, openErr := tracker.Open()
		if openErr != nil {
			return openErr
		}
		defer func() { err = resource.Finish(err) }()
		if frame == "bad" {
			return ErrBadItem
		}
	}
	return nil
}
