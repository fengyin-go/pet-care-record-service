package flow10

import "pet-care/internal/state10"

func Capture(store *state10.Store, values []string) {
	store.Replace(values)
}

func Read(store *state10.Store) []string { return store.Snapshot() }
