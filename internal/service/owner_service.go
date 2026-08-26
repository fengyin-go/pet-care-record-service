package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/pkg/idgen"
)

func (s *Service) CreateOwner(input model.Owner) (*model.Owner, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if existing, _ := s.store.GetOwnerByEmail(input.Email); existing != nil {
		return nil, model.NewValidationError("email", "邮箱已被注册")
	}
	input.ID = idgen.Hex()
	input.CreatedAt = time.Now()
	if err := s.store.CreateOwner(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetOwner(id string) (*model.Owner, error) {
	return s.store.GetOwner(id)
}

func (s *Service) GetOwnerByEmail(email string) (*model.Owner, error) {
	return s.store.GetOwnerByEmail(email)
}

func (s *Service) ListOwners(filter model.OwnerFilter, page, size int) ([]*model.Owner, int, error) {
	all := s.store.ListOwners()
	matched := make([]*model.Owner, 0, len(all))
	for _, o := range all {
		if filter.Match(o) {
			matched = append(matched, o)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Owner{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateOwner(id string, input model.Owner) (*model.Owner, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	o, err := s.store.GetOwner(id)
	if err != nil {
		return nil, err
	}
	if input.Email != o.Email {
		if existing, _ := s.store.GetOwnerByEmail(input.Email); existing != nil {
			return nil, model.NewValidationError("email", "邮箱已被注册")
		}
	}
	o.Name = input.Name
	o.Phone = input.Phone
	o.Email = input.Email
	if err := s.store.UpdateOwner(o); err != nil {
		return nil, err
	}
	return o, nil
}

func (s *Service) DeleteOwner(id string) error {
	return s.store.DeleteOwner(id)
}
