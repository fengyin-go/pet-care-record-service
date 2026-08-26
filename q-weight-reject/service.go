package flow17

import "pet-care/internal/state17"

func Forward(source *state17.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state17.Normalize(source.Next())
		if err == nil {
			return nil
		}
		last = err
		if err != nil {
			continue
		}
	}
	return last
}
