// Package export 提供数据导出功能。
package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// JSONExporter JSON 导出器。
type JSONExporter struct {
	indent bool
}

// NewJSONExporter 创建新的 JSON 导出器。
func NewJSONExporter(indent bool) *JSONExporter {
	return &JSONExporter{indent: indent}
}

// Export 将数据导出为 JSON 字节数组。
func (e *JSONExporter) Export(data interface{}) ([]byte, error) {
	if e.indent {
		return json.MarshalIndent(data, "", "  ")
	}
	return json.Marshal(data)
}

// ExportToString 将数据导出为 JSON 字符串。
func (e *JSONExporter) ExportToString(data interface{}) (string, error) {
	b, err := e.Export(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ExportWithHeader 导出数据并添加 HTTP 响应头信息。
func (e *JSONExporter) ExportWithHeader(data interface{}, filename string) ([]byte, map[string]string, error) {
	b, err := e.Export(data)
	if err != nil {
		return nil, nil, err
	}
	headers := map[string]string{
		"Content-Type":        "application/json; charset=utf-8",
		"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
	}
	return b, headers, nil
}

// HealthProfileSummary 健康档案汇总结构。
type HealthProfileSummary struct {
	PetID           string    `json:"pet_id"`
	PetName         string    `json:"pet_name"`
	OwnerName       string    `json:"owner_name"`
	SpeciesName     string    `json:"species_name"`
	TotalMedical    int       `json:"total_medical"`
	TotalVaccine    int       `json:"total_vaccine"`
	TotalWeight     int       `json:"total_weight"`
	TotalFeeding    int       `json:"total_feeding"`
	LatestWeight    float64   `json:"latest_weight"`
	VaccinesExpired int       `json:"vaccines_expired"`
	VaccinesValid   int       `json:"vaccines_valid"`
	GeneratedAt     time.Time `json:"generated_at"`
}

// SummarizeHealthProfile 汇总健康档案信息。
func SummarizeHealthProfile(petID, petName, ownerName, speciesName string, medical, vaccine, weight, feeding int, latestWeight float64, expired, valid int) *HealthProfileSummary {
	return &HealthProfileSummary{
		PetID:           petID,
		PetName:         petName,
		OwnerName:       ownerName,
		SpeciesName:     speciesName,
		TotalMedical:    medical,
		TotalVaccine:    vaccine,
		TotalWeight:     weight,
		TotalFeeding:    feeding,
		LatestWeight:    latestWeight,
		VaccinesExpired: expired,
		VaccinesValid:   valid,
		GeneratedAt:     time.Now(),
	}
}

// ExportBatch 批量导出多个记录。
func ExportBatch(items []interface{}) ([]byte, error) {
	return json.Marshal(items)
}

// FormatBytes 格式化字节数为人类可读字符串。
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// WriteJSONBuffer 将数据写入 bytes.Buffer。
func WriteJSONBuffer(buf *bytes.Buffer, data interface{}) error {
	enc := json.NewEncoder(buf)
	return enc.Encode(data)
}
