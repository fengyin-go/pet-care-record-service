package model

import (
	"strings"
	"time"
)

const (
	MedicalTypeRoutine     = "routine"
	MedicalTypeEmergency   = "emergency"
	MedicalTypeSurgery     = "surgery"
	MedicalTypeVaccination = "vaccination"
)

type MedicalRecord struct {
	ID        string    `json:"id"`
	PetID     string    `json:"pet_id"`
	VetName   string    `json:"vet_name"`
	Diagnosis string    `json:"diagnosis"`
	Treatment string    `json:"treatment"`
	VisitDate string    `json:"visit_date"`
	Type      string    `json:"type"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *MedicalRecord) Validate() error {
	m.VetName = strings.TrimSpace(m.VetName)
	m.Diagnosis = strings.TrimSpace(m.Diagnosis)
	m.Treatment = strings.TrimSpace(m.Treatment)
	m.VisitDate = strings.TrimSpace(m.VisitDate)
	m.Notes = strings.TrimSpace(m.Notes)
	m.Type = strings.TrimSpace(m.Type)
	if m.PetID == "" {
		return NewValidationError("pet_id", "宠物不能为空")
	}
	if m.VetName == "" {
		return NewValidationError("vet_name", "兽医姓名不能为空")
	}
	if m.Diagnosis == "" {
		return NewValidationError("diagnosis", "诊断不能为空")
	}
	if m.VisitDate == "" {
		return NewValidationError("visit_date", "就诊日期不能为空")
	}
	if m.Type == "" {
		m.Type = MedicalTypeRoutine
	}
	return nil
}

type MedicalRecordFilter struct {
	PetID     string
	VetName   string
	Keyword   string
	VisitDate string
	Type      string
}

func (f MedicalRecordFilter) Match(m *MedicalRecord) bool {
	if f.PetID != "" && m.PetID != f.PetID {
		return false
	}
	if f.VetName != "" && m.VetName != f.VetName {
		return false
	}
	if f.VisitDate != "" && m.VisitDate != f.VisitDate {
		return false
	}
	if f.Type != "" && m.Type != f.Type {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(m.Diagnosis), k) &&
			!strings.Contains(strings.ToLower(m.Treatment), k) &&
			!strings.Contains(strings.ToLower(m.VetName), k) {
			return false
		}
	}
	return true
}
