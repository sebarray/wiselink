package controller

import (
	"os"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationUser"
	"github.com/sebarray/wiselink/model"
)

func PostUser(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	typedb := operationUser.GetProvider(os.Getenv("TYPE_DB"))
	user, err = typedb.CreateUser(user)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	user.Password = ""
	return ctx.JSON(http.StatusOK, user)
}
