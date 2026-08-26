package flow16

import (
	"errors"
	"pet-care/internal/state16"
)

func Forward(source *state16.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state16.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		var temporary *state16.Temporary
		if errors.As(err, &temporary) {
			continue
		}
		return err
	}
	return last
}
