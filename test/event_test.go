package test

import (
	"testing"
	"time"

	"github.com/sebarray/wiselink/config"
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
	config.Loadenv()
	evento, err := SqlEvent.CreateEvent(Event)
	assert.Equal(t, nil, err)
	_, err = SqlEvent.CreateEvent(eventerror)
	assert.NotEqual(t, nil, err)
	_ = SqlEvent.DeleteEvent(evento.Id)

}

func TestDeleteEvent(t *testing.T) {
	config.Loadenv()
	e, _ := SqlEvent.CreateEvent(Event)
	err := SqlEvent.DeleteEvent(e.Id)
	assert.Equal(t, nil, err)
	e.Id = ""
	err = SqlEvent.DeleteEvent(e.Id)
	assert.NotEqual(t, nil, err)

}
func TestUpdateteEvent(t *testing.T) {
	config.Loadenv()
	e, _ := SqlEvent.CreateEvent(Event)
	e.Organizer = "gabin belson"
	event, _ := SqlEvent.UpdateEvent(e)
	assert.Equal(t, e.Organizer, event.Organizer)
}

func TestReadEvents(t *testing.T) {
	config.Loadenv()
	var filter model.Filter
	events, _ := SqlEvent.ReadEvents(filter)
	assert.NotEqual(t, nil, events)

}
