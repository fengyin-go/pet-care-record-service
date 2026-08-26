// Package report 提供各类统计报表生成能力。
package report

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"pet-care/internal/model"
	"pet-care/internal/store"
)

// ReportGenerator 报表生成器。
type ReportGenerator struct {
	store store.Store
}

// NewReportGenerator 创建报表生成器。
func NewReportGenerator(st store.Store) *ReportGenerator {
	return &ReportGenerator{store: st}
}

// OwnerSummary 主人汇总信息。
type OwnerSummary struct {
	OwnerID      string `json:"owner_id"`
	OwnerName    string `json:"owner_name"`
	PetCount     int    `json:"pet_count"`
	ActivePets   int    `json:"active_pets"`
	ArchivedPets int    `json:"archived_pets"`
}

// GenerateOwnerSummary 生成主人宠物汇总报表。
func (rg *ReportGenerator) GenerateOwnerSummary() ([]OwnerSummary, error) {
	owners := rg.store.ListOwners()
	pets := rg.store.ListPets()
	m := make(map[string]*OwnerSummary)
	for _, o := range owners {
		m[o.ID] = &OwnerSummary{
			OwnerID:   o.ID,
			OwnerName: o.Name,
		}
	}
	for _, p := range pets {
		if s, ok := m[p.OwnerID]; ok {
			s.PetCount++
			if p.Status == model.PetStatusActive {
				s.ActivePets++
			} else {
				s.ArchivedPets++
			}
		}
	}
	result := make([]OwnerSummary, 0, len(m))
	for _, s := range m {
		result = append(result, *s)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].PetCount > result[j].PetCount
	})
	return result, nil
}

// SpeciesSummary 品种汇总信息。
type SpeciesSummary struct {
	SpeciesID   string `json:"species_id"`
	SpeciesName string `json:"species_name"`
	Category    string `json:"category"`
	PetCount    int    `json:"pet_count"`
}

// GenerateSpeciesSummary 生成品种宠物汇总报表。
func (rg *ReportGenerator) GenerateSpeciesSummary() ([]SpeciesSummary, error) {
	species := rg.store.ListSpecies()
	pets := rg.store.ListPets()
	m := make(map[string]*SpeciesSummary)
	for _, sp := range species {
		m[sp.ID] = &SpeciesSummary{
			SpeciesID:   sp.ID,
			SpeciesName: sp.Name,
			Category:    sp.Category,
		}
	}
	for _, p := range pets {
		if s, ok := m[p.SpeciesID]; ok {
			s.PetCount++
		}
	}
	result := make([]SpeciesSummary, 0, len(m))
	for _, s := range m {
		result = append(result, *s)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].PetCount > result[j].PetCount
	})
	return result, nil
}

// MonthlyRecordSummary 月度记录汇总。
type MonthlyRecordSummary struct {
	YearMonth    string `json:"year_month"`
	MedicalCount int    `json:"medical_count"`
	VaccineCount int    `json:"vaccine_count"`
	WeightCount  int    `json:"weight_count"`
	FeedingCount int    `json:"feeding_count"`
	TotalRecords int    `json:"total_records"`
}

