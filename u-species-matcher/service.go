package flow21

import "pet-care/internal/state21"

func Accept(gate state21.Gate, payload string) error {
	if gate == nil {
		// 品种匹配停用时安全跳过，放行信号给后续步骤。
		return nil
	}
	return gate.Validate(payload)
}
