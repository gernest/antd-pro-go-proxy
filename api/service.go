package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ApiError struct {
	Timestamp int64  `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

func Service() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	api := e.Group("/api")
	api.GET("/currentUser", CurrentUser())
	api.GET("/users", Users())
	api.POST("/login/account", Login())
	api.POST("/login/register", Register())
	api.GET("/500", RenderError(ApiError{
		Timestamp: 1513932555104,
		Status:    500,
		Error:     "error",
		Message:   "error",
		Path:      "/base/category/list",
	}))
	api.GET("/404", RenderError(ApiError{
		Timestamp: 1513932643431,
		Status:    404,
		Error:     "Not Found",
		Message:   "No message available",
		Path:      "/base/category/list/2121212",
	}))
	api.GET("/403", RenderError(ApiError{
		Timestamp: 1513932555104,
		Status:    403,
		Error:     "Unauthorized",
		Message:   "Unauthorized",
		Path:      "/base/category/list",
	}))
	api.GET("/401", RenderError(ApiError{
		Timestamp: 1513932555104,
		Status:    401,
		Error:     "Unauthorized",
		Message:   "Unauthorized",
		Path:      "/base/category/list",
	}))
	api.GET("/fake_list", GetFakeList())
	api.POST("/fake_list", PostFakeList())
	api.GET("/captcha", GetFakeCaptcha())
	api.POST("/forms", PostForms())
	api.GET("/activities", GetActivity())
	api.GET("/tags", GetTags())
	api.GET("/project/notice", GetNotice())
	api.GET("/notices", GetNotices())
	api.GET("/profile/advanced", GetProfileAdvancedData())
	api.GET("/profile/basic", GetProfileBasic())
	api.GET("/fake_chart_data", GetChartData())
	api.GET("/auth_routes", GetAuthRoutes())
	api.GET("/rule", GetRule())
	api.POST("/rule", PostRule())
	api.GET("/api/geographic/city/:province", GetCity())
	api.GET("/geographic/province", GetProvince())
	return e
}

func RenderError(e ApiError) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(e.Status, e)
	}
}
