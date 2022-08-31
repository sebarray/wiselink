package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sebarray/wiselink/model"
)

func CreateSuscribe(userId string, event model.Event) error {
	var err error
	if !time.Now().Before(event.Date) {
		return fmt.Errorf("no es posible anotarse a eventos completados")
	}

	id := uuid.New().String()
	fmt.Println(QueryCreateSuscribe(id, userId, event.Id))
	conn := ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(QueryCreateSuscribe(id, userId, event.Id))
	if err != nil {
		return err
	}
	_, err = insert.Exec()
	if err != nil {
		return err
	}

	return nil
}
