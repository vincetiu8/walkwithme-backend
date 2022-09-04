package handlers

import (
	"encoding/json"
	"net/http"

	"walkwithme-backend/accounts"
)

func (s *Server) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var user accounts.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	user.NumRatings = 1
	user.TotalRating = 5

	builder, err := s.Mtr.NewObjectBuilder("did:snr:QmTYGoTAsamNDN2UtGBdHeY3GAigFB41fwXmcSjoAY5Fvd")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	builder.Set("username", user.Username)
	builder.Set("password", user.Password)
	builder.Set("photo_url", user.PhotoURL)
	builder.Set("total_rating", user.TotalRating)
	builder.Set("num_ratings", user.NumRatings)
	builder.SetLabel(user.Username)
	uploadResp, err := builder.Upload()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	user.UserDID = uploadResp.Reference.Did

	s.Users = append(s.Users, user)

	b, err := json.Marshal(&user)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user accounts.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	for _, u := range s.Users {
		if u.Username == user.Username {
			if u.Password == user.Password {
				b, _ := json.Marshal(&u)

				w.WriteHeader(http.StatusOK)
				w.Write(b)
				return
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Incorrect password"))
				return
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("User not found"))
}

func (s *Server) ChangeNameHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Name     string `json:"new_name"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	for i, u := range s.Users {
		if u.Username == user.Username {
			s.Users[i].Name = user.Name
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Successfully changed name"))
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("User not found"))
}

func (s *Server) ChangeUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username    string `json:"username"`
		NewUsername string `json:"new_username"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	for i, u := range s.Users {
		if u.Username == user.Username {
			s.Users[i].Username = user.NewUsername
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Successfully changed username"))
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("User not found"))
}

func (s *Server) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username    string `json:"username"`
		NewPassword string `json:"new_password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	for i, u := range s.Users {
		if u.Username == user.Username {
			s.Users[i].Username = user.NewPassword
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Successfully changed password"))
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("User not found"))
}
