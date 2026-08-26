package state21

import "errors"

type Gate interface{ Validate(string) error }

type validator struct{ enabled bool }

func (v *validator) Validate(payload string) error {
	if !v.enabled {
		return errors.New("disabled gate invoked")
	}
	if payload == "" {
		return errors.New("empty payload")
	}
	return nil
}

func NewGate(enabled bool) Gate {
	if !enabled {
		// 返回真正的 nil interface，避免"非空 interface 包 nil 指针"导致
		// 调用方在 nil 检查后仍触发方法调用而崩溃。停用即安全跳过。
		return nil
	}
	return &validator{enabled: true}
}
