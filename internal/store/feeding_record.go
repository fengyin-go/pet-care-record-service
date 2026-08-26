package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreateFeedingRecord(f *model.FeedingRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.feedingRecords[f.ID] = f
	return nil
}

func (s *MemoryStore) GetFeedingRecord(id string) (*model.FeedingRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f, ok := s.feedingRecords[id]
	if !ok {
		return nil, ErrNotFound
	}
	return f, nil
}

func (s *MemoryStore) ListFeedingRecords() []*model.FeedingRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.FeedingRecord, 0, len(s.feedingRecords))
	for _, f := range s.feedingRecords {
		list = append(list, f)
	}
	return list
}

func (s *MemoryStore) UpdateFeedingRecord(f *model.FeedingRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.feedingRecords[f.ID]; !ok {
		return ErrNotFound
	}
	s.feedingRecords[f.ID] = f
	return nil
}

func (s *MemoryStore) DeleteFeedingRecord(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.feedingRecords[id]; !ok {
		return ErrNotFound
	}
	delete(s.feedingRecords, id)
	return nil
}

func (s *MemoryStore) DeleteFeedingRecordsByPet(petID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, f := range s.feedingRecords {
		if f.PetID == petID {
			delete(s.feedingRecords, id)
		}
	}
	return nil
}
