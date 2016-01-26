package home

import (
	"net/http"

	us "github.com/atamis/wt3/data/usersession"
	"github.com/atamis/wt3/met/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	sess, err := us.MakeSession(r)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sess.Save(w, r)

	err = template.RenderTemplate(w, "index.tmpl", sess.Locals())

	if err != nil {
		panic(err)
	}
}
