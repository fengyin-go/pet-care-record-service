package handler

import (
	"net/http"
	"strconv"
	"time"

	"pet-care/internal/model"
	"pet-care/pkg/httpx"
)

func (s *Server) registerVaccineRecordRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/vaccine-records", s.createVaccineRecord)
	mux.HandleFunc("GET /api/vaccine-records", s.listVaccineRecords)
	mux.HandleFunc("GET /api/vaccine-records/{id}", s.getVaccineRecord)
	mux.HandleFunc("PUT /api/vaccine-records/{id}", s.updateVaccineRecord)
	mux.HandleFunc("DELETE /api/vaccine-records/{id}", s.deleteVaccineRecord)
	mux.HandleFunc("GET /api/vaccine-records/expiring", s.listExpiringVaccines)
	mux.HandleFunc("GET /api/vaccine-records/expired", s.listExpiredVaccines)
}

type createVaccineRecordRequest struct {
	PetID        string    `json:"pet_id"`
	VaccineName  string    `json:"vaccine_name"`
	Dose         int       `json:"dose"`
	VaccinatedAt time.Time `json:"vaccinated_at"`
	NextDueAt    time.Time `json:"next_due_at"`
}

func (s *Server) createVaccineRecord(w http.ResponseWriter, r *http.Request) {
	var req createVaccineRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	v, err := s.svc.CreateVaccineRecord(model.VaccineRecord{PetID: req.PetID, VaccineName: req.VaccineName, Dose: req.Dose, VaccinatedAt: req.VaccinatedAt, NextDueAt: req.NextDueAt})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, v)
}

func (s *Server) listVaccineRecords(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.VaccineRecordFilter{
		PetID:       r.URL.Query().Get("pet_id"),
		VaccineName: r.URL.Query().Get("vaccine_name"),
		Status:      r.URL.Query().Get("status"),
	}
	items, total, err := s.svc.ListVaccineRecords(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getVaccineRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	v, err := s.svc.GetVaccineRecord(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, v)
}

type updateVaccineRecordRequest struct {
	PetID        string    `json:"pet_id"`
	VaccineName  string    `json:"vaccine_name"`
	Dose         int       `json:"dose"`
	VaccinatedAt time.Time `json:"vaccinated_at"`
	NextDueAt    time.Time `json:"next_due_at"`
}

func (s *Server) updateVaccineRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateVaccineRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	v, err := s.svc.UpdateVaccineRecord(id, model.VaccineRecord{PetID: req.PetID, VaccineName: req.VaccineName, Dose: req.Dose, VaccinatedAt: req.VaccinatedAt, NextDueAt: req.NextDueAt})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, v)
}

func (s *Server) deleteVaccineRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteVaccineRecord(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

func (s *Server) listExpiringVaccines(w http.ResponseWriter, r *http.Request) {
	days, _ := strconv.Atoi(r.URL.Query().Get("days"))
	if days <= 0 {
		days = 30
	}
	items, err := s.svc.ListExpiringVaccines(days)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, items)
}

func (s *Server) listExpiredVaccines(w http.ResponseWriter, r *http.Request) {
	items, err := s.svc.ListExpiredVaccines()
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, items)
}
