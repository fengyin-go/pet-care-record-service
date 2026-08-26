package handler

import (
	"net/http"

	"pet-care/internal/model"
	"pet-care/pkg/httpx"
)

func (s *Server) registerMedicalRecordRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/medical-records", s.createMedicalRecord)
	mux.HandleFunc("GET /api/medical-records", s.listMedicalRecords)
	mux.HandleFunc("GET /api/medical-records/{id}", s.getMedicalRecord)
	mux.HandleFunc("PUT /api/medical-records/{id}", s.updateMedicalRecord)
	mux.HandleFunc("DELETE /api/medical-records/{id}", s.deleteMedicalRecord)
}

type createMedicalRecordRequest struct {
	PetID     string `json:"pet_id"`
	VetName   string `json:"vet_name"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
	VisitDate string `json:"visit_date"`
	Notes     string `json:"notes"`
}

func (s *Server) createMedicalRecord(w http.ResponseWriter, r *http.Request) {
	var req createMedicalRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	m, err := s.svc.CreateMedicalRecord(model.MedicalRecord{PetID: req.PetID, VetName: req.VetName, Diagnosis: req.Diagnosis, Treatment: req.Treatment, VisitDate: req.VisitDate, Notes: req.Notes})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, m)
}

func (s *Server) listMedicalRecords(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.MedicalRecordFilter{
		PetID:     r.URL.Query().Get("pet_id"),
		VetName:   r.URL.Query().Get("vet_name"),
		Keyword:   r.URL.Query().Get("keyword"),
		VisitDate: r.URL.Query().Get("visit_date"),
	}
	items, total, err := s.svc.ListMedicalRecords(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getMedicalRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	m, err := s.svc.GetMedicalRecord(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, m)
}

type updateMedicalRecordRequest struct {
	PetID     string `json:"pet_id"`
	VetName   string `json:"vet_name"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
	VisitDate string `json:"visit_date"`
	Notes     string `json:"notes"`
}

func (s *Server) updateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateMedicalRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	m, err := s.svc.UpdateMedicalRecord(id, model.MedicalRecord{PetID: req.PetID, VetName: req.VetName, Diagnosis: req.Diagnosis, Treatment: req.Treatment, VisitDate: req.VisitDate, Notes: req.Notes})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, m)
}

func (s *Server) deleteMedicalRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteMedicalRecord(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
