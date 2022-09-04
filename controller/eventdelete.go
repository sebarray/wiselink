package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationEvent"
	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func DeleteEvent(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
	if !claims.Admin {
		return echo.ErrUnauthorized
	}
	var event model.Event
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}
	err = json.Unmarshal(body, &event)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	typedb := operationEvent.GetProvider(os.Getenv("TYPE_DB"))
	err = typedb.DeleteEvent(event.Id)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	return ctx.JSON(http.StatusOK, event)
}
