package flow09

import "pet-care/internal/state09"

func Capture(store *state09.Store, values []string) {
	snapshot := append([]string(nil), values...)
	store.Replace(snapshot)
}

func Read(store *state09.Store) []string { return store.Snapshot() }
