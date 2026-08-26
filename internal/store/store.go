// Package store 定义数据访问接口与内存实现。
package store

import (
	"errors"

	"pet-care/internal/model"
)

var (
	ErrNotFound = errors.New("记录不存在")
	ErrConflict = errors.New("记录已存在或状态冲突")
)

type Store interface {
	// Owner
	CreateOwner(o *model.Owner) error
	GetOwner(id string) (*model.Owner, error)
	GetOwnerByEmail(email string) (*model.Owner, error)
	ListOwners() []*model.Owner
	UpdateOwner(o *model.Owner) error
	DeleteOwner(id string) error

	// Species
	CreateSpecies(s *model.Species) error
	GetSpecies(id string) (*model.Species, error)
	GetSpeciesByName(name string) (*model.Species, error)
	ListSpecies() []*model.Species
	UpdateSpecies(s *model.Species) error
	DeleteSpecies(id string) error

	// Pet
	CreatePet(p *model.Pet) error
	GetPet(id string) (*model.Pet, error)
	ListPets() []*model.Pet
	UpdatePet(p *model.Pet) error
	DeletePet(id string) error

	// MedicalRecord
	CreateMedicalRecord(m *model.MedicalRecord) error
	GetMedicalRecord(id string) (*model.MedicalRecord, error)
	ListMedicalRecords() []*model.MedicalRecord
	UpdateMedicalRecord(m *model.MedicalRecord) error
	DeleteMedicalRecord(id string) error
	DeleteMedicalRecordsByPet(petID string) error

	// VaccineRecord
	CreateVaccineRecord(v *model.VaccineRecord) error
	GetVaccineRecord(id string) (*model.VaccineRecord, error)
	ListVaccineRecords() []*model.VaccineRecord
	UpdateVaccineRecord(v *model.VaccineRecord) error
	DeleteVaccineRecord(id string) error
	DeleteVaccineRecordsByPet(petID string) error

	// WeightRecord
	CreateWeightRecord(w *model.WeightRecord) error
	GetWeightRecord(id string) (*model.WeightRecord, error)
	ListWeightRecords() []*model.WeightRecord
	UpdateWeightRecord(w *model.WeightRecord) error
	DeleteWeightRecord(id string) error
	DeleteWeightRecordsByPet(petID string) error

	// FeedingRecord
	CreateFeedingRecord(f *model.FeedingRecord) error
	GetFeedingRecord(id string) (*model.FeedingRecord, error)
	ListFeedingRecords() []*model.FeedingRecord
	UpdateFeedingRecord(f *model.FeedingRecord) error
	DeleteFeedingRecord(id string) error
	DeleteFeedingRecordsByPet(petID string) error
}
