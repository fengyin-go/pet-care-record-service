package model

import (
	"strings"
	"time"
)

const (
	VaccineStatusValid   = "valid"
	VaccineStatusExpired = "expired"
)

type VaccineRecord struct {
	ID           string    `json:"id"`
	PetID        string    `json:"pet_id"`
	VaccineName  string    `json:"vaccine_name"`
	Dose         int       `json:"dose"`
	BatchNumber  string    `json:"batch_number"`
	VaccinatedAt time.Time `json:"vaccinated_at"`
	NextDueAt    time.Time `json:"next_due_at"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

func (v *VaccineRecord) Validate() error {
	v.VaccineName = strings.TrimSpace(v.VaccineName)
	v.BatchNumber = strings.TrimSpace(v.BatchNumber)
	v.Notes = strings.TrimSpace(v.Notes)
	if v.PetID == "" {
		return NewValidationError("pet_id", "宠物不能为空")
	}
	if v.VaccineName == "" {
		return NewValidationError("vaccine_name", "疫苗名称不能为空")
	}
	if v.Dose <= 0 {
		return NewValidationError("dose", "剂次必须大于 0")
	}
	if v.VaccinatedAt.IsZero() {
		return NewValidationError("vaccinated_at", "接种时间不能为空")
	}
	if v.NextDueAt.IsZero() {
		return NewValidationError("next_due_at", "下次到期时间不能为空")
	}
	return nil
}

func (v *VaccineRecord) Status(now time.Time) string {
	if v.NextDueAt.Before(now) {
		return VaccineStatusExpired
	}
	return VaccineStatusValid
}

func (v *VaccineRecord) DaysUntilDue(now time.Time) int {
	diff := v.NextDueAt.Sub(now)
	return int(diff.Hours() / 24)
}

type VaccineRecordFilter struct {
	PetID       string
	VaccineName string
	Status      string
}

func (f VaccineRecordFilter) Match(v *VaccineRecord) bool {
	if f.PetID != "" && v.PetID != f.PetID {
		return false
	}
	if f.VaccineName != "" && v.VaccineName != f.VaccineName {
		return false
	}
	if f.Status != "" {
		s := v.Status(time.Now())
		if s != f.Status {
			return false
		}
	}
	return true
}
