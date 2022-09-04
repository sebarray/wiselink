package operationEvent

import (
	"fmt"

	"github.com/sebarray/wiselink/db"
)

func (e EventSql) DeleteEvent(id string) error {
	err := deleteSuscribe(id)
	if err != nil {
		return err
	}
	querysql := fmt.Sprintf("DELETE FROM event WHERE id='%s' ;", id)
	conn := db.ConnectionDB()
	defer conn.Close()
	query, err := conn.Prepare(querysql)
	if err != nil {
		return err
	}
	result, err := query.Exec()
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

func deleteSuscribe(id string) error {
	querysql := fmt.Sprintf("DELETE FROM suscribe WHERE idevent='%s' ;", id)
	conn := db.ConnectionDB()
	defer conn.Close()
	query, err := conn.Prepare(querysql)
	if err != nil {
		return err
	}
	_, err = query.Exec()
	if err != nil {
		return err
	}

	return nil
}
