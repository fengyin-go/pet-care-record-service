package flow21

import "pet-care/internal/state21"

func Accept(gate state21.Gate, payload string) error {
	if gate == nil {
		return nil
	}
	return gate.Validate(payload)
}
