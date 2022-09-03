package handlers

import (
	"encoding/json"
	"net/http"

	"walkwithme-backend/search"
)

func (s *Server) RegisterPlanHandler(w http.ResponseWriter, r *http.Request) {
	var req search.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	s.Requests = append(s.Requests, req)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully registered plan"))
}

func (s *Server) FindPartnerHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}

	var walk *search.Walk
	for _, wa := range s.OngoingWalks {
		if wa.User1 == user.Username || wa.User2 == user.Username {
			walk = &wa
			break
		}
	}
	if walk != nil {
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(walk.Path)
		w.Write(body)
		return
	}

	var req *search.Request
	for _, u := range s.Requests {
		if u.Username == user.Username {
			req = &u
			break
		}
	}

	if req == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No requests made by user"))
		return
	}

	for _, other := range s.Requests {
		if other.Username != req.Username && req.IsValidPartner(other) {
			path := search.GetClosestLocations(*req, other)
			walk := search.Walk{
				User1: req.Username,
				User2: other.Username,
				Path:  path,
			}
			s.OngoingWalks = append(s.OngoingWalks, walk)

			w.WriteHeader(http.StatusOK)
			body, _ := json.Marshal(walk.Path)
			w.Write(body)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("No suitable partners found"))
}
