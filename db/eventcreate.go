package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sebarray/wiselink/model"
)

func CreateEvent(event model.Event) (model.Event, error) {

	event.Id = uuid.New().String()
	err := create(event)
	if err != nil {
		return event, err
	}
	conn := ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(QueryCreateEvent(event))
	if err != nil {
		return event, err
	}
	result, err := insert.Exec()

	if err != nil {
		return event, err
	}
	rowAf, err := result.RowsAffected()
	if err != nil {
		return event, err
	}
	if rowAf == 0 {
		return event, fmt.Errorf("se ocasiono un error resvise su request")
	}

	return event, nil

}

func create(event model.Event) error {
	if event.Id == "" && event.Description == "" && event.State == "" {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}
	if event.Date.String() == "" && event.DescriptionShort == "" && event.Place == "" {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}
	if event.Date.String() == "" && event.DescriptionShort == "" && event.Place == "" {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}
	if event.Title == "" && event.Organizer == "" {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}
	if event.State != "publicado" && event.State != "borrador" {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}

	return nil
}
