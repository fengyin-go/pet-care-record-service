package model

import (
	"strings"
	"time"
)

const (
	SpeciesCategoryDog   = "犬"
	SpeciesCategoryCat   = "猫"
	SpeciesCategoryOther = "其他"
)

var ValidSpeciesCategories = []string{SpeciesCategoryDog, SpeciesCategoryCat, SpeciesCategoryOther}

type Species struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Origin      string    `json:"origin"`
	LifeSpan    string    `json:"life_span"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s *Species) Validate() error {
	s.Name = strings.TrimSpace(s.Name)
	s.Category = strings.TrimSpace(s.Category)
	s.Description = strings.TrimSpace(s.Description)
	s.Origin = strings.TrimSpace(s.Origin)
	s.LifeSpan = strings.TrimSpace(s.LifeSpan)
	if s.Name == "" {
		return NewValidationError("name", "品种名称不能为空")
	}
	if s.Category == "" {
		s.Category = SpeciesCategoryOther
	}
	valid := false
	for _, c := range ValidSpeciesCategories {
		if s.Category == c {
			valid = true
			break
		}
	}
	if !valid {
		return NewValidationError("category", "品种分类不合法，可选：犬/猫/其他")
	}
	return nil
}

type SpeciesFilter struct {
	Category string
	Name     string
	Keyword  string
}

func (f SpeciesFilter) Match(s *Species) bool {
	if f.Category != "" && s.Category != f.Category {
		return false
	}
	if f.Name != "" && s.Name != f.Name {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(s.Name), k) &&
			!strings.Contains(strings.ToLower(s.Description), k) {
			return false
		}
	}
	return true
}
