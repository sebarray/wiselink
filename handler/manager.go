package handler

import (
	"os"

	"github.com/labstack/echo/v4"

	"github.com/sebarray/wiselink/routes/v1"
)

func Manager() {

	e := echo.New()
	r := e.Group("/event")
	routesSuscribe := e.Group("/suscribe")
	routes.Event(r)
	routes.Suscribe(routesSuscribe)
	routes.Login(e)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3282"
	}
	e.Start(":" + PORT)
}
