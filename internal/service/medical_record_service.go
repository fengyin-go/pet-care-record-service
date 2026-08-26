package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/internal/store"
	"pet-care/pkg/idgen"
)

func (s *Service) CreateMedicalRecord(input model.MedicalRecord) (*model.MedicalRecord, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetPet(input.PetID); err != nil {
		if err == store.ErrNotFound {
			return nil, model.NewValidationError("pet_id", "宠物不存在")
		}
		return nil, err
	}
	input.ID = idgen.Hex()
	input.CreatedAt = time.Now()
	if err := s.store.CreateMedicalRecord(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetMedicalRecord(id string) (*model.MedicalRecord, error) {
	return s.store.GetMedicalRecord(id)
}

func (s *Service) ListMedicalRecords(filter model.MedicalRecordFilter, page, size int) ([]*model.MedicalRecord, int, error) {
	all := s.store.ListMedicalRecords()
	matched := make([]*model.MedicalRecord, 0, len(all))
	for _, m := range all {
		if filter.Match(m) {
			matched = append(matched, m)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.MedicalRecord{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateMedicalRecord(id string, input model.MedicalRecord) (*model.MedicalRecord, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	m, err := s.store.GetMedicalRecord(id)
	if err != nil {
		return nil, err
	}
	if input.PetID != m.PetID {
		if _, err := s.store.GetPet(input.PetID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("pet_id", "宠物不存在")
			}
			return nil, err
		}
	}
	m.PetID = input.PetID
	m.VetName = input.VetName
	m.Diagnosis = input.Diagnosis
	m.Treatment = input.Treatment
	m.VisitDate = input.VisitDate
	m.Notes = input.Notes
	if err := s.store.UpdateMedicalRecord(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Service) DeleteMedicalRecord(id string) error {
	return s.store.DeleteMedicalRecord(id)
}
