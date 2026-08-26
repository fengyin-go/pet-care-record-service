package flow13

import "pet-care/internal/state13"

func Forward(source *state13.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state13.Normalize(source.Next())
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
