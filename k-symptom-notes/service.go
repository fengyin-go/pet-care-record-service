package flow11

import "pet-care/internal/state11"

func Capture(store *state11.Store, values []string) {
	store.Replace(values)
}

func Read(store *state11.Store) []string { return store.Snapshot() }
