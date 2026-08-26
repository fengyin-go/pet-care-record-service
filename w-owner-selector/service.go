package flow23

import "pet-care/internal/state23"

func Accept(gate state23.Gate, payload string) error {
	return gate.Validate(payload)
}
