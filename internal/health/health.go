// Package health 提供服务健康检查能力。
package health

import (
	"runtime"
	"time"
)

// Status 健康状态。
type Status string

const (
	StatusHealthy   Status = "healthy"
	StatusDegraded  Status = "degraded"
	StatusUnhealthy Status = "unhealthy"
)

// Check 健康检查项。
type Check struct {
	Name      string        `json:"name"`
	Status    Status        `json:"status"`
	Message   string        `json:"message"`
	Latency   time.Duration `json:"latency"`
	CheckedAt time.Time     `json:"checked_at"`
}

// Report 健康检查报告。
type Report struct {
	Status    Status    `json:"status"`
	Version   string    `json:"version"`
	Uptime    string    `json:"uptime"`
	Checks    []Check   `json:"checks"`
	CheckedAt time.Time `json:"checked_at"`
}

// Checker 健康检查器。
type Checker struct {
	version   string
	startTime time.Time
	checks    []func() Check
}

// NewChecker 创建健康检查器。
func NewChecker(version string) *Checker {
	return &Checker{
		version:   version,
		startTime: time.Now(),
		checks:    make([]func() Check, 0),
	}
}

// Register 注册健康检查函数。
func (c *Checker) Register(name string, checkFn func() (Status, string)) {
	c.checks = append(c.checks, func() Check {
		start := time.Now()
		status, msg := checkFn()
		return Check{
			Name:      name,
			Status:    status,
			Message:   msg,
			Latency:   time.Since(start),
			CheckedAt: time.Now(),
		}
	})
}

// Check 执行所有健康检查。
func (c *Checker) Check() *Report {
	checks := make([]Check, 0, len(c.checks))
	overall := StatusHealthy
	for _, fn := range c.checks {
		check := fn()
		checks = append(checks, check)
		if check.Status == StatusUnhealthy {
			overall = StatusUnhealthy
		} else if check.Status == StatusDegraded && overall == StatusHealthy {
			overall = StatusDegraded
		}
	}
	return &Report{
		Status:    overall,
		Version:   c.version,
		Uptime:    time.Since(c.startTime).String(),
		Checks:    checks,
		CheckedAt: time.Now(),
	}
}

// MemoryCheck 内存检查。
func MemoryCheck() (Status, string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	allocMB := m.Alloc / 1024 / 1024
	if allocMB > 1024 {
		return StatusDegraded, "内存使用超过 1GB"
	}
	return StatusHealthy, "内存使用正常"
}

// GoroutineCheck Goroutine 数量检查。
func GoroutineCheck() (Status, string) {
	count := runtime.NumGoroutine()
	if count > 1000 {
		return StatusDegraded, "Goroutine 数量超过 1000"
	}
	return StatusHealthy, "Goroutine 数量正常"
}

// SimpleCheck 简单健康检查（始终健康）。
func SimpleCheck() (Status, string) {
	return StatusHealthy, "服务运行正常"
}

// BuildInfo 构建信息。
type BuildInfo struct {
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	NumCPU    int    `json:"num_cpu"`
}

// GetBuildInfo 获取构建信息。
func GetBuildInfo(version string) *BuildInfo {
	return &BuildInfo{
		Version:   version,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		NumCPU:    runtime.NumCPU(),
	}
}

// SystemStats 系统统计。
type SystemStats struct {
	NumGoroutine int    `json:"num_goroutine"`
	NumCPU       int    `json:"num_cpu"`
	MemoryAlloc  uint64 `json:"memory_alloc"`
	MemoryTotal  uint64 `json:"memory_total"`
	NumGC        uint32 `json:"num_gc"`
}

// GetSystemStats 获取系统统计。
func GetSystemStats() *SystemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &SystemStats{
		NumGoroutine: runtime.NumGoroutine(),
		NumCPU:       runtime.NumCPU(),
		MemoryAlloc:  m.Alloc,
		MemoryTotal:  m.TotalAlloc,
		NumGC:        m.NumGC,
	}
}
