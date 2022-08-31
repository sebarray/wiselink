package db

import (
	"testing"
	"time"

	"github.com/sebarray/wiselink/model"
	"github.com/stretchr/testify/assert"
)

var Event = model.Event{
	Id:               "",
	DescriptionShort: "description sh",
	Description:      "description long",
	State:            "publicado",
	Organizer:        "gabrielx",
	Place:            "san luiz",
	Date:             time.Now().Add(time.Hour * 72),
	Title:            "hackaton",
}

var eventerror = model.Event{
	Id:               "",
	DescriptionShort: "",
	Description:      "description long",
	State:            "publicado",
	Organizer:        "",
	Place:            "",
	Date:             time.Now(),
	Title:            "",
}

func TestCreateEvent(t *testing.T) {

	evento, err := CreateEvent(Event)
	assert.Equal(t, nil, err)
	_, err = CreateEvent(eventerror)
	assert.NotEqual(t, nil, err)
	_ = DeleteEvent(evento.Id)

}

func TestDeleteEvent(t *testing.T) {
	e, _ := CreateEvent(Event)
	err := DeleteEvent(e.Id)
	assert.Equal(t, nil, err)
	e.Id = ""
	err = DeleteEvent(e.Id)
	assert.NotEqual(t, nil, err)

}
func TestUpdateteEvent(t *testing.T) {
	e, _ := CreateEvent(Event)
	e.Organizer = "gabin belson"
	event, _ := EventUpdate(e)
	assert.Equal(t, e.Organizer, event.Organizer)
}

func TestReadEvents(t *testing.T) {
	var filter model.Filter
	events, _ := ReadEvents(filter)
	assert.NotEqual(t, nil, events)

}
