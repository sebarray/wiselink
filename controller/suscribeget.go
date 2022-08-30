package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/wiselink/db"
	"github.com/sebarray/wiselink/service"
)

func GetSuscribe(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
	index := map[string]string{"completados": "<", "activos": ">"}
	filter := ctx.QueryParam("filtro")
	suscribes, err := db.ReadSuscribe(index[filter], claims.Id)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	return ctx.JSON(http.StatusOK, suscribes)
}
