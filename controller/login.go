package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/model"
)

func Login(ctx echo.Context) error {
	var user model.User
	var t string
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.Error(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	t, err = db.Login(user)

	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}
