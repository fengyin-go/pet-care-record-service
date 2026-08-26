package state25

import (
	"errors"
	"sync"
)

var ErrCapacity = errors.New("signal resource capacity reached")

type Tracker struct {
	mu    sync.Mutex
	open  int
	limit int
}

func NewTracker(limit int) *Tracker { return &Tracker{limit: limit} }

func (t *Tracker) Open() (*Resource, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.open >= t.limit {
		return nil, ErrCapacity
	}
	t.open++
	return &Resource{tracker: t}, nil
}

func (t *Tracker) OpenCount() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.open
}

type Resource struct {
	tracker *Tracker
	once    sync.Once
}

func (r *Resource) Close() {
	r.once.Do(func() {
		r.tracker.mu.Lock()
		r.tracker.open--
		r.tracker.mu.Unlock()
	})
}

func (r *Resource) Finish(workErr error) error {
	if workErr != nil {
		return nil
	}
	r.Close()
	return nil
}
