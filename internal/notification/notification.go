// Package notification 提供通知提醒服务。
package notification

import (
	"fmt"
	"time"
)

// NotificationType 通知类型。
type NotificationType string

const (
	NotificationVaccine NotificationType = "vaccine"
	NotificationMedical NotificationType = "medical"
	NotificationWeight  NotificationType = "weight"
	NotificationFeeding NotificationType = "feeding"
	NotificationGeneral NotificationType = "general"
)

// Priority 优先级。
type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
)

// Notification 通知消息。
type Notification struct {
	ID        string           `json:"id"`
	Type      NotificationType `json:"type"`
	Title     string           `json:"title"`
	Message   string           `json:"message"`
	Priority  Priority         `json:"priority"`
	PetID     string           `json:"pet_id"`
	Read      bool             `json:"read"`
	CreatedAt time.Time        `json:"created_at"`
}

// NotificationService 通知服务。
type NotificationService struct {
	notifications []*Notification
	maxSize       int
}

// NewNotificationService 创建通知服务。
func NewNotificationService(maxSize int) *NotificationService {
	if maxSize <= 0 {
		maxSize = 100
	}
	return &NotificationService{
		notifications: make([]*Notification, 0),
		maxSize:       maxSize,
	}
}

// Create 创建通知。
func (ns *NotificationService) Create(nType NotificationType, title, message string, priority Priority, petID string) *Notification {
	n := &Notification{
		ID:        generateID(),
		Type:      nType,
		Title:     title,
		Message:   message,
		Priority:  priority,
		PetID:     petID,
		Read:      false,
		CreatedAt: time.Now(),
	}
	ns.notifications = append(ns.notifications, n)
	if len(ns.notifications) > ns.maxSize {
		ns.notifications = ns.notifications[len(ns.notifications)-ns.maxSize:]
	}
	return n
}

// List 列出所有通知。
func (ns *NotificationService) List() []*Notification {
	result := make([]*Notification, len(ns.notifications))
	copy(result, ns.notifications)
	return result
}

// ListUnread 列出未读通知。
func (ns *NotificationService) ListUnread() []*Notification {
	result := make([]*Notification, 0)
	for _, n := range ns.notifications {
		if !n.Read {
			result = append(result, n)
		}
	}
	return result
}

// ListByType 按类型列出通知。
func (ns *NotificationService) ListByType(nType NotificationType) []*Notification {
	result := make([]*Notification, 0)
	for _, n := range ns.notifications {
		if n.Type == nType {
			result = append(result, n)
		}
	}
	return result
}

// ListByPet 按宠物列出通知。
func (ns *NotificationService) ListByPet(petID string) []*Notification {
	result := make([]*Notification, 0)
	for _, n := range ns.notifications {
		if n.PetID == petID {
			result = append(result, n)
		}
	}
	return result
}

// MarkAsRead 标记通知为已读。
func (ns *NotificationService) MarkAsRead(id string) bool {
	for _, n := range ns.notifications {
		if n.ID == id {
			n.Read = true
			return true
		}
	}
	return false
}

// MarkAllAsRead 标记所有通知为已读。
func (ns *NotificationService) MarkAllAsRead() {
	for _, n := range ns.notifications {
		n.Read = true
	}
}

// Delete 删除通知。
func (ns *NotificationService) Delete(id string) bool {
	for i, n := range ns.notifications {
		if n.ID == id {
			ns.notifications = append(ns.notifications[:i], ns.notifications[i+1:]...)
			return true
		}
	}
	return false
}

// CountUnread 统计未读数量。
func (ns *NotificationService) CountUnread() int {
	count := 0
	for _, n := range ns.notifications {
		if !n.Read {
			count++
		}
	}
	return count
}

// CountByPriority 按优先级统计。
func (ns *NotificationService) CountByPriority(priority Priority) int {
	count := 0
	for _, n := range ns.notifications {
		if n.Priority == priority {
			count++
		}
	}
	return count
}

// VaccineDueNotification 生成疫苗到期通知。
func VaccineDueNotification(petID, petName, vaccineName string, daysUntil int) *Notification {
	priority := PriorityMedium
	if daysUntil <= 0 {
		priority = PriorityHigh
	} else if daysUntil <= 7 {
		priority = PriorityHigh
	}
	return &Notification{
		ID:        generateID(),
		Type:      NotificationVaccine,
		Title:     fmt.Sprintf("%s 的疫苗即将到期", petName),
		Message:   fmt.Sprintf("疫苗 %s 还有 %d 天到期，请及时接种", vaccineName, daysUntil),
		Priority:  priority,
		PetID:     petID,
		Read:      false,
		CreatedAt: time.Now(),
	}
}

// WeightCheckNotification 生成体重检查通知。
func WeightCheckNotification(petID, petName string) *Notification {
	return &Notification{
		ID:        generateID(),
		Type:      NotificationWeight,
		Title:     fmt.Sprintf("该为 %s 记录体重了", petName),
		Message:   "建议定期记录宠物体重，关注健康变化",
		Priority:  PriorityLow,
		PetID:     petID,
		Read:      false,
		CreatedAt: time.Now(),
	}
}

func generateID() string {
	return fmt.Sprintf("notif_%d", time.Now().UnixNano())
}
