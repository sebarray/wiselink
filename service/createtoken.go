package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/model"
)

func CreateToken(user model.User) (string, error) {
	config.Loadenv()
	claims := &model.JwtClaims{
		Name:  user.Name,
		Id:    user.Id,
		Admin: user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return t, err
	}
	return t, nil
}
