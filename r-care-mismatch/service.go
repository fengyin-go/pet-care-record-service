package flow18

import (
	"errors"
	"pet-care/internal/state18"
)

func Forward(source *state18.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state18.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		var temporary *state18.Temporary
		if errors.As(err, &temporary) {
			continue
		}
		return err
	}
	return last
}
