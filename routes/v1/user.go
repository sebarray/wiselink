package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/controller"
)

func Login(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(config.GetCors()))
	e.POST("/login", controller.Login)
	e.POST("/register", controller.PostUser)
}
