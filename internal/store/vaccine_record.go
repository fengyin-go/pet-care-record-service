package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreateVaccineRecord(v *model.VaccineRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.vaccineRecords[v.ID] = v
	return nil
}

func (s *MemoryStore) GetVaccineRecord(id string) (*model.VaccineRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.vaccineRecords[id]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (s *MemoryStore) ListVaccineRecords() []*model.VaccineRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.VaccineRecord, 0, len(s.vaccineRecords))
	for _, v := range s.vaccineRecords {
		list = append(list, v)
	}
	return list
}

func (s *MemoryStore) UpdateVaccineRecord(v *model.VaccineRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.vaccineRecords[v.ID]; !ok {
		return ErrNotFound
	}
	s.vaccineRecords[v.ID] = v
	return nil
}

func (s *MemoryStore) DeleteVaccineRecord(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.vaccineRecords[id]; !ok {
		return ErrNotFound
	}
	delete(s.vaccineRecords, id)
	return nil
}

func (s *MemoryStore) DeleteVaccineRecordsByPet(petID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, v := range s.vaccineRecords {
		if v.PetID == petID {
			delete(s.vaccineRecords, id)
		}
	}
	return nil
}