// GenerateMonthlySummary 生成月度记录汇总报表。
func (rg *ReportGenerator) GenerateMonthlySummary() ([]MonthlyRecordSummary, error) {
	m := make(map[string]*MonthlyRecordSummary)
	for _, rec := range rg.store.ListMedicalRecords() {
		key := rec.CreatedAt.Format("2006-01")
		if _, ok := m[key]; !ok {
			m[key] = &MonthlyRecordSummary{YearMonth: key}
		}
		m[key].MedicalCount++
		m[key].TotalRecords++
	}
	for _, rec := range rg.store.ListVaccineRecords() {
		key := rec.CreatedAt.Format("2006-01")
		if _, ok := m[key]; !ok {
			m[key] = &MonthlyRecordSummary{YearMonth: key}
		}
		m[key].VaccineCount++
		m[key].TotalRecords++
	}
	for _, rec := range rg.store.ListWeightRecords() {
		key := rec.CreatedAt.Format("2006-01")
		if _, ok := m[key]; !ok {
			m[key] = &MonthlyRecordSummary{YearMonth: key}
		}
		m[key].WeightCount++
		m[key].TotalRecords++
	}
	for _, rec := range rg.store.ListFeedingRecords() {
		key := rec.CreatedAt.Format("2006-01")
		if _, ok := m[key]; !ok {
			m[key] = &MonthlyRecordSummary{YearMonth: key}
		}
		m[key].FeedingCount++
		m[key].TotalRecords++
	}
	result := make([]MonthlyRecordSummary, 0, len(m))
	for _, s := range m {
		result = append(result, *s)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].YearMonth > result[j].YearMonth
	})
	return result, nil
}

// VaccineAlert 疫苗预警信息。
type VaccineAlert struct {
	PetID        string    `json:"pet_id"`
	PetName      string    `json:"pet_name"`
	VaccineName  string    `json:"vaccine_name"`
	NextDueAt    time.Time `json:"next_due_at"`
	DaysUntilDue int       `json:"days_until_due"`
	Status       string    `json:"status"`
}

// GenerateVaccineAlerts 生成疫苗预警报表。
func (rg *ReportGenerator) GenerateVaccineAlerts(now time.Time) ([]VaccineAlert, error) {
	records := rg.store.ListVaccineRecords()
	pets := rg.store.ListPets()
	petMap := make(map[string]*model.Pet)
	for _, p := range pets {
		petMap[p.ID] = p
	}
	result := make([]VaccineAlert, 0)
	for _, v := range records {
		status := v.Status(now)
		days := v.DaysUntilDue(now)
		petName := ""
		if p, ok := petMap[v.PetID]; ok {
			petName = p.Name
		}
		result = append(result, VaccineAlert{
			PetID:        v.PetID,
			PetName:      petName,
			VaccineName:  v.VaccineName,
			NextDueAt:    v.NextDueAt,
			DaysUntilDue: days,
			Status:       status,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].DaysUntilDue < result[j].DaysUntilDue
	})
	return result, nil
}

// WeightChangeReport 体重变化报告。
type WeightChangeReport struct {
	PetID         string  `json:"pet_id"`
	PetName       string  `json:"pet_name"`
	FirstWeight   float64 `json:"first_weight"`
	LatestWeight  float64 `json:"latest_weight"`
	ChangeAmount  float64 `json:"change_amount"`
	ChangePercent float64 `json:"change_percent"`
	RecordCount   int     `json:"record_count"`
}

