package flow13

import (
	"errors"
	"pet-care/internal/state13"
)

func Forward(source *state13.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state13.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		var temporary *state13.Temporary
		if errors.As(err, &temporary) {
			continue
		}
		return err
	}
	return last
}
