package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db/operationSuscribe"
	"github.com/sebarray/wiselink/service"
)

func GetSuscribe(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
	index := map[string]string{"completados": "<", "activos": ">"}
	filter := ctx.QueryParam("filtro")
	typedb := operationSuscribe.GetProvider(os.Getenv("TYPE_DB"))
	suscribes, err := typedb.ReadSuscribe(index[filter], claims.Id)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, suscribes)
}
