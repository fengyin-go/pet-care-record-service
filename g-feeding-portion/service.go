package flow07

import "pet-care/internal/state07"

func Capture(store *state07.Store, values []string) {
	store.Replace(values)
}

func Read(store *state07.Store) []string { return store.Snapshot() }
