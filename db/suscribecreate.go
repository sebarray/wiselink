package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sebarray/wiselink/model"
)

func CreateSuscribe(userId string, event model.Event) (bool, error) {

	if !time.Now().Before(event.Date) {
		fmt.Println("ODIO MI VIDA")
		return false, nil
	}

	id := uuid.New().String()
	fmt.Println(QueryCreateSuscribe(id, userId, event.Id))
	conn := ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(QueryCreateSuscribe(id, userId, event.Id))
	if err != nil {
		return false, err
	}
	_, err = insert.Exec()
	if err != nil {
		return false, err
	}

	return true, nil
}
