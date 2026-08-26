package flow28

import (
	"errors"
	"pet-care/internal/state28"
)

var ErrBadItem = errors.New("invalid vaccine-segment")

func Process(tracker *state28.Tracker, frames []string) (err error) {
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
