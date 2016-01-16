package session

import (
	"fmt"
	us "github.com/atamis/wt3/data/usersession"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	sess, err := us.MakeSession(r)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	login := sess.LogIn(r.FormValue("username"), r.FormValue("password"))

	sess.Save(w, r)

	w.Write([]byte(sess.Debug()))
	w.Write([]byte("Gorilla!\n"))
	w.Write([]byte(r.FormValue("username")))
	w.Write([]byte(":"))
	w.Write([]byte(r.FormValue("password")))
	w.Write([]byte("\nGorilla!\n"))
	w.Write([]byte(fmt.Sprintf("%v\n", sess.UserId)))
	w.Write([]byte(fmt.Sprintf("Logged in successfully: %v", login)))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sess, err := us.MakeSession(r)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = sess.Logout()

	sess.Save(w, r)

	if err != nil {
		w.Write([]byte("Not logged in"))
	}

	w.Write([]byte("Logged out"))
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	sess, err := us.MakeSession(r)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	sess.Save(w, r)

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(sess.Debug()))
	w.Write([]byte("<form method=\"post\" action=\"/login\">"))
	w.Write([]byte("<input type=\"text\" name=\"username\">"))
	w.Write([]byte("<input type=\"password\" name=\"password\">"))

	w.Write([]byte("<input type=\"submit\" name=\"submit\">"))
	w.Write([]byte("</form>"))
}
