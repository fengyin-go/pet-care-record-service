package flow20

import "pet-care/internal/state20"

func Accept(gate state20.Gate, payload string) error {
	if gate == nil {
		return nil
	}
	return gate.Validate(payload)
}
