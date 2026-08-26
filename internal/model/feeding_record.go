package model

import (
	"strings"
	"time"
)

type FeedingRecord struct {
	ID        string    `json:"id"`
	PetID     string    `json:"pet_id"`
	Food      string    `json:"food"`
	Brand     string    `json:"brand"`
	Amount    float64   `json:"amount"`
	Unit      string    `json:"unit"`
	FedAt     time.Time `json:"fed_at"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

func (f *FeedingRecord) Validate() error {
	f.Food = strings.TrimSpace(f.Food)
	f.Brand = strings.TrimSpace(f.Brand)
	f.Notes = strings.TrimSpace(f.Notes)
	if f.PetID == "" {
		return NewValidationError("pet_id", "宠物不能为空")
	}
	if f.Food == "" {
		return NewValidationError("food", "食物名称不能为空")
	}
	if f.Amount <= 0 {
		return NewValidationError("amount", "食量必须大于 0")
	}
	if f.Unit == "" {
		f.Unit = "g"
	}
	if f.FedAt.IsZero() {
		return NewValidationError("fed_at", "喂养时间不能为空")
	}
	return nil
}

type FeedingRecordFilter struct {
	PetID string
	Food  string
}

func (f FeedingRecordFilter) Match(r *FeedingRecord) bool {
	if f.PetID != "" && r.PetID != f.PetID {
		return false
	}
	if f.Food != "" && r.Food != f.Food {
		return false
	}
	return true
}
