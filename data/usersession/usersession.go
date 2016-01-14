package usersession

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/atamis/wt3/data/user"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Session struct {
	UserId   user.ID
	LoggedIn bool
}

func (s *Session) LogIn(username string, pass_attempt string) bool {
	log, user := user.AuthenticateUser(username, pass_attempt)

	if !log {
		return false
	}

	s.UserId = user.Id
	s.LoggedIn = true

	return true
}

func (s *Session) Logout() error {
	if s.LoggedIn {
		return errors.New("not logged in")
	}
	s.UserId = 0
	s.LoggedIn = false
	return nil
}

func (s *Session) CurrentUser() (user.User, error) {
	if s.LoggedIn {
		return user.FindUserId(s.UserId)
	}
	return user.User{}, errors.New("user not found")
}

func (s *Session) Debug() string {
	return fmt.Sprintf("%v", s)
}

func MakeSession(r *http.Request) (Session, error) {
	var s = Session{
		UserId:   0,
		LoggedIn: false,
	}

	session, err := store.Get(r, "session-name")

	if err != nil {
		return Session{}, err
	}

	_loggedin := session.Values["loggedin"]

	switch _loggedin.(type) {
	case bool:
		s.LoggedIn = _loggedin.(bool)
	default:
		s.LoggedIn = false
	}

	if s.LoggedIn {
		_id := session.Values["userid"]

		switch _id.(type) {
		case user.ID:
			s.UserId = _id.(user.ID)
		default:
			s.LoggedIn = false
			s.UserId = 0
		}
	}

	return s, nil
}

func (s *Session) Save(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "session-name")

	if err != nil {
		return err
	}

	session.Values["userid"] = s.UserId
	session.Values["loggedin"] = s.LoggedIn

	session.Save(r, w)

	return nil
}
