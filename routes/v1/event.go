package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/controller"
)

func Event(r *echo.Group) {
	r.Use(middleware.JWTWithConfig(config.Jwt()), middleware.CORSWithConfig(config.GetCors()))
	r.GET("", controller.GetEvents)
	r.POST("", controller.PostEvent)
	r.PUT("", controller.PutEvent)
	r.DELETE("", controller.DeleteEvent)
}
