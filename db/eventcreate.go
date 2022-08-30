package db

import (
	"github.com/google/uuid"
	"github.com/sebarray/wiselink/model"
)

func CreateEvent(event model.Event) (model.Event, error) {
	event.Id = uuid.New().String()
	conn := ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(QueryCreateEvent(event))
	if err != nil {
		return event, err
	}
	_, err = insert.Exec()

	if err != nil {
		return event, err
	}
	return event, nil

}
