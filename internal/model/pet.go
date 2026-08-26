package model

import (
	"strings"
	"time"
)

const (
	PetStatusActive   = "active"
	PetStatusArchived = "archived"
)

var ValidPetStatuses = []string{PetStatusActive, PetStatusArchived}

var petTransitions = map[string]map[string]bool{
	PetStatusActive:   {PetStatusArchived: true},
	PetStatusArchived: {PetStatusActive: true},
}

func CanTransitionPetStatus(from, to string) bool {
	if m, ok := petTransitions[from]; ok {
		return m[to]
	}
	return false
}

func IsValidPetStatus(status string) bool {
	for _, s := range ValidPetStatuses {
		if status == s {
			return true
		}
	}
	return false
}

type Pet struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	SpeciesID   string    `json:"species_id"`
	Breed       string    `json:"breed"`
	Gender      string    `json:"gender"`
	BirthDate   string    `json:"birth_date"`
	Color       string    `json:"color"`
	MicrochipID string    `json:"microchip_id"`
	Status      string    `json:"status"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *Pet) Validate() error {
	p.Name = strings.TrimSpace(p.Name)
	p.Breed = strings.TrimSpace(p.Breed)
	p.Gender = strings.TrimSpace(p.Gender)
	p.BirthDate = strings.TrimSpace(p.BirthDate)
	p.Color = strings.TrimSpace(p.Color)
	p.MicrochipID = strings.TrimSpace(p.MicrochipID)
	if p.Name == "" {
		return NewValidationError("name", "宠物名称不能为空")
	}
	if p.SpeciesID == "" {
		return NewValidationError("species_id", "品种不能为空")
	}
	if p.OwnerID == "" {
		return NewValidationError("owner_id", "主人不能为空")
	}
	if p.Status == "" {
		p.Status = PetStatusActive
	}
	if !IsValidPetStatus(p.Status) {
		return NewValidationError("status", "宠物状态不合法")
	}
	if p.Gender != "" && p.Gender != "公" && p.Gender != "母" {
		return NewValidationError("gender", "性别必须为公或母")
	}
	return nil
}

type PetFilter struct {
	Status    string
	SpeciesID string
	OwnerID   string
	Keyword   string
	Gender    string
}

func (f PetFilter) Match(p *Pet) bool {
	if f.Status != "" && p.Status != f.Status {
		return false
	}
	if f.SpeciesID != "" && p.SpeciesID != f.SpeciesID {
		return false
	}
	if f.OwnerID != "" && p.OwnerID != f.OwnerID {
		return false
	}
	if f.Gender != "" && p.Gender != f.Gender {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(p.Name), k) &&
			!strings.Contains(strings.ToLower(p.Breed), k) &&
			!strings.Contains(strings.ToLower(p.Color), k) {
			return false
		}
	}
	return true
}
