package flow16

import "pet-care/internal/state16"

func Forward(source *state16.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state16.Normalize(source.Next())
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
