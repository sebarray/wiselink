package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationEvent"
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
	typedb := operationEvent.GetProvider(os.Getenv("TYPE_DB"))
	events, err := typedb.ReadEvents(filter)

	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, events)

}
