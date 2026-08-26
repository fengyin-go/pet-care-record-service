// Package converter 提供单位转换和数据格式化工具。
package converter

import (
	"fmt"
	"strconv"
	"strings"
)

// WeightUnit 重量单位。
type WeightUnit string

const (
	WeightUnitKG WeightUnit = "kg"
	WeightUnitG  WeightUnit = "g"
	WeightUnitLB WeightUnit = "lb"
)

// ConvertWeight 在不同重量单位之间转换。
func ConvertWeight(value float64, from, to WeightUnit) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("重量不能为负数")
	}
	var kg float64
	switch from {
	case WeightUnitKG:
		kg = value
	case WeightUnitG:
		kg = value / 1000
	case WeightUnitLB:
		kg = value * 0.45359237
	default:
		return 0, fmt.Errorf("不支持的源单位: %s", from)
	}
	switch to {
	case WeightUnitKG:
		return kg, nil
	case WeightUnitG:
		return kg * 1000, nil
	case WeightUnitLB:
		return kg / 0.45359237, nil
	default:
		return 0, fmt.Errorf("不支持的目标单位: %s", to)
	}
}

// FormatWeight 格式化重量输出。
func FormatWeight(value float64, unit WeightUnit) string {
	return fmt.Sprintf("%.2f %s", value, unit)
}

// FoodUnit 食物单位。
type FoodUnit string

const (
	FoodUnitG     FoodUnit = "g"
	FoodUnitKG    FoodUnit = "kg"
	FoodUnitCup   FoodUnit = "cup"
	FoodUnitPiece FoodUnit = "piece"
)

// ConvertFoodAmount 食物量转换（cup 按 100g 估算）。
func ConvertFoodAmount(value float64, from, to FoodUnit) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("食量不能为负数")
	}
	var grams float64
	switch from {
	case FoodUnitG:
		grams = value
	case FoodUnitKG:
		grams = value * 1000
	case FoodUnitCup:
		grams = value * 100
	case FoodUnitPiece:
		grams = value * 50
	default:
		return 0, fmt.Errorf("不支持的源单位: %s", from)
	}
	switch to {
	case FoodUnitG:
		return grams, nil
	case FoodUnitKG:
		return grams / 1000, nil
	case FoodUnitCup:
		return grams / 100, nil
	case FoodUnitPiece:
		return grams / 50, nil
	default:
		return 0, fmt.Errorf("不支持的目标单位: %s", to)
	}
}

// AgeFormatter 年龄格式化工具。
type AgeFormatter struct{}

// NewAgeFormatter 创建年龄格式化器。
func NewAgeFormatter() *AgeFormatter {
	return &AgeFormatter{}
}

// FormatMonths 将月龄格式化为人类可读字符串。
func (af *AgeFormatter) FormatMonths(totalMonths int) string {
	if totalMonths < 0 {
		return "未知"
	}
	years := totalMonths / 12
	months := totalMonths % 12
	if years > 0 && months > 0 {
		return fmt.Sprintf("%d岁%d个月", years, months)
	}
	if years > 0 {
		return fmt.Sprintf("%d岁", years)
	}
	return fmt.Sprintf("%d个月", months)
}

// FormatDays 将天数格式化为人类可读字符串。
func (af *AgeFormatter) FormatDays(days int) string {
	if days < 0 {
		return "未知"
	}
	if days < 30 {
		return fmt.Sprintf("%d天", days)
	}
	if days < 365 {
		months := days / 30
		return fmt.Sprintf("%d个月", months)
	}
	years := days / 365
	remainingDays := days % 365
	months := remainingDays / 30
	if months > 0 {
		return fmt.Sprintf("%d岁%d个月", years, months)
	}
	return fmt.Sprintf("%d岁", years)
}

// TemperatureConverter 温度转换器。
type TemperatureConverter struct{}

// CelsiusToFahrenheit 摄氏度转华氏度。
func (tc *TemperatureConverter) CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// FahrenheitToCelsius 华氏度转摄氏度。
func (tc *TemperatureConverter) FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// FormatTemperature 格式化温度输出。
func (tc *TemperatureConverter) FormatTemperature(value float64, unit string) string {
	return fmt.Sprintf("%.1f°%s", value, strings.ToUpper(unit))
}

// DurationConverter 时长转换器。
type DurationConverter struct{}

// MinutesToHoursMinutes 将分钟转为小时分钟格式。
func (dc *DurationConverter) MinutesToHoursMinutes(minutes int) string {
	if minutes < 0 {
		return "0分钟"
	}
	h := minutes / 60
	m := minutes % 60
	if h > 0 && m > 0 {
		return fmt.Sprintf("%d小时%d分钟", h, m)
	}
	if h > 0 {
		return fmt.Sprintf("%d小时", h)
	}
	return fmt.Sprintf("%d分钟", m)
}

// ParseDurationString 解析时长字符串（如 "1小时30分钟"）。
func (dc *DurationConverter) ParseDurationString(s string) (minutes int, err error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	parts := strings.Fields(s)
	for i := 0; i < len(parts); i += 2 {
		if i+1 >= len(parts) {
			break
		}
		val, err := strconv.Atoi(parts[i])
		if err != nil {
			return 0, fmt.Errorf("无法解析数值: %s", parts[i])
		}
		unit := parts[i+1]
		switch {
		case strings.Contains(unit, "小时") || unit == "h":
			minutes += val * 60
		case strings.Contains(unit, "分钟") || unit == "m":
			minutes += val
		case strings.Contains(unit, "天") || unit == "d":
			minutes += val * 24 * 60
		}
	}
	return minutes, nil
}

// StringFormatter 字符串格式化工具。
type StringFormatter struct{}

// NewStringFormatter 创建字符串格式化器。
func NewStringFormatter() *StringFormatter {
	return &StringFormatter{}
}

// PadLeft 左填充字符串。
func (sf *StringFormatter) PadLeft(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(pad), length-len(s))
	return padding + s
}

// PadRight 右填充字符串。
func (sf *StringFormatter) PadRight(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(pad), length-len(s))
	return s + padding
}

// TruncateWithEllipsis 截断字符串并添加省略号。
func (sf *StringFormatter) TruncateWithEllipsis(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// NumberFormatter 数字格式化工具。
type NumberFormatter struct{}

// FormatWithCommas 千分位格式化整数。
func (nf *NumberFormatter) FormatWithCommas(n int64) string {
	s := strconv.FormatInt(n, 10)
	if n < 0 {
		s = s[1:]
	}
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	if n < 0 {
		result = "-" + result
	}
	return result
}

// RoundToDecimals 四舍五入到指定小数位。
func (nf *NumberFormatter) RoundToDecimals(value float64, decimals int) float64 {
	format := fmt.Sprintf("%%.%df", decimals)
	s := fmt.Sprintf(format, value)
	result, _ := strconv.ParseFloat(s, 64)
	return result
}
