package test

import (
	"testing"

	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/db/operationUser"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	config.Loadenv()
	_, err := sqluser.CreateUser(userErr)
	assert.NotEqual(t, nil, err)

	userdelete, err := sqluser.CreateUser(User)
	assert.Equal(t, nil, err)
	operationUser.DeleteUser(userdelete.Id)
}

func TestLogin(t *testing.T) {
	config.Loadenv()
	User.Email = "sebastian.demello@wiselink.com"
	_, err := sqluser.Login(User)
	assert.Equal(t, nil, err)
}
