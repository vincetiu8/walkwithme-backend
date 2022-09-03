package handlers

import "net/http"

func (s *Server) FoundPartnerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Found Partner"))
}

func (s *Server) FinishedWalkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Finished Walk"))
}

func (s *Server) CancelWalkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cancel Walk"))
}
