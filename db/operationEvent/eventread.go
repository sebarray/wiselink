package operationEvent

import (
	"fmt"

	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/model"
)

func (e EventSql) ReadEvents(filter model.Filter) ([]model.Event, error) {

	db.Checkconnection()
	var events []model.Event
	var event model.Event
	conn := db.ConnectionDB()
	defer conn.Close()
	registry, err := conn.Query(db.QueryReadEvents(filter))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for registry.Next() {
		err := registry.Scan(&event.Id, &event.Description,
			&event.DescriptionShort, &event.State, &event.Organizer,
			&event.Place, &event.Date, &event.Title)
		if err != nil {
			fmt.Println(err.Error())
		}
		events = append(events, event)
	}

	return events, nil
}

func ReadEvent(id string) (model.Event, error) {
	var events []model.Event
	var event model.Event
	query := fmt.Sprintf("SELECT * FROM event where id='%s'", id)
	registry, err := db.ConnectionDB().Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return event, err
	}
	for registry.Next() {
		err := registry.Scan(&event.Id, &event.Description,
			&event.DescriptionShort, &event.State, &event.Organizer,
			&event.Place, &event.Date, &event.Title)
		if err != nil {
			fmt.Println(err.Error())
		}

		events = append(events, event)
	}
	return events[0], nil
}
