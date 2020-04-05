package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUser(t *testing.T) {
	usr, err := GetUser(0)
	assert.Nil(t, usr, "We are not accepting user with ID 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not Found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	usr, err := GetUser(365341)
	assert.Nil(t, err, "Success Error stuct should be nil")
	assert.EqualValues(t, usr.Id, 365341)
	assert.EqualValues(t, usr.Firstname, "ABHISHEK")
	assert.EqualValues(t, usr.Lastname, "SHARMA")
	assert.EqualValues(t, usr.Email, "ab258841@wipro.com")
}
