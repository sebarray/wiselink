package controller

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationUser"
	"github.com/sebarray/wiselink/model"
)

func PostUser(ctx echo.Context) error {
	var user model.User
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.Error(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	typedb := operationUser.GetProvider(os.Getenv("TYPE_DB"))
	user, err = typedb.CreateUser(user)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	user.Password = ""
	return ctx.JSON(http.StatusOK, user)
}
