package state10

type Store struct{ values []string }

func (s *Store) Replace(values []string) { s.values = values }

func (s *Store) Snapshot() []string {
	return append([]string(nil), s.values...)
}
