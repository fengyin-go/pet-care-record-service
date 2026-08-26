package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreateSpecies(sp *model.Species) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.species {
		if exist.Name == sp.Name {
			return ErrConflict
		}
	}
	s.species[sp.ID] = sp
	return nil
}

func (s *MemoryStore) GetSpecies(id string) (*model.Species, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sp, ok := s.species[id]
	if !ok {
		return nil, ErrNotFound
	}
	return sp, nil
}

func (s *MemoryStore) GetSpeciesByName(name string) (*model.Species, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, sp := range s.species {
		if sp.Name == name {
			return sp, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) ListSpecies() []*model.Species {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Species, 0, len(s.species))
	for _, sp := range s.species {
		list = append(list, sp)
	}
	return list
}

func (s *MemoryStore) UpdateSpecies(sp *model.Species) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.species[sp.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.species {
		if exist.ID != sp.ID && exist.Name == sp.Name {
			return ErrConflict
		}
	}
	s.species[sp.ID] = sp
	return nil
}

func (s *MemoryStore) DeleteSpecies(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.species[id]; !ok {
		return ErrNotFound
	}
	delete(s.species, id)
	return nil
}
