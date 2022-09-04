package operationUser

import "github.com/sebarray/wiselink/model"

type UserSql struct{}

type ISuscribe interface {
	Login(user model.User) (string, error)
	CreateUser(user model.User) (model.User, error)
}

var providers = map[string]UserSql{}

func init() {
	providers["mysql"] = UserSql{}
}

func GetProvider(typeDb string) UserSql {
	return providers[typeDb]
}
