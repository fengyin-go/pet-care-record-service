package service

import (
	"pet-care/internal/model"
)

type GlobalStats struct {
	OwnerCount         int `json:"owner_count"`
	PetCount           int `json:"pet_count"`
	MedicalRecordCount int `json:"medical_record_count"`
	VaccineRecordCount int `json:"vaccine_record_count"`
	WeightRecordCount  int `json:"weight_record_count"`
	FeedingRecordCount int `json:"feeding_record_count"`
}

func (s *Service) GlobalStats() (*GlobalStats, error) {
	stats := &GlobalStats{
		OwnerCount:         len(s.store.ListOwners()),
		PetCount:           len(s.store.ListPets()),
		MedicalRecordCount: len(s.store.ListMedicalRecords()),
		VaccineRecordCount: len(s.store.ListVaccineRecords()),
		WeightRecordCount:  len(s.store.ListWeightRecords()),
		FeedingRecordCount: len(s.store.ListFeedingRecords()),
	}
	return stats, nil
}

type PetRecordStats struct {
	MedicalRecordCount int `json:"medical_record_count"`
	VaccineRecordCount int `json:"vaccine_record_count"`
	WeightRecordCount  int `json:"weight_record_count"`
	FeedingRecordCount int `json:"feeding_record_count"`
}

func (s *Service) PetRecordStats(petID string) (*PetRecordStats, error) {
	if _, err := s.store.GetPet(petID); err != nil {
		return nil, err
	}
	stats := &PetRecordStats{}
	for _, m := range s.store.ListMedicalRecords() {
		if m.PetID == petID {
			stats.MedicalRecordCount++
		}
	}
	for _, v := range s.store.ListVaccineRecords() {
		if v.PetID == petID {
			stats.VaccineRecordCount++
		}
	}
	for _, w := range s.store.ListWeightRecords() {
		if w.PetID == petID {
			stats.WeightRecordCount++
		}
	}
	for _, f := range s.store.ListFeedingRecords() {
		if f.PetID == petID {
			stats.FeedingRecordCount++
		}
	}
	return stats, nil
}

type SpeciesPetCount struct {
	SpeciesID   string `json:"species_id"`
	SpeciesName string `json:"species_name"`
	Count       int    `json:"count"`
}

func (s *Service) CountPetsBySpecies() ([]SpeciesPetCount, error) {
	pets := s.store.ListPets()
	m := make(map[string]int)
	for _, p := range pets {
		m[p.SpeciesID]++
	}
	result := make([]SpeciesPetCount, 0, len(m))
	for speciesID, count := range m {
		name := ""
		if sp, err := s.store.GetSpecies(speciesID); err == nil {
			name = sp.Name
		}
		result = append(result, SpeciesPetCount{SpeciesID: speciesID, SpeciesName: name, Count: count})
	}
	return result, nil
}

type PetHealthProfile struct {
	Pet            *model.Pet             `json:"pet"`
	Owner          *model.Owner           `json:"owner"`
	Species        *model.Species         `json:"species"`
	MedicalRecords []*model.MedicalRecord `json:"medical_records"`
	VaccineRecords []*model.VaccineRecord `json:"vaccine_records"`
	WeightRecords  []*model.WeightRecord  `json:"weight_records"`
	FeedingRecords []*model.FeedingRecord `json:"feeding_records"`
}

func (s *Service) ExportPetHealthProfile(petID string) (*PetHealthProfile, error) {
	pet, err := s.store.GetPet(petID)
	if err != nil {
		return nil, err
	}
	profile := &PetHealthProfile{Pet: pet}
	if owner, err := s.store.GetOwner(pet.OwnerID); err == nil {
		profile.Owner = owner
	}
	if sp, err := s.store.GetSpecies(pet.SpeciesID); err == nil {
		profile.Species = sp
	}
	for _, m := range s.store.ListMedicalRecords() {
		if m.PetID == petID {
			profile.MedicalRecords = append(profile.MedicalRecords, m)
		}
	}
	for _, v := range s.store.ListVaccineRecords() {
		if v.PetID == petID {
			profile.VaccineRecords = append(profile.VaccineRecords, v)
		}
	}
	for _, w := range s.store.ListWeightRecords() {
		if w.PetID == petID {
			profile.WeightRecords = append(profile.WeightRecords, w)
		}
	}
	for _, f := range s.store.ListFeedingRecords() {
		if f.PetID == petID {
			profile.FeedingRecords = append(profile.FeedingRecords, f)
		}
	}
	return profile, nil
}
