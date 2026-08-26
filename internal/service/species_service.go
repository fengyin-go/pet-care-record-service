package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/pkg/idgen"
)

func (s *Service) CreateSpecies(input model.Species) (*model.Species, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if existing, _ := s.store.GetSpeciesByName(input.Name); existing != nil {
		return nil, model.NewValidationError("name", "品种名称已存在")
	}
	input.ID = idgen.Hex()
	input.CreatedAt = time.Now()
	if err := s.store.CreateSpecies(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetSpecies(id string) (*model.Species, error) {
	return s.store.GetSpecies(id)
}

func (s *Service) GetSpeciesByName(name string) (*model.Species, error) {
	return s.store.GetSpeciesByName(name)
}

func (s *Service) ListSpecies(filter model.SpeciesFilter, page, size int) ([]*model.Species, int, error) {
	all := s.store.ListSpecies()
	matched := make([]*model.Species, 0, len(all))
	for _, sp := range all {
		if filter.Match(sp) {
			matched = append(matched, sp)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Species{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateSpecies(id string, input model.Species) (*model.Species, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	sp, err := s.store.GetSpecies(id)
	if err != nil {
		return nil, err
	}
	if input.Name != sp.Name {
		if existing, _ := s.store.GetSpeciesByName(input.Name); existing != nil {
			return nil, model.NewValidationError("name", "品种名称已存在")
		}
	}
	sp.Name = input.Name
	sp.Category = input.Category
	sp.Description = input.Description
	if err := s.store.UpdateSpecies(sp); err != nil {
		return nil, err
	}
	return sp, nil
}

func (s *Service) DeleteSpecies(id string) error {
	return s.store.DeleteSpecies(id)
}
