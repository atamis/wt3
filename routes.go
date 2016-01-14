package main

import (
	"github.com/atamis/wt3/act/session"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/login", session.Login).
		Methods("POST")

	r.HandleFunc("/logout", session.Logout).
		Methods("GET")

	r.HandleFunc("/login", session.LoginPage).
		Methods("GET")

	return r
}
