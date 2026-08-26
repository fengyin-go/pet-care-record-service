// Package datetime 提供日期时间相关工具函数。
package utils

import (
	"fmt"
	"time"
)

// TimeRange 时间范围结构。
type TimeRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// NewTimeRange 创建时间范围。
func NewTimeRange(start, end time.Time) *TimeRange {
	return &TimeRange{Start: start, End: end}
}

// Duration 返回时间范围持续时长。
func (tr *TimeRange) Duration() time.Duration {
	return tr.End.Sub(tr.Start)
}

// Contains 检查时间点是否在范围内。
func (tr *TimeRange) Contains(t time.Time) bool {
	return (t.Equal(tr.Start) || t.After(tr.Start)) && (t.Equal(tr.End) || t.Before(tr.End))
}

// Overlaps 检查两个时间范围是否重叠。
func (tr *TimeRange) Overlaps(other *TimeRange) bool {
	return tr.Start.Before(other.End) && tr.End.After(other.Start)
}

// AgeFromBirthDate 根据出生日期计算年龄（年/月）。
func AgeFromBirthDate(birthDate string) (years, months int, err error) {
	t, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return 0, 0, fmt.Errorf("出生日期格式错误: %w", err)
	}
	now := time.Now()
	years = now.Year() - t.Year()
	months = int(now.Month()) - int(t.Month())
	if now.Day() < t.Day() {
		months--
	}
	if months < 0 {
		years--
		months += 12
	}
	return years, months, nil
}

// AgeString 返回人类可读的年龄字符串。
func AgeString(birthDate string) (string, error) {
	years, months, err := AgeFromBirthDate(birthDate)
	if err != nil {
		return "", err
	}
	if years > 0 {
		if months > 0 {
			return fmt.Sprintf("%d岁%d个月", years, months), nil
		}
		return fmt.Sprintf("%d岁", years), nil
	}
	return fmt.Sprintf("%d个月", months), nil
}

// NextOccurrence 计算下一个指定时间的发生时刻。
func NextOccurrence(hour, minute int) time.Time {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
	if next.Before(now) || next.Equal(now) {
		next = next.AddDate(0, 0, 1)
	}
	return next
}

// MonthStartEnd 返回指定月份的起止时间。
func MonthStartEnd(year, month int) (start, end time.Time) {
	loc := time.Now().Location()
	start = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	if month == 12 {
		end = time.Date(year+1, 1, 1, 0, 0, 0, 0, loc).Add(-time.Nanosecond)
	} else {
		end = time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, loc).Add(-time.Nanosecond)
	}
	return start, end
}

// WeekRange 返回本周的起止时间（周一到周日）。
func WeekRange(t time.Time) (start, end time.Time) {
	loc := t.Location()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	daysFromMonday := weekday - 1
	start = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -daysFromMonday)
	end = start.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return start, end
}

// DurationToHuman 将时长转为人类可读字符串。
func DurationToHuman(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%d秒", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%d分钟", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%d小时", int(d.Hours()))
	}
	days := int(d.Hours() / 24)
	return fmt.Sprintf("%d天", days)
}

// AddBusinessDays 添加工作日（跳过周末）。
func AddBusinessDays(t time.Time, days int) time.Time {
	result := t
	added := 0
	for added < days {
		result = result.AddDate(0, 0, 1)
		weekday := result.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			added++
		}
	}
	return result
}
