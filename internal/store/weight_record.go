package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreateWeightRecord(w *model.WeightRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.weightRecords[w.ID] = w
	return nil
}

func (s *MemoryStore) GetWeightRecord(id string) (*model.WeightRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	w, ok := s.weightRecords[id]
	if !ok {
		return nil, ErrNotFound
	}
	return w, nil
}

func (s *MemoryStore) ListWeightRecords() []*model.WeightRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.WeightRecord, 0, len(s.weightRecords))
	for _, w := range s.weightRecords {
		list = append(list, w)
	}
	return list
}

func (s *MemoryStore) UpdateWeightRecord(w *model.WeightRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.weightRecords[w.ID]; !ok {
		return ErrNotFound
	}
	s.weightRecords[w.ID] = w
	return nil
}

func (s *MemoryStore) DeleteWeightRecord(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.weightRecords[id]; !ok {
		return ErrNotFound
	}
	delete(s.weightRecords, id)
	return nil
}

func (s *MemoryStore) DeleteWeightRecordsByPet(petID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, w := range s.weightRecords {
		if w.PetID == petID {
			delete(s.weightRecords, id)
		}
	}
	return nil
}
