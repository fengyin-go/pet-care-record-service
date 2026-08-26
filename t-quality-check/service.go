package flow20

import "pet-care/internal/state20"

func Accept(gate state20.Gate, payload string) error {
	return gate.Validate(payload)
}
