package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSuscribe(t *testing.T) {
	u, _ := CreateUser(User)
	event, _ := CreateEvent(Event)
	err := CreateSuscribe(u.Id, event)
	assert.Equal(t, nil, err)

	DeleteEvent(event.Id)
	DeleteUser(u.Id)

}

func TestReadSuscribe(t *testing.T) {
	u, _ := CreateUser(User)
	event, _ := CreateEvent(Event)
	_ = CreateSuscribe(u.Id, event)
	_, err := ReadSuscribe("", u.Id)
	assert.Equal(t, nil, err)
	DeleteEvent(event.Id)
	DeleteUser(u.Id)

}
