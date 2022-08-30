package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/controller"
)

func Login(e *echo.Echo) {
	e.POST("/login", controller.Login)
	e.POST("/register", controller.PostUser)
}
