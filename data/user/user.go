package user

import "errors"

type User struct {
	Id       int
	Username string
	Password string
}

var andrew = User{
	Id:       1,
	Username: "andrew",
	Password: "password",
}

var notfound = errors.New("user not found")

func FindUserName(username string) (User, error) {
	if username == "andrew" {
		return andrew, nil
	}

	return User{}, notfound
}

func FindUserId(id int) (User, error) {
	if id == 1 {
		return andrew, nil
	}

	return User{}, notfound
}

func AuthenticateUser(username string, pass_attempt string) (bool, User) {
	var u, err = FindUserName(username)

	if err != nil {
		return false, User{}
	}

	return u.Password == pass_attempt, u

}
