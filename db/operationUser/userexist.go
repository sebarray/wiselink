package operationUser

import (
	"fmt"

	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func UserExist(u model.User) (model.User, error) {
	var user model.User
	var users []model.User
	var err error
	user.Password, err = service.Encryptpswd(user.Password)
	if err != nil {
		return model.User{}, err
	}
	query := fmt.Sprintf("SELECT * FROM user where email='%s';", u.Email)
	conn := db.ConnectionDB()
	defer conn.Close()
	registry, err := conn.Query(query)
	if err != nil {
		return user, err
	}
	for registry.Next() {
		err = registry.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Admin)
		if err != nil {

			return user, err
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return users[0], fmt.Errorf("hubo un error en el login")
	}

	return users[0], err

}
