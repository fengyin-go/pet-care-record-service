package flow14

import "pet-care/internal/state14"

func Forward(source *state14.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state14.Normalize(source.Next())
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
