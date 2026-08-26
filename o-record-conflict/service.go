package flow15

import (
	"errors"
	"pet-care/internal/state15"
)

func Forward(source *state15.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state15.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		var temporary *state15.Temporary
		if errors.As(err, &temporary) {
			continue
		}
		return err
	}
	return last
}
