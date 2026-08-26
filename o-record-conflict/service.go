package flow15

import "pet-care/internal/state15"

func Forward(source *state15.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state15.Normalize(source.Next())
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
