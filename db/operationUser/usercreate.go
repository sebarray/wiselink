package operationUser

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/model"
)

func (us UserSql) CreateUser(user model.User) (model.User, error) {
	if user.Email == "" || user.Password == "" {
		return user, fmt.Errorf("ocurrio un erro revise un request")

	}
	user.Id = uuid.New().String()
	conn := db.ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(db.QueryCreateUser(user))
	if err != nil {
		return user, err
	}
	_, err = insert.Exec()
	if err != nil {
		return user, err
	}
	return user, nil

}
