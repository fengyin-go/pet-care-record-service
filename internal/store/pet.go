package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreatePet(p *model.Pet) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pets[p.ID] = p
	return nil
}

func (s *MemoryStore) GetPet(id string) (*model.Pet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.pets[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (s *MemoryStore) ListPets() []*model.Pet {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Pet, 0, len(s.pets))
	for _, p := range s.pets {
		list = append(list, p)
	}
	return list
}

func (s *MemoryStore) UpdatePet(p *model.Pet) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.pets[p.ID]; !ok {
		return ErrNotFound
	}
	s.pets[p.ID] = p
	return nil
}

func (s *MemoryStore) DeletePet(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.pets[id]; !ok {
		return ErrNotFound
	}
	delete(s.pets, id)
	return nil
}
