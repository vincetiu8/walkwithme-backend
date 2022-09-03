package main

import (
	"fmt"

	"github.com/gorilla/mux"
	motor2 "github.com/sonr-io/sonr/third_party/types/motor"
	"github.com/sonr-io/sonr/x/schema/types"
	"walkwithme-backend/handlers"
)

func createRouter(s *handlers.Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.PingHandler).Methods("GET")

	// Account Management
	// Create Account
	r.HandleFunc("/accounts/create", s.CreateAccountHandler).Methods("POST")

	// Log into account
	r.HandleFunc("/accounts/login", s.LoginHandler).Methods("POST")

	// Change Username
	r.HandleFunc("/accounts/username", s.ChangeUsernameHandler).Methods("PUT")

	// Change Password
	r.HandleFunc("/accounts/password", s.ChangePasswordHandler).Methods("PUT")

	// Finding other users
	// Register travel plan
	r.HandleFunc("/search/registerplan", s.RegisterPlanHandler).Methods("POST")

	// Search for users
	r.HandleFunc("/search/findpartner", s.FindPartnerHandler).Methods("GET")

	// Confirm walk
	r.HandleFunc("/search/confirmwalk", s.ConfirmWalkHandler).Methods("POST")

	// Operations when walking
	// Confirm seen other user
	r.HandleFunc("/walk/foundpartner", s.FoundPartnerHandler).Methods("POST")

	// Confirm arrived at destination
	r.HandleFunc("/walk/finishedwalk", s.FinishedWalkHandler).Methods("POST")

	// Cancel walk
	r.HandleFunc("/walk/cancelwalk", s.CancelWalkHandler).Methods("POST")

	return r
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	s, err := handlers.NewServer()
	if err != nil {
		panic(err)
	}

	aesDscKey := []byte("12345678901234567890123456789012")

	createResponse, err := s.Mtr.CreateAccount(motor2.CreateAccountRequest{
		Password:  "amongus",
		AesDscKey: aesDscKey,
	})
	if err != nil {
		panic(err)
	}

	loginResponse, err := s.Mtr.Login(motor2.LoginRequest{
		Did:       string(createResponse.GetDidDocument()),
		Password:  "amongus",
		AesPskKey: createResponse.AesPsk,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(createResponse.AesPsk)
	fmt.Println(loginResponse)

	createSchemaResponse, err := s.Mtr.CreateSchema(motor2.CreateSchemaRequest{
		Label: "user",
		Fields: map[string]types.SchemaKind{
			"username": types.SchemaKind_STRING,
			"password": types.SchemaKind_STRING,
		},
		Metadata: nil,
	})

	fmt.Println(createSchemaResponse)

	//r := createRouter(s)
	//fmt.Println("Listening on port 8080")
	//log.Fatal(http.ListenAndServe("localhost:8080", r))
}
