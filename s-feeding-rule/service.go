package flow19

import "pet-care/internal/state19"

func Accept(gate state19.Gate, payload string) error {
	return gate.Validate(payload)
}
