package flow14

import (
	"errors"
	"pet-care/internal/state14"
)

func Forward(source *state14.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state14.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		var temporary *state14.Temporary
		if errors.As(err, &temporary) {
			continue
		}
		return err
	}
	return last
}
