// Package utils 提供通用工具函数。
package utils

import (
	"strings"
	"time"
	"unicode"
)

// ToPointer 返回值的指针。
func ToPointer[T any](v T) *T {
	return &v
}

// Coalesce 返回第一个非零值。
func Coalesce[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// FormatDate 格式化日期为 YYYY-MM-DD。
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseDate 解析 YYYY-MM-DD 格式日期。
func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// FormatDateTime 格式化日期时间为 ISO8601。
func FormatDateTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ContainsString 检查字符串切片是否包含指定元素。
func ContainsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

// RemoveEmpty 移除字符串切片中的空元素。
func RemoveEmpty(slice []string) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if strings.TrimSpace(s) != "" {
			result = append(result, s)
		}
	}
	return result
}

// TruncateString 截断字符串到指定长度并添加省略号。
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// IsBlank 检查字符串是否为空或仅含空白。
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank 检查字符串是否非空。
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// DefaultString 如果字符串为空则返回默认值。
func DefaultString(s, def string) string {
	if IsBlank(s) {
		return def
	}
	return s
}

// SanitizeString 去除字符串首尾空白并规范化内部空格。
func SanitizeString(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// IsValidEmail 简单邮箱格式校验。
func IsValidEmail(s string) bool {
	parts := strings.Split(s, "@")
	if len(parts) != 2 {
		return false
	}
	if strings.Contains(parts[1], ".") {
		return true
	}
	return false
}

// IsNumeric 检查字符串是否全为数字。
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// ClampInt 将整数限制在 [min, max] 范围内。
func ClampInt(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// DaysBetween 计算两个日期之间的天数差。
func DaysBetween(a, b time.Time) int {
	return int(b.Sub(a).Hours() / 24)
}

// StartOfDay 返回日期的零点时刻。
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 返回日期的最后时刻。
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// IsToday 判断日期是否为今天。
func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// PageBounds 计算分页的起止索引。
func PageBounds(page, size, total int) (start, end int) {
	start = (page - 1) * size
	if start < 0 {
		start = 0
	}
	if start > total {
		start = total
	}
	end = start + size
	if end > total {
		end = total
	}
	return start, end
}
