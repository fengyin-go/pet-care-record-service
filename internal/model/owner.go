package model

import (
	"strings"
	"time"
)

type Owner struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (o *Owner) Validate() error {
	o.Name = strings.TrimSpace(o.Name)
	o.Email = strings.TrimSpace(o.Email)
	o.Phone = strings.TrimSpace(o.Phone)
	o.Address = strings.TrimSpace(o.Address)
	if o.Name == "" {
		return NewValidationError("name", "主人姓名不能为空")
	}
	if o.Email == "" {
		return NewValidationError("email", "邮箱不能为空")
	}
	if o.Phone == "" {
		return NewValidationError("phone", "电话不能为空")
	}
	if len(o.Phone) < 7 {
		return NewValidationError("phone", "电话格式不正确")
	}
	return nil
}

type OwnerFilter struct {
	Name    string
	Email   string
	Keyword string
}

func (f OwnerFilter) Match(o *Owner) bool {
	if f.Name != "" && o.Name != f.Name {
		return false
	}
	if f.Email != "" && o.Email != f.Email {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(o.Name), k) &&
			!strings.Contains(strings.ToLower(o.Email), k) &&
			!strings.Contains(strings.ToLower(o.Phone), k) {
			return false
		}
	}
	return true
}
