package test

import (
	"testing"

	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/db/operationEvent"
	"github.com/sebarray/wiselink/db/operationSuscribe"
	"github.com/sebarray/wiselink/db/operationUser"
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

var SqlEvent = operationEvent.EventSql{}
var SqlSuscribe = operationSuscribe.SuscribeSql{}
var sqluser = operationUser.UserSql{}

func TestCreateSuscribe(t *testing.T) {
	config.Loadenv()
	u, _ := sqluser.CreateUser(User)
	event, _ := SqlEvent.CreateEvent(Event)
	err := SqlSuscribe.CreateSuscribe(u.Id, event)
	assert.Equal(t, nil, err)

	SqlEvent.DeleteEvent(event.Id)
	operationUser.DeleteUser(u.Id)

}

func TestReadSuscribe(t *testing.T) {
	config.Loadenv()
	u, _ := sqluser.CreateUser(User)
	event, _ := SqlEvent.CreateEvent(Event)
	_ = SqlSuscribe.CreateSuscribe(u.Id, event)
	_, err := SqlSuscribe.ReadSuscribe("", u.Id)
	assert.Equal(t, nil, err)
	SqlEvent.DeleteEvent(event.Id)
	operationUser.DeleteUser(u.Id)

}
