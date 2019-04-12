package api

import (
	"github.com/labstack/echo"
)

func GetProvince() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		ctx.Response().Write([]byte(province))
		return nil
	}
}

func GetCity() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Write([]byte(city))
		return nil
	}
}
