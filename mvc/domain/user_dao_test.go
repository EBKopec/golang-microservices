package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T){
	// Initialization:

	// Execution:
	user, err := UserDao.GetUser(0)

	// Validation:
	assert.Nil(t, user, "we were expecting a user with id 0.")
	assert.NotNil(t, err, "we were expecting an error when user id is 0.")

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user: 0 does not exists!", err.Message)
}

func TestGetUserNoError(t *testing.T){

	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Everton", user.FirstName)
	assert.EqualValues(t, "Kopec", user.LastName)
	assert.EqualValues(t, "evertonkopec@hotmail.com", user.Email)
}