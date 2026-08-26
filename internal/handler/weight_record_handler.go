package handler

import (
	"net/http"
	"time"

	"pet-care/internal/model"
	"pet-care/pkg/httpx"
)

func (s *Server) registerWeightRecordRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/weight-records", s.createWeightRecord)
	mux.HandleFunc("GET /api/weight-records", s.listWeightRecords)
	mux.HandleFunc("GET /api/weight-records/{id}", s.getWeightRecord)
	mux.HandleFunc("PUT /api/weight-records/{id}", s.updateWeightRecord)
	mux.HandleFunc("DELETE /api/weight-records/{id}", s.deleteWeightRecord)
	mux.HandleFunc("GET /api/pets/{id}/weight-trend", s.getWeightTrend)
}

type createWeightRecordRequest struct {
	PetID      string    `json:"pet_id"`
	Weight     float64   `json:"weight"`
	MeasuredAt time.Time `json:"measured_at"`
}

func (s *Server) createWeightRecord(w http.ResponseWriter, r *http.Request) {
	var req createWeightRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	wrec, err := s.svc.CreateWeightRecord(model.WeightRecord{PetID: req.PetID, Weight: req.Weight, MeasuredAt: req.MeasuredAt})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, wrec)
}

func (s *Server) listWeightRecords(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.WeightRecordFilter{
		PetID: r.URL.Query().Get("pet_id"),
	}
	items, total, err := s.svc.ListWeightRecords(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getWeightRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	wrec, err := s.svc.GetWeightRecord(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, wrec)
}

type updateWeightRecordRequest struct {
	PetID      string    `json:"pet_id"`
	Weight     float64   `json:"weight"`
	MeasuredAt time.Time `json:"measured_at"`
}

func (s *Server) updateWeightRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateWeightRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	wrec, err := s.svc.UpdateWeightRecord(id, model.WeightRecord{PetID: req.PetID, Weight: req.Weight, MeasuredAt: req.MeasuredAt})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, wrec)
}

func (s *Server) deleteWeightRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteWeightRecord(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

func (s *Server) getWeightTrend(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	items, err := s.svc.GetWeightTrend(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, items)
}
