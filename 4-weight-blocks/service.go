package flow29

import (
	"errors"
	"pet-care/internal/state29"
)

var ErrBadItem = errors.New("invalid weight-block")

func Process(tracker *state29.Tracker, frames []string) (err error) {
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
