package state08

type Store struct{ values []string }

func (s *Store) Replace(values []string) { s.values = values }

func (s *Store) Snapshot() []string {
	return s.values
}
