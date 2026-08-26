package flow24

import "pet-care/internal/state24"

func Accept(gate state24.Gate, payload string) error {
	return gate.Validate(payload)
}
