package domain

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		365341: &User{365341, "ABHISHEK", "SHARMA", "ab258841@wipro.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{fmt.Sprintf("User %v not found", userId), http.StatusNotFound, "Not Found"}
}