package oauth

import (
	"github.com/evertonkopec/golang-microservices-main/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"kopec": {Id: 123, Username: "kopec"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError("no user found with given parameters")
	}
	return user, nil
}
