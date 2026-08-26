// Package handler 实现 HTTP 处理器层。
package handler

import (
	"errors"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"pet-care/internal/config"
	"pet-care/internal/model"
	"pet-care/internal/service"
	"pet-care/internal/store"
	"pet-care/pkg/httpx"
	"pet-care/pkg/logger"
)

type Server struct {
	svc *service.Service
	log *logger.Logger
	cfg *config.Config
}

func NewServer(svc *service.Service, log *logger.Logger, cfg *config.Config) *Server {
	return &Server{svc: svc, log: log, cfg: cfg}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	s.registerOwnerRoutes(mux)
	s.registerSpeciesRoutes(mux)
	s.registerPetRoutes(mux)
	s.registerMedicalRecordRoutes(mux)
	s.registerVaccineRecordRoutes(mux)
	s.registerWeightRecordRoutes(mux)
	s.registerFeedingRecordRoutes(mux)
	s.registerStatsRoutes(mux)
	return s.authMiddleware(s.loggingMiddleware(s.recoveryMiddleware(mux)))
}

func (s *Server) maxPageSize() int {
	if s.cfg != nil && s.cfg.MaxPageSize > 0 {
		return s.cfg.MaxPageSize
	}
	return 100
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			if token != s.cfg.APIToken {
				httpx.Unauthorized(w, "无效的访问令牌")
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.log.Infof("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func (s *Server) recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				s.log.Errorf("panic: %v\n%s", rec, debug.Stack())
				httpx.InternalError(w, "服务器内部错误")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func writeServiceError(w http.ResponseWriter, err error) {
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
