package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAuthRoutes() echo.HandlerFunc {
	a := map[string]interface{}{
		"/form/advanced-form": map[string]interface{}{
			"authority": []string{"admin", "user"},
		},
	}
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, a)
	}
}
