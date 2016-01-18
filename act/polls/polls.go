package polls

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/atamis/wt3/data/poll"
	"github.com/gorilla/mux"
)

func Poll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	p, err := poll.Find(id)
	p.LoadAnswers()

	if err != nil {
		http.Error(w, "Poll not found "+err.Error(), 404)
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", p.Collated())))
}
