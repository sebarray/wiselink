package db

import (
	"testing"

	"github.com/sebarray/wiselink/model"
	"github.com/stretchr/testify/assert"
)

var userErr = model.User{
	Email:    "",
	Password: "montevideo724",
	Admin:    true,
}
var User = model.User{
	Email:    "sebastian.demello@wiselink.com",
	Password: "montevideo724",
	Admin:    true,
}

func TestUserCreate(t *testing.T) {
	_, err := CreateUser(userErr)
	assert.NotEqual(t, nil, err)

	userdelete, err := CreateUser(User)
	assert.Equal(t, nil, err)
	DeleteUser(userdelete.Id)
}

func TestLogin(t *testing.T) {
	User.Email = "sebastian.demello@wiselink.com"
	_, err := Login(User)
	assert.Equal(t, nil, err)
}
