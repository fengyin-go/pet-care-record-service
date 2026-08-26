package flow26

import (
	"errors"
	"pet-care/internal/state26"
)

var ErrBadItem = errors.New("invalid feeding-session")

func Process(tracker *state26.Tracker, frames []string) (err error) {
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
