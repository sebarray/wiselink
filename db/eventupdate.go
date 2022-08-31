package db

import (
	"fmt"

	"github.com/sebarray/wiselink/model"
)

func update(event model.Event) error {

	query := QueryUpdateEvents(event)
	conn := ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(query)
	if err != nil {
		return err
	}
	result, err := insert.Exec()
	if err != nil {
		return err
	}
	rowAf, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAf == 0 {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}

	return nil
}

func EventUpdate(event model.Event) (model.Event, error) {
	eventf, err := ReadEvent(event.Id)
	if err != nil {
		return event, err
	}
	if event.Description == "" {
		event.Description = eventf.Description
	}
	if event.DescriptionShort == "" {
		event.DescriptionShort = eventf.DescriptionShort
	}
	if event.State == "" {
		event.State = eventf.State
	}
	if event.Organizer == "" {
		event.Organizer = eventf.Organizer
	}
	if event.Place == "" {
		event.Place = eventf.Place
	}
	if event.Date.Format("2006-01-02 15:04:05") == "" {
		event.Date = eventf.Date
	}
	if event.Title == "" {
		event.Title = eventf.Title
	}
	err = update(event)
	return event, err
}
