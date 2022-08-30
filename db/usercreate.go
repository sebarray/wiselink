package db

import (
	"github.com/google/uuid"
	"github.com/sebarray/wiselink/model"
)

func CreateUser(user model.User) (model.User, error) {
	user.Id = uuid.New().String()
	conn := ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(QueryCreateUser(user))
	if err != nil {
		return user, err
	}
	_, err = insert.Exec()
	if err != nil {
		return user, err
	}
	return user, nil

}
