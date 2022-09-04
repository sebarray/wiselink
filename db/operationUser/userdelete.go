package operationUser

import (
	"fmt"

	"github.com/sebarray/wiselink/db"
)

func DeleteUser(id string) error {
	err := deleteSuscribeUser(id)
	if err != nil {
		return err
	}
	querysql := fmt.Sprintf("DELETE FROM user WHERE id='%s' ;", id)
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

func deleteSuscribeUser(id string) error {
	querysql := fmt.Sprintf("DELETE FROM suscribe WHERE iduser='%s' ;", id)
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
