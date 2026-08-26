package handler

import (
	"net/http"

	"pet-care/internal/model"
	"pet-care/pkg/httpx"
)

func (s *Server) registerPetRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/pets", s.createPet)
	mux.HandleFunc("GET /api/pets", s.listPets)
	mux.HandleFunc("GET /api/pets/{id}", s.getPet)
	mux.HandleFunc("PUT /api/pets/{id}", s.updatePet)
	mux.HandleFunc("DELETE /api/pets/{id}", s.deletePet)
	mux.HandleFunc("POST /api/pets/{id}/transition", s.transitionPetStatus)
	mux.HandleFunc("DELETE /api/pets/{id}/with-records", s.deletePetWithRecords)
}

type createPetRequest struct {
	Name      string `json:"name"`
	SpeciesID string `json:"species_id"`
	Breed     string `json:"breed"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	OwnerID   string `json:"owner_id"`
}

func (s *Server) createPet(w http.ResponseWriter, r *http.Request) {
	var req createPetRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	p, err := s.svc.CreatePet(model.Pet{Name: req.Name, SpeciesID: req.SpeciesID, Breed: req.Breed, Gender: req.Gender, BirthDate: req.BirthDate, OwnerID: req.OwnerID})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, p)
}

func (s *Server) listPets(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.PetFilter{
		Status:    r.URL.Query().Get("status"),
		SpeciesID: r.URL.Query().Get("species_id"),
		OwnerID:   r.URL.Query().Get("owner_id"),
		Keyword:   r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListPets(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getPet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	p, err := s.svc.GetPet(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, p)
}

type updatePetRequest struct {
	Name      string `json:"name"`
	SpeciesID string `json:"species_id"`
	Breed     string `json:"breed"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	OwnerID   string `json:"owner_id"`
}

func (s *Server) updatePet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updatePetRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	p, err := s.svc.UpdatePet(id, model.Pet{Name: req.Name, SpeciesID: req.SpeciesID, Breed: req.Breed, Gender: req.Gender, BirthDate: req.BirthDate, OwnerID: req.OwnerID})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, p)
}

func (s *Server) deletePet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeletePet(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type transitionPetStatusRequest struct {
	Status string `json:"status"`
}

func (s *Server) transitionPetStatus(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req transitionPetStatusRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	p, err := s.svc.TransitionPetStatus(id, req.Status)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, p)
}

func (s *Server) deletePetWithRecords(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeletePetWithRecords(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
