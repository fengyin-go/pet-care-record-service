package flow08

import "pet-care/internal/state08"

func Capture(store *state08.Store, values []string) {
	snapshot := append([]string(nil), values...)
	store.Replace(snapshot)
}

func Read(store *state08.Store) []string { return store.Snapshot() }
