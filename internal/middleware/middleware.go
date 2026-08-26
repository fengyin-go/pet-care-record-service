// Package middleware 提供 HTTP 中间件。
package middleware

import (
	"context"
	"errors"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"pet-care/internal/config"
	"pet-care/internal/model"
	"pet-care/internal/store"
	"pet-care/pkg/httpx"
	"pet-care/pkg/logger"
)

// Recovery 恢复 panic 中间件。
func Recovery(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Errorf("panic: %v\n%s", rec, debug.Stack())
					httpx.InternalError(w, "服务器内部错误")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// Logging 请求日志中间件。
func Logging(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Infof("%s %s %s", r.Method, r.URL.Path, time.Since(start))
		})
	}
}

// Auth 鉴权中间件。
func Auth(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api/") {
				token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
				if token != cfg.APIToken {
					httpx.Unauthorized(w, "无效的访问令牌")
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

// CORSMiddleware 跨域中间件。
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// rateLimiter 基于滑动窗口的简易限流器。
type rateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *rateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-rl.window)
	list := rl.requests[key]
	valid := make([]time.Time, 0, len(list))
	for _, t := range list {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}
	if len(valid) >= rl.limit {
		rl.requests[key] = valid
		return false
	}
	valid = append(valid, now)
	rl.requests[key] = valid
	return true
}

// RateLimit 限流中间件（每 IP 每分钟 60 次）。
func RateLimit() func(http.Handler) http.Handler {
	rl := newRateLimiter(60, time.Minute)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr
			if !rl.allow(key) {
				httpx.Error(w, http.StatusTooManyRequests, 429, "请求过于频繁")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// RequestID 注入请求 ID 中间件。
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = generateRequestID()
		}
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), "request_id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateRequestID() string {
	return time.Now().Format("20060102150405") + randomHex(4)
}

func randomHex(n int) string {
	const hexChars = "0123456789abcdef"
	b := make([]byte, n)
	for i := range b {
		b[i] = hexChars[time.Now().UnixNano()%int64(len(hexChars))]
	}
	return string(b)
}

// WriteServiceError 统一服务错误映射写入。
func WriteServiceError(w http.ResponseWriter, err error) {
	switch {
	case model.IsValidationError(err):
		httpx.BadRequest(w, err.Error())
	case errors.Is(err, store.ErrNotFound):
		httpx.NotFound(w, err.Error())
	case errors.Is(err, store.ErrConflict):
		httpx.Conflict(w, err.Error())
	default:
		httpx.InternalError(w, err.Error())
	}
}
