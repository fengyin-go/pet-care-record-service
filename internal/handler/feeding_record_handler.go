package handler

import (
	"net/http"
	"time"

	"pet-care/internal/model"
	"pet-care/pkg/httpx"
)

func (s *Server) registerFeedingRecordRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/feeding-records", s.createFeedingRecord)
	mux.HandleFunc("GET /api/feeding-records", s.listFeedingRecords)
	mux.HandleFunc("GET /api/feeding-records/{id}", s.getFeedingRecord)
	mux.HandleFunc("PUT /api/feeding-records/{id}", s.updateFeedingRecord)
	mux.HandleFunc("DELETE /api/feeding-records/{id}", s.deleteFeedingRecord)
}

type createFeedingRecordRequest struct {
	PetID  string    `json:"pet_id"`
	Food   string    `json:"food"`
	Amount float64   `json:"amount"`
	FedAt  time.Time `json:"fed_at"`
}

func (s *Server) createFeedingRecord(w http.ResponseWriter, r *http.Request) {
	var req createFeedingRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	f, err := s.svc.CreateFeedingRecord(model.FeedingRecord{PetID: req.PetID, Food: req.Food, Amount: req.Amount, FedAt: req.FedAt})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, f)
}

func (s *Server) listFeedingRecords(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.FeedingRecordFilter{
		PetID: r.URL.Query().Get("pet_id"),
		Food:  r.URL.Query().Get("food"),
	}
	items, total, err := s.svc.ListFeedingRecords(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getFeedingRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	f, err := s.svc.GetFeedingRecord(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, f)
}

type updateFeedingRecordRequest struct {
	PetID  string    `json:"pet_id"`
	Food   string    `json:"food"`
	Amount float64   `json:"amount"`
	FedAt  time.Time `json:"fed_at"`
}

func (s *Server) updateFeedingRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateFeedingRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	f, err := s.svc.UpdateFeedingRecord(id, model.FeedingRecord{PetID: req.PetID, Food: req.Food, Amount: req.Amount, FedAt: req.FedAt})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, f)
}

func (s *Server) deleteFeedingRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteFeedingRecord(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
