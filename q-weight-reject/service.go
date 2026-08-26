package flow17

import (
	"errors"
	"pet-care/internal/state17"
)

func Forward(source *state17.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state17.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		var temporary *state17.Temporary
		if errors.As(err, &temporary) {
			continue
		}
		return err
	}
	return last
}
