package db

import (
	"fmt"

	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func UserExist(u model.User) (bool, model.User, error) {
	var user model.User
	var users []model.User
	var err error
	user.Password, err = service.Encryptpswd(user.Password)
	if err != nil {
		return false, model.User{}, err
	}
	query := fmt.Sprintf("SELECT * FROM user where email='%s';", u.Email)
	conn := ConnectionDB()
	defer conn.Close()
	registry, err := conn.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return false, user, err
	}
	for registry.Next() {
		err = registry.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Admin)
		if err != nil {
			fmt.Println(err.Error())
			return false, user, err
		}
		users = append(users, user)
	}
	return true, users[0], err

}
