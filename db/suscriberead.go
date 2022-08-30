package db

import (
	"fmt"

	"github.com/sebarray/wiselink/model"
)

func ReadSuscribe(filter, id string) ([]model.Event, error) {
	Checkconnection()
	var suscribes []model.Event
	var suscribe model.Event
	conn := ConnectionDB()
	defer conn.Close()
	registry, err := conn.Query(QueryReadSuscribe(id, filter))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for registry.Next() {
		err := registry.Scan(&suscribe.Id, &suscribe.Description,
			&suscribe.DescriptionShort, &suscribe.State, &suscribe.Organizer,
			&suscribe.Place, &suscribe.Date, &suscribe.Title)
		if err != nil {
			fmt.Println(err.Error())
		}
		suscribes = append(suscribes, suscribe)
	}

	return suscribes, nil
}
