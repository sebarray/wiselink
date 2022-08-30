package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/model"
)

func GetJwtClaims(ctx echo.Context) *model.JwtClaims {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtClaims)
	return claims

}
