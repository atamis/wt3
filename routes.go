package main

import (
	"github.com/atamis/wt3/act/home"
	"github.com/atamis/wt3/act/polls"
	"github.com/atamis/wt3/act/session"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", home.Home).
		Methods("GET")

	r.HandleFunc("/login", session.Login).
		Methods("POST")

	r.HandleFunc("/logout", session.Logout).
		Methods("GET")

	r.HandleFunc("/login", session.LoginPage).
		Methods("GET")

	r.HandleFunc("/polls/{id}", polls.Poll).
		Methods("GET")

	return r
}
