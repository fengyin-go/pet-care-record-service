package flow25

import (
	"errors"
	"pet-care/internal/state25"
)

var ErrBadItem = errors.New("invalid profile-export")

func Process(tracker *state25.Tracker, frames []string) error {
	for _, frame := range frames {
		err := func() (err error) {
			resource, openErr := tracker.Open()
			if openErr != nil {
				return openErr
			}
			defer func() { err = resource.Finish(err) }()
			if frame == "bad" {
				return ErrBadItem
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
