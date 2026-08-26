package flow26

import (
	"errors"
	"pet-care/internal/state26"
)

var ErrBadItem = errors.New("invalid feeding-session")

func Process(tracker *state26.Tracker, frames []string) error {
	for _, frame := range frames {
		resource, openErr := tracker.Open()
		if openErr != nil {
			return openErr
		}
		// Close immediately after each item so the next Open() in the loop has
		// capacity; a deferred close here would accumulate until function exit
		// and exhaust the tracker mid-batch. ProcessItem releases the resource
		// and surfaces the per-item error together.
		if err := processItem(resource, frame); err != nil {
			return err
		}
	}
	return nil
}

func processItem(resource *state26.Resource, frame string) error {
	err := ErrBadItem
	if frame != "bad" {
		err = nil
	}
	return resource.Finish(err)
}
