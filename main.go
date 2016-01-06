package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

type State struct {
	On      bool
	Answers []uint
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func LoginFunc(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if r.FormValue("password") == "correct" {
		session.Values["userid"] = r.FormValue("username")
	}

	session.Save(r, w)

	r.ParseForm()

	w.Write([]byte("Gorilla!\n"))
	w.Write([]byte(r.FormValue("username")))
	w.Write([]byte(":"))
	w.Write([]byte(r.FormValue("password")))
	w.Write([]byte("\nGorilla!\n"))
	w.Write([]byte(fmt.Sprint(session.Values["userid"])))

}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<form method=\"post\" action=\"/login\">"))
	w.Write([]byte("<input type=\"text\" name=\"username\">"))
	w.Write([]byte("<input type=\"password\" name=\"password\">"))

	w.Write([]byte("<input type=\"submit\" name=\"submit\">"))
	w.Write([]byte("</form>"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", LoginFunc).
		Methods("POST")

	r.HandleFunc("/login", LoginPage).
		Methods("GET")

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)

}
