package handler

import (
	"net/http"

	"pet-care/pkg/httpx"
)

func (s *Server) registerStatsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/stats/global", s.globalStats)
	mux.HandleFunc("GET /api/stats/pets/{id}/records", s.petRecordStats)
	mux.HandleFunc("GET /api/stats/pets-by-species", s.petsBySpecies)
	mux.HandleFunc("GET /api/pets/{id}/health-profile", s.exportPetHealthProfile)
}

func (s *Server) globalStats(w http.ResponseWriter, r *http.Request) {
	stats, err := s.svc.GlobalStats()
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, stats)
}

func (s *Server) petRecordStats(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	stats, err := s.svc.PetRecordStats(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, stats)
}

func (s *Server) petsBySpecies(w http.ResponseWriter, r *http.Request) {
	counts, err := s.svc.CountPetsBySpecies()
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, counts)
}

func (s *Server) exportPetHealthProfile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	profile, err := s.svc.ExportPetHealthProfile(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, profile)
}
