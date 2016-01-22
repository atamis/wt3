package polls

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/atamis/wt3/data/poll"
	us "github.com/atamis/wt3/data/usersession"
	"github.com/gorilla/mux"
)

func findPoll(r *http.Request) (poll.Poll, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		return poll.Poll{}, err
	}

	p, err := poll.Find(id)
	p.LoadAnswers()

	if err != nil {
		return poll.Poll{}, err
	}

	return p, nil

}

func Poll(w http.ResponseWriter, r *http.Request) {
	poll, err := findPoll(r)

	if err != nil {
		http.Error(w, "Error: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf("%v", poll.Collated())))
	fmt.Fprintf(w, `<form method="post" action="%s">
	<input type="text" name="answer">
	<input type="submit" name="submit">
	</form>`, r.URL.Path)
}

func Answer(w http.ResponseWriter, r *http.Request) {
	poll, err := findPoll(r)
	sess, err := us.MakeSession(r)
	user, err := sess.CurrentUser()

	if err != nil {
		http.Error(w, "Error: "+err.Error(), 500)
		return
	}

	r.ParseForm()

	ans, err := strconv.Atoi(r.FormValue("answer"))

	if err != nil {
		http.Error(w, "Error: "+err.Error(), 500)
		return
	}

	poll.AddAnswer(ans, user.Id)

	http.Redirect(w, r, fmt.Sprintf("/polls/%v", poll.Id), http.StatusFound)

}
