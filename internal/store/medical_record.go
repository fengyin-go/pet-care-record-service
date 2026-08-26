package store

import (
	"pet-care/internal/model"
)

func (s *MemoryStore) CreateMedicalRecord(m *model.MedicalRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.medicalRecords[m.ID] = m
	return nil
}

func (s *MemoryStore) GetMedicalRecord(id string) (*model.MedicalRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m, ok := s.medicalRecords[id]
	if !ok {
		return nil, ErrNotFound
	}
	return m, nil
}

func (s *MemoryStore) ListMedicalRecords() []*model.MedicalRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.MedicalRecord, 0, len(s.medicalRecords))
	for _, m := range s.medicalRecords {
		list = append(list, m)
	}
	return list
}

func (s *MemoryStore) UpdateMedicalRecord(m *model.MedicalRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.medicalRecords[m.ID]; !ok {
		return ErrNotFound
	}
	s.medicalRecords[m.ID] = m
	return nil
}

func (s *MemoryStore) DeleteMedicalRecord(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.medicalRecords[id]; !ok {
		return ErrNotFound
	}
	delete(s.medicalRecords, id)
	return nil
}

func (s *MemoryStore) DeleteMedicalRecordsByPet(petID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, m := range s.medicalRecords {
		if m.PetID == petID {
			delete(s.medicalRecords, id)
		}
	}
	return nil
}
