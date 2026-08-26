package handler

import (
	"net/http"

	"pet-care/internal/model"
	"pet-care/pkg/httpx"
)

func (s *Server) registerSpeciesRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/species", s.createSpecies)
	mux.HandleFunc("GET /api/species", s.listSpecies)
	mux.HandleFunc("GET /api/species/{id}", s.getSpecies)
	mux.HandleFunc("PUT /api/species/{id}", s.updateSpecies)
	mux.HandleFunc("DELETE /api/species/{id}", s.deleteSpecies)
}

type createSpeciesRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (s *Server) createSpecies(w http.ResponseWriter, r *http.Request) {
	var req createSpeciesRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	sp, err := s.svc.CreateSpecies(model.Species{Name: req.Name, Category: req.Category, Description: req.Description})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, sp)
}

func (s *Server) listSpecies(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.SpeciesFilter{
		Category: r.URL.Query().Get("category"),
		Name:     r.URL.Query().Get("name"),
		Keyword:  r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListSpecies(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getSpecies(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sp, err := s.svc.GetSpecies(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, sp)
}

type updateSpeciesRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (s *Server) updateSpecies(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateSpeciesRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	sp, err := s.svc.UpdateSpecies(id, model.Species{Name: req.Name, Category: req.Category, Description: req.Description})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, sp)
}

func (s *Server) deleteSpecies(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteSpecies(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
