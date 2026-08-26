package state13

import "fmt"

type Rejected struct{ Reason string }

func (e *Rejected) Error() string { return "rejected: " + e.Reason }

type Temporary struct{ Reason string }

func (e *Temporary) Error() string { return "temporary: " + e.Reason }

type Source struct {
	steps []error
	calls int
}

func NewSource(steps ...error) *Source {
	return &Source{steps: append([]error(nil), steps...)}
}

func (s *Source) Next() error {
	s.calls++
	if len(s.steps) == 0 {
		return nil
	}
	err := s.steps[0]
	s.steps = s.steps[1:]
	return err
}

func (s *Source) Calls() int { return s.calls }

func Normalize(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("signal outcome: %v", err)
}
