package operationUser

import (
	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
	"golang.org/x/crypto/bcrypt"
)

func (us UserSql) Login(user model.User) (string, error) {

	u, err := UserExist(user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	t, err := service.CreateToken(u)

	if err != nil {
		return "", err
	}

	return t, nil
}
