package home

import (
	"fmt"
	"net/http"

	us "github.com/atamis/wt3/data/usersession"
)

func Home(w http.ResponseWriter, r *http.Request) {
	sess, err := us.MakeSession(r)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sess.Save(w, r)

	w.Header().Set("Content-Type", "text/html")

	if sess.LoggedIn {
		homeLoggedOn(sess, w, r)
	} else {
		homeLoggedOff(sess, w, r)
	}

	w.Write([]byte("<a href=\"/login\">Login</a> or <a href=\"/logout\">Logout</a>"))
}

func homeLoggedOn(sess us.Session, w http.ResponseWriter, r *http.Request) {
	user, err := sess.CurrentUser()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error. Session state: %v", sess.Debug())))
		return
	}
	w.Write([]byte(fmt.Sprintf("You are logged in. Welcome, %v.", user.Username)))
}

func homeLoggedOff(sess us.Session, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are not logged on."))
}
