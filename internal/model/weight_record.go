package model

import (
	"time"
)

type WeightRecord struct {
	ID         string    `json:"id"`
	PetID      string    `json:"pet_id"`
	Weight     float64   `json:"weight"`
	Unit       string    `json:"unit"`
	MeasuredAt time.Time `json:"measured_at"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
}

func (w *WeightRecord) Validate() error {
	if w.PetID == "" {
		return NewValidationError("pet_id", "宠物不能为空")
	}
	if w.Weight <= 0 {
		return NewValidationError("weight", "体重必须大于 0")
	}
	if w.Unit == "" {
		w.Unit = "kg"
	}
	if w.MeasuredAt.IsZero() {
		return NewValidationError("measured_at", "测量时间不能为空")
	}
	return nil
}

func (w *WeightRecord) WeightInKG() float64 {
	if w.Unit == "g" {
		return w.Weight / 1000
	}
	if w.Unit == "lb" {
		return w.Weight * 0.453592
	}
	return w.Weight
}

type WeightRecordFilter struct {
	PetID string
}

func (f WeightRecordFilter) Match(w *WeightRecord) bool {
	if f.PetID != "" && w.PetID != f.PetID {
		return false
	}
	return true
}
