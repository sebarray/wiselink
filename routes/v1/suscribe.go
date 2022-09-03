package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/controller"
)

func Suscribe(s *echo.Group) {
	s.Use(middleware.JWTWithConfig(config.Jwt()), middleware.CORSWithConfig(config.GetCors()))
	s.GET("", controller.GetSuscribe)
	s.POST("", controller.PostSuscribe)
}
