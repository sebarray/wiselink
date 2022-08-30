package model

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	Id    string `json:"id"`
	jwt.StandardClaims
}
