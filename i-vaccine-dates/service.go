package flow09

import "pet-care/internal/state09"

func Capture(store *state09.Store, values []string) {
	store.Replace(values)
}

func Read(store *state09.Store) []string { return store.Snapshot() }
