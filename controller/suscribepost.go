package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationSuscribe"
	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func PostSuscribe(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
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
	typedb := operationSuscribe.GetProvider(os.Getenv("TYPE_DB"))
	err = typedb.CreateSuscribe(claims.Id, event)

	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	return ctx.JSON(http.StatusOK, event)
}
