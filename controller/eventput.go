package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationEvent"
	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func PutEvent(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
	if !claims.Admin {
		return echo.ErrUnauthorized
	}
	var event model.Event

	err := ctx.Bind(&event)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	typedb := operationEvent.GetProvider(os.Getenv("TYPE_DB"))
	event, err = typedb.UpdateEvent(event)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, event)
}
