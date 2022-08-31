package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func GetEvents(ctx echo.Context) error {
	var filter model.Filter
	claims := service.GetJwtClaims(ctx)
	filter.Title = ctx.QueryParam("titulo")
	filter.Date = ctx.QueryParam("fecha")
	filter.State = ctx.QueryParam("estado")
	if !claims.Admin {
		filter.State = "publicado"
	}
	events, err := db.ReadEvents(filter)

	if err != nil {
		ctx.Error(err)
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	return ctx.JSON(http.StatusOK, events)

}
