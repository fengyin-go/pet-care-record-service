package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreateOwner(o *model.Owner) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.owners {
		if exist.Email == o.Email {
			return ErrConflict
		}
	}
	s.owners[o.ID] = o
	return nil
}

func (s *MemoryStore) GetOwner(id string) (*model.Owner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	o, ok := s.owners[id]
	if !ok {
		return nil, ErrNotFound
	}
	return o, nil
}

func (s *MemoryStore) GetOwnerByEmail(email string) (*model.Owner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, o := range s.owners {
		if o.Email == email {
			return o, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) ListOwners() []*model.Owner {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Owner, 0, len(s.owners))
	for _, o := range s.owners {
		list = append(list, o)
	}
	return list
}

func (s *MemoryStore) UpdateOwner(o *model.Owner) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.owners[o.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.owners {
		if exist.ID != o.ID && exist.Email == o.Email {
			return ErrConflict
		}
	}
	s.owners[o.ID] = o
	return nil
}

func (s *MemoryStore) DeleteOwner(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.owners[id]; !ok {
		return ErrNotFound
	}
	delete(s.owners, id)
	return nil
}
