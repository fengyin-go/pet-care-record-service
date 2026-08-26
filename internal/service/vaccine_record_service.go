package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/internal/store"
	"pet-care/pkg/idgen"
)

func (s *Service) CreateVaccineRecord(input model.VaccineRecord) (*model.VaccineRecord, error) {
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
	if err := s.store.CreateVaccineRecord(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetVaccineRecord(id string) (*model.VaccineRecord, error) {
	return s.store.GetVaccineRecord(id)
}

func (s *Service) ListVaccineRecords(filter model.VaccineRecordFilter, page, size int) ([]*model.VaccineRecord, int, error) {
	all := s.store.ListVaccineRecords()
	matched := make([]*model.VaccineRecord, 0, len(all))
	for _, v := range all {
		if filter.Match(v) {
			matched = append(matched, v)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.VaccineRecord{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateVaccineRecord(id string, input model.VaccineRecord) (*model.VaccineRecord, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	v, err := s.store.GetVaccineRecord(id)
	if err != nil {
		return nil, err
	}
	if input.PetID != v.PetID {
		if _, err := s.store.GetPet(input.PetID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("pet_id", "宠物不存在")
			}
			return nil, err
		}
	}
	v.PetID = input.PetID
	v.VaccineName = input.VaccineName
	v.Dose = input.Dose
	v.VaccinatedAt = input.VaccinatedAt
	v.NextDueAt = input.NextDueAt
	if err := s.store.UpdateVaccineRecord(v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Service) DeleteVaccineRecord(id string) error {
	return s.store.DeleteVaccineRecord(id)
}

func (s *Service) ListExpiringVaccines(days int) ([]*model.VaccineRecord, error) {
	all := s.store.ListVaccineRecords()
	now := time.Now()
	deadline := now.AddDate(0, 0, days)
	matched := make([]*model.VaccineRecord, 0)
	for _, v := range all {
		if v.Status(now) == model.VaccineStatusValid && (v.NextDueAt.Before(deadline) || v.NextDueAt.Equal(deadline)) {
			matched = append(matched, v)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].NextDueAt.Before(matched[j].NextDueAt)
	})
	return matched, nil
}

func (s *Service) ListExpiredVaccines() ([]*model.VaccineRecord, error) {
	all := s.store.ListVaccineRecords()
	now := time.Now()
	matched := make([]*model.VaccineRecord, 0)
	for _, v := range all {
		if v.Status(now) == model.VaccineStatusExpired {
			matched = append(matched, v)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].NextDueAt.Before(matched[j].NextDueAt)
	})
	return matched, nil
}
