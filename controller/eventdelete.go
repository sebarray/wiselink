package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db"
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
		ctx.Error(err)
	}
	err = json.Unmarshal(body, &event)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	err = db.DeleteEvent(event.Id)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}

	return ctx.JSON(http.StatusOK, event)
}
