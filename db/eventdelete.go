package db

import "fmt"

func DeleteEvent(id string) error {
	querysql := fmt.Sprintf("DELETE FROM event WHERE id='%s' ;", id)
	conn := ConnectionDB()
	defer conn.Close()
	query, err := conn.Prepare(querysql)
	if err != nil {
		fmt.Println(err.Error(), " delete 1")
		return err
	}
	_, err = query.Exec()
	if err != nil {
		fmt.Println(err.Error(), " delete 2")
		return err
	}

	return nil
}
