package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationUser"
	"github.com/sebarray/wiselink/model"
)

func Login(ctx echo.Context) error {
	var user model.User
	var t string

	err := ctx.Bind(&user)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}

	typedb := operationUser.GetProvider(os.Getenv("TYPE_DB"))
	t, err = typedb.Login(user)

	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)

	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}
