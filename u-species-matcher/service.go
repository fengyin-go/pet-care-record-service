package flow21

import "pet-care/internal/state21"

func Accept(gate state21.Gate, payload string) error {
	return gate.Validate(payload)
}
