package service

import (
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/internal/store"
	"pet-care/pkg/idgen"
)

func (s *Service) CreatePet(input model.Pet) (*model.Pet, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetSpecies(input.SpeciesID); err != nil {
		if err == store.ErrNotFound {
			return nil, model.NewValidationError("species_id", "品种不存在")
		}
		return nil, err
	}
	if _, err := s.store.GetOwner(input.OwnerID); err != nil {
		if err == store.ErrNotFound {
			return nil, model.NewValidationError("owner_id", "主人不存在")
		}
		return nil, err
	}
	input.ID = idgen.Hex()
	input.CreatedAt = time.Now()
	if err := s.store.CreatePet(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetPet(id string) (*model.Pet, error) {
	return s.store.GetPet(id)
}

func (s *Service) ListPets(filter model.PetFilter, page, size int) ([]*model.Pet, int, error) {
	all := s.store.ListPets()
	matched := make([]*model.Pet, 0, len(all))
	for _, p := range all {
		if filter.Match(p) {
			matched = append(matched, p)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Pet{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdatePet(id string, input model.Pet) (*model.Pet, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	p, err := s.store.GetPet(id)
	if err != nil {
		return nil, err
	}
	if input.SpeciesID != p.SpeciesID {
		if _, err := s.store.GetSpecies(input.SpeciesID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("species_id", "品种不存在")
			}
			return nil, err
		}
	}
	if input.OwnerID != p.OwnerID {
		if _, err := s.store.GetOwner(input.OwnerID); err != nil {
			if err == store.ErrNotFound {
				return nil, model.NewValidationError("owner_id", "主人不存在")
			}
			return nil, err
		}
	}
	p.Name = input.Name
	p.SpeciesID = input.SpeciesID
	p.Breed = input.Breed
	p.Gender = input.Gender
	p.BirthDate = input.BirthDate
	p.OwnerID = input.OwnerID
	if err := s.store.UpdatePet(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) TransitionPetStatus(id string, to string) (*model.Pet, error) {
	p, err := s.store.GetPet(id)
	if err != nil {
		return nil, err
	}
	if !model.CanTransitionPetStatus(p.Status, to) {
		return nil, model.NewValidationError("status", "状态流转不合法")
	}
	p.Status = to
	if err := s.store.UpdatePet(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) DeletePet(id string) error {
	return s.store.DeletePet(id)
}

func (s *Service) DeletePetWithRecords(id string) error {
	if err := s.store.DeleteMedicalRecordsByPet(id); err != nil {
		return err
	}
	if err := s.store.DeleteVaccineRecordsByPet(id); err != nil {
		return err
	}
	if err := s.store.DeleteWeightRecordsByPet(id); err != nil {
		return err
	}
	if err := s.store.DeleteFeedingRecordsByPet(id); err != nil {
		return err
	}
	return s.store.DeletePet(id)
}
