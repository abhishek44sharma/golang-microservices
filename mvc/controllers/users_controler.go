package controllers

import (
	"golang-microservices/mvc/services"
	"strconv"
	"net/http"
	"encoding/json"
	"golang-microservices/mvc/utils"
	"fmt"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:	fmt.Sprintf("Invalid User Id Sent"),
			StatusCode:	http.StatusBadRequest,
			Code:		"Bad Request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}