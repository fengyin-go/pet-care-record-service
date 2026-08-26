package flow18

import "pet-care/internal/state18"

func Forward(source *state18.Source, attempts int) error {
	var last error
	for i := 0; i < attempts; i++ {
		err := state18.Normalize(source.Next())
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
