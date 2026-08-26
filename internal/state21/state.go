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
		return nil
	}
	return &validator{enabled: true}
}
