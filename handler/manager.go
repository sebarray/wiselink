package handler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sebarray/wiselink/routes/v1"
)

func Manager() {

	e := echo.New()
	e.Use(middleware.Logger())
	r := e.Group("/event")
	r.Use(middleware.Logger())
	routesSuscribe := e.Group("/suscribe")
	routesSuscribe.Use(middleware.Logger())
	routes.Event(r)
	routes.Suscribe(routesSuscribe)
	routes.Login(e)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3282"
	}
	e.Start(":" + PORT)
}
