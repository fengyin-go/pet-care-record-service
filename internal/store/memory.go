package store

import (
	"sync"

	"pet-care/internal/model"
)

type MemoryStore struct {
	mu             sync.RWMutex
	owners         map[string]*model.Owner
	species        map[string]*model.Species
	pets           map[string]*model.Pet
	medicalRecords map[string]*model.MedicalRecord
	vaccineRecords map[string]*model.VaccineRecord
	weightRecords  map[string]*model.WeightRecord
	feedingRecords map[string]*model.FeedingRecord
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		owners:         make(map[string]*model.Owner),
		species:        make(map[string]*model.Species),
		pets:           make(map[string]*model.Pet),
		medicalRecords: make(map[string]*model.MedicalRecord),
		vaccineRecords: make(map[string]*model.VaccineRecord),
		weightRecords:  make(map[string]*model.WeightRecord),
		feedingRecords: make(map[string]*model.FeedingRecord),
	}
}

var _ Store = (*MemoryStore)(nil)