// GenerateWeightChangeReport 生成体重变化报表。
func (rg *ReportGenerator) GenerateWeightChangeReport() ([]WeightChangeReport, error) {
	pets := rg.store.ListPets()
	allWeights := rg.store.ListWeightRecords()
	result := make([]WeightChangeReport, 0)
	for _, pet := range pets {
		petWeights := make([]*model.WeightRecord, 0)
		for _, w := range allWeights {
			if w.PetID == pet.ID {
				petWeights = append(petWeights, w)
			}
		}
		if len(petWeights) < 2 {
			continue
		}
		sort.Slice(petWeights, func(i, j int) bool {
			return petWeights[i].MeasuredAt.Before(petWeights[j].MeasuredAt)
		})
		first := petWeights[0].WeightInKG()
		latest := petWeights[len(petWeights)-1].WeightInKG()
		change := latest - first
		percent := 0.0
		if first != 0 {
			percent = (change / first) * 100
		}
		result = append(result, WeightChangeReport{
			PetID:         pet.ID,
			PetName:       pet.Name,
			FirstWeight:   first,
			LatestWeight:  latest,
			ChangeAmount:  change,
			ChangePercent: percent,
			RecordCount:   len(petWeights),
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ChangeAmount > result[j].ChangeAmount
	})
	return result, nil
}

// DashboardStats 仪表盘统计。
type DashboardStats struct {
	TotalOwners        int `json:"total_owners"`
	TotalPets          int `json:"total_pets"`
	ActivePets         int `json:"active_pets"`
	ArchivedPets       int `json:"archived_pets"`
	TotalMedical       int `json:"total_medical"`
	TotalVaccine       int `json:"total_vaccine"`
	TotalWeight        int `json:"total_weight"`
	TotalFeeding       int `json:"total_feeding"`
	ExpiredVaccines    int `json:"expired_vaccines"`
	ExpiringVaccines30 int `json:"expiring_vaccines_30"`
}

// GenerateDashboardStats 生成仪表盘统计。
func (rg *ReportGenerator) GenerateDashboardStats(now time.Time) (*DashboardStats, error) {
	stats := &DashboardStats{
		TotalOwners:  len(rg.store.ListOwners()),
		TotalPets:    len(rg.store.ListPets()),
		TotalMedical: len(rg.store.ListMedicalRecords()),
		TotalVaccine: len(rg.store.ListVaccineRecords()),
		TotalWeight:  len(rg.store.ListWeightRecords()),
		TotalFeeding: len(rg.store.ListFeedingRecords()),
	}
	for _, p := range rg.store.ListPets() {
		if p.Status == model.PetStatusActive {
			stats.ActivePets++
		} else {
			stats.ArchivedPets++
		}
	}
	for _, v := range rg.store.ListVaccineRecords() {
		if v.Status(now) == model.VaccineStatusExpired {
			stats.ExpiredVaccines++
		} else if v.NextDueAt.Before(now.AddDate(0, 0, 30)) {
			stats.ExpiringVaccines30++
		}
	}
	return stats, nil
}

// FeedingFrequencyReport 喂养频率报告。
type FeedingFrequencyReport struct {
	PetID       string  `json:"pet_id"`
	PetName     string  `json:"pet_name"`
	FoodType    string  `json:"food_type"`
	TotalAmount float64 `json:"total_amount"`
	AvgAmount   float64 `json:"avg_amount"`
	Count       int     `json:"count"`
}

// GenerateFeedingFrequencyReport 生成喂养频率报表。
func (rg *ReportGenerator) GenerateFeedingFrequencyReport() ([]FeedingFrequencyReport, error) {
	pets := rg.store.ListPets()
	allFeedings := rg.store.ListFeedingRecords()
	petMap := make(map[string]string)
	for _, p := range pets {
		petMap[p.ID] = p.Name
	}
	type key struct {
		PetID    string
		FoodType string
	}
	m := make(map[key]*FeedingFrequencyReport)
	for _, f := range allFeedings {
		k := key{PetID: f.PetID, FoodType: f.Food}
		if _, ok := m[k]; !ok {
			m[k] = &FeedingFrequencyReport{
				PetID:    f.PetID,
				PetName:  petMap[f.PetID],
				FoodType: f.Food,
			}
		}
		m[k].TotalAmount += f.Amount
		m[k].Count++
	}
	result := make([]FeedingFrequencyReport, 0, len(m))
	for _, r := range m {
		if r.Count > 0 {
			r.AvgAmount = r.TotalAmount / float64(r.Count)
		}
		result = append(result, *r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})
	return result, nil
}

// ExportFormat 导出格式枚举。
type ExportFormat string

const (
	ExportFormatJSON ExportFormat = "json"
	ExportFormatCSV  ExportFormat = "csv"
)

// ExportReport 导出报表数据（JSON 格式）。
func (rg *ReportGenerator) ExportReport(data interface{}, format ExportFormat) ([]byte, error) {
	if format == ExportFormatJSON {
		return json.MarshalIndent(data, "", "  ")
	}
	return nil, fmt.Errorf("不支持的导出格式: %s", format)
}
