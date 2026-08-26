// Package validator 提供通用验证工具。
package validator

import (
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	phoneRegex   = regexp.MustCompile(`^\+?[0-9\-\s]{7,20}$`)
	chineseRegex = regexp.MustCompile(`^[\p{Han}]+$`)
)

// Validator 通用验证器。
type Validator struct {
	errors []string
}

// New 创建新验证器。
func New() *Validator {
	return &Validator{errors: make([]string, 0)}
}

// IsValid 返回是否有错误。
func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

// Errors 返回错误列表。
func (v *Validator) Errors() []string {
	return v.errors
}

// ErrorString 返回合并的错误字符串。
func (v *Validator) ErrorString() string {
	return strings.Join(v.errors, "; ")
}

// Required 验证必填。
func (v *Validator) Required(value, fieldName string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, fmt.Sprintf("%s 不能为空", fieldName))
	}
	return v
}

// MinLength 验证最小长度。
func (v *Validator) MinLength(value string, min int, fieldName string) *Validator {
	if utf8.RuneCountInString(value) < min {
		v.errors = append(v.errors, fmt.Sprintf("%s 长度不能小于 %d", fieldName, min))
	}
	return v
}

// MaxLength 验证最大长度。
func (v *Validator) MaxLength(value string, max int, fieldName string) *Validator {
	if utf8.RuneCountInString(value) > max {
		v.errors = append(v.errors, fmt.Sprintf("%s 长度不能大于 %d", fieldName, max))
	}
	return v
}

// Range 验证数值范围。
func (v *Validator) Range(value float64, min, max float64, fieldName string) *Validator {
	if value < min || value > max {
		v.errors = append(v.errors, fmt.Sprintf("%s 必须在 %.2f 到 %.2f 之间", fieldName, min, max))
	}
	return v
}

// Positive 验证正数。
func (v *Validator) Positive(value float64, fieldName string) *Validator {
	if value <= 0 {
		v.errors = append(v.errors, fmt.Sprintf("%s 必须大于 0", fieldName))
	}
	return v
}

// Email 验证邮箱格式。
func (v *Validator) Email(value string, fieldName string) *Validator {
	if strings.TrimSpace(value) == "" {
		return v
	}
	_, err := mail.ParseAddress(value)
	if err != nil {
		v.errors = append(v.errors, fmt.Sprintf("%s 格式不正确", fieldName))
	}
	return v
}

// Phone 验证电话格式。
func (v *Validator) Phone(value string, fieldName string) *Validator {
	if !phoneRegex.MatchString(value) {
		v.errors = append(v.errors, fmt.Sprintf("%s 格式不正确", fieldName))
	}
	return v
}

// InEnum 验证值在枚举列表中。
func (v *Validator) InEnum(value string, allowed []string, fieldName string) *Validator {
	for _, a := range allowed {
		if value == a {
			return v
		}
	}
	v.errors = append(v.errors, fmt.Sprintf("%s 值不合法", fieldName))
	return v
}

// Numeric 验证字符串是否为数字。
func (v *Validator) Numeric(value, fieldName string) *Validator {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		v.errors = append(v.errors, fmt.Sprintf("%s 必须是数字", fieldName))
	}
	return v
}

// DateFormat 验证日期格式。
func (v *Validator) DateFormat(value, layout, fieldName string) *Validator {
	if value == "" {
		return v
	}
	// 简单校验 YYYY-MM-DD 格式
	if layout == "2006-01-02" {
		parts := strings.Split(value, "-")
		if len(parts) != 3 {
			v.errors = append(v.errors, fmt.Sprintf("%s 日期格式不正确", fieldName))
			return v
		}
		year, err1 := strconv.Atoi(parts[0])
		month, err2 := strconv.Atoi(parts[1])
		day, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			v.errors = append(v.errors, fmt.Sprintf("%s 日期格式不正确", fieldName))
			return v
		}
		if year < 1900 || year > 2100 || month < 1 || month > 12 || day < 1 || day > 31 {
			v.errors = append(v.errors, fmt.Sprintf("%s 日期范围不正确", fieldName))
			return v
		}
	}
	return v
}

// IsChinese 检查是否为纯中文。
func IsChinese(s string) bool {
	return chineseRegex.MatchString(s)
}

// IsASCII 检查是否为纯 ASCII。
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}
