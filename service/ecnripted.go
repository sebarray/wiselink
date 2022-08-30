package service

import "golang.org/x/crypto/bcrypt"

func Encryptpswd(pswd string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pswd), costo)
	return string(bytes), err

}
