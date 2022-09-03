package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetCors() middleware.CORSConfig {

	return middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}

}
