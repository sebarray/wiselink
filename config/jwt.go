package config

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/wiselink/model"
)

func Jwt() middleware.JWTConfig {
	cofingJwt := middleware.JWTConfig{
		Claims:     &model.JwtClaims{},
		SigningKey: []byte("secret"),
	}
	return cofingJwt
}
