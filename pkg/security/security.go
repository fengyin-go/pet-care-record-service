// Package security 提供安全相关工具。
package security

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// HashString 使用 SHA-256 对字符串进行哈希。
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// HashWithSalt 带盐哈希。
func HashWithSalt(s, salt string) string {
	return HashString(s + salt)
}

// GenerateSalt 生成随机盐值。
func GenerateSalt(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// TokenGenerator Token 生成器。
type TokenGenerator struct {
	prefix string
}

// NewTokenGenerator 创建 Token 生成器。
func NewTokenGenerator(prefix string) *TokenGenerator {
	return &TokenGenerator{prefix: prefix}
}

// Generate 生成随机 Token。
func (tg *TokenGenerator) Generate() string {
	timestamp := time.Now().UnixNano()
	randomPart := GenerateSalt(16)
	raw := fmt.Sprintf("%s:%d:%s", tg.prefix, timestamp, randomPart)
	return HashString(raw)[:32]
}

// GenerateWithPrefix 生成带前缀的 Token。
func (tg *TokenGenerator) GenerateWithPrefix() string {
	return tg.prefix + "_" + tg.Generate()
}

// MaskString 对字符串进行脱敏处理。
func MaskString(s string, keepPrefix, keepSuffix int) string {
	if len(s) <= keepPrefix+keepSuffix {
		return strings.Repeat("*", len(s))
	}
	return s[:keepPrefix] + strings.Repeat("*", len(s)-keepPrefix-keepSuffix) + s[len(s)-keepSuffix:]
}

// MaskEmail 对邮箱进行脱敏。
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return strings.Repeat("*", len(email))
	}
	local := parts[0]
	domain := parts[1]
	if len(local) <= 2 {
		return strings.Repeat("*", len(local)) + "@" + domain
	}
	return local[:2] + strings.Repeat("*", len(local)-2) + "@" + domain
}

// MaskPhone 对电话进行脱敏。
func MaskPhone(phone string) string {
	if len(phone) < 7 {
		return strings.Repeat("*", len(phone))
	}
	return phone[:3] + strings.Repeat("*", len(phone)-6) + phone[len(phone)-3:]
}

// RateLimiterConfig 限流配置。
type RateLimiterConfig struct {
	MaxRequests int
	WindowSize  time.Duration
}

// ValidateAPIKey 校验 API Key 格式。
func ValidateAPIKey(key string) error {
	if len(key) < 16 {
		return fmt.Errorf("API Key 长度不能小于 16")
	}
	if strings.Contains(key, " ") {
		return fmt.Errorf("API Key 不能包含空格")
	}
	return nil
}

// SecureCompare 常量时间字符串比较（防时序攻击）。
func SecureCompare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}

// PasswordPolicy 密码策略。
type PasswordPolicy struct {
	MinLength      int
	RequireUpper   bool
	RequireLower   bool
	RequireDigit   bool
	RequireSpecial bool
}

// ValidatePassword 校验密码强度。
func ValidatePassword(password string, policy PasswordPolicy) error {
	if len(password) < policy.MinLength {
		return fmt.Errorf("密码长度不能小于 %d", policy.MinLength)
	}
	if policy.RequireUpper && !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return fmt.Errorf("密码必须包含大写字母")
	}
	if policy.RequireLower && !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("密码必须包含小写字母")
	}
	if policy.RequireDigit && !strings.ContainsAny(password, "0123456789") {
		return fmt.Errorf("密码必须包含数字")
	}
	return nil
}
