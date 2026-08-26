package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/internal/store"
	"pet-care/pkg/idgen"
)

func (s *Service) CreateWeightRecord(input model.WeightRecord) (*model.WeightRecord, error) {
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
	if err := s.store.CreateWeightRecord(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetWeightRecord(id string) (*model.WeightRecord, error) {
	return s.store.GetWeightRecord(id)
}

func (s *Service) ListWeightRecords(filter model.WeightRecordFilter, page, size int) ([]*model.WeightRecord, int, error) {
	all := s.store.ListWeightRecords()
	matched := make([]*model.WeightRecord, 0, len(all))
	for _, w := range all {
		if filter.Match(w) {
			matched = append(matched, w)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].MeasuredAt.After(matched[j].MeasuredAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.WeightRecord{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) GetWeightTrend(petID string) ([]*model.WeightRecord, error) {
	all := s.store.ListWeightRecords()
	matched := make([]*model.WeightRecord, 0)
	for _, w := range all {
		if w.PetID == petID {
			matched = append(matched, w)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].MeasuredAt.Before(matched[j].MeasuredAt)
	})
	return matched, nil
}

func (s *Service) UpdateWeightRecord(id string, input model.WeightRecord) (*model.WeightRecord, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	w, err := s.store.GetWeightRecord(id)
	if err != nil {
		return nil, err
	}
	if input.PetID != w.PetID {
		if _, err := s.store.GetPet(input.PetID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("pet_id", "宠物不存在")
			}
			return nil, err
		}
	}
	w.PetID = input.PetID
	w.Weight = input.Weight
	w.MeasuredAt = input.MeasuredAt
	if err := s.store.UpdateWeightRecord(w); err != nil {
		return nil, err
	}
	return w, nil
}

func (s *Service) DeleteWeightRecord(id string) error {
	return s.store.DeleteWeightRecord(id)
}
