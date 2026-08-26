package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/internal/store"
	"pet-care/pkg/idgen"
)

func (s *Service) CreateFeedingRecord(input model.FeedingRecord) (*model.FeedingRecord, error) {
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
	if err := s.store.CreateFeedingRecord(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetFeedingRecord(id string) (*model.FeedingRecord, error) {
	return s.store.GetFeedingRecord(id)
}

func (s *Service) ListFeedingRecords(filter model.FeedingRecordFilter, page, size int) ([]*model.FeedingRecord, int, error) {
	all := s.store.ListFeedingRecords()
	matched := make([]*model.FeedingRecord, 0, len(all))
	for _, f := range all {
		if filter.Match(f) {
			matched = append(matched, f)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].FedAt.After(matched[j].FedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.FeedingRecord{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateFeedingRecord(id string, input model.FeedingRecord) (*model.FeedingRecord, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	f, err := s.store.GetFeedingRecord(id)
	if err != nil {
		return nil, err
	}
	if input.PetID != f.PetID {
		if _, err := s.store.GetPet(input.PetID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("pet_id", "宠物不存在")
			}
			return nil, err
		}
	}
	f.PetID = input.PetID
	f.Food = input.Food
	f.Amount = input.Amount
	f.FedAt = input.FedAt
	if err := s.store.UpdateFeedingRecord(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Service) DeleteFeedingRecord(id string) error {
	return s.store.DeleteFeedingRecord(id)
}
