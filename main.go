package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	motor2 "github.com/sonr-io/sonr/third_party/types/motor"
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

	// Operations when walking
	// Confirm arrived at destination
	r.HandleFunc("/walk/finishedwalk", s.FinishedWalkHandler).Methods("POST")

	return r
}

func main() {
	s, err := handlers.NewServer()
	if err != nil {
		panic(err)
	}

	aesPskKey, _ := os.ReadFile("aesPskKey")

	_, err = s.Mtr.Login(motor2.LoginRequest{
		Did:       "snr1d7w5cr7nxa84gtwgcpv6fhgfrjquvpqygmxq2y",
		Password:  "amongus",
		AesPskKey: aesPskKey,
	})
	if err != nil {
		panic(err)
	}

	queryResp, err := s.Mtr.QueryWhatIsByDid("did:snr:QmTYGoTAsamNDN2UtGBdHeY3GAigFB41fwXmcSjoAY5Fvd")
	fmt.Println(queryResp.WhatIs)

	r := createRouter(s)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
