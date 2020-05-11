package oauth

import (
	"fmt"
	"golang-microservices/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = ""
)

var (
	users = map[string]*User{
		"Abhishek": {Id: 123, Username: "Abhishek"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprint("no user found with given parameters"))
	}
	return user, nil
}
