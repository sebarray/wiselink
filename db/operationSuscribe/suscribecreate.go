package operationSuscribe

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/model"
)

func (s SuscribeSql) CreateSuscribe(userId string, event model.Event) error {
	var err error
	if !time.Now().Before(event.Date) {
		return fmt.Errorf("no es posible anotarse a eventos completados")
	}

	id := uuid.New().String()

	conn := db.ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(db.QueryCreateSuscribe(id, userId, event.Id))
	if err != nil {
		return err
	}
	_, err = insert.Exec()
	if err != nil {
		return err
	}

	return nil
}
