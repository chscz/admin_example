package router

import (
	"example/controller"
	"github.com/labstack/echo"
)

func SetRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", controller.Home)
	e.GET("/about", controller.About)
	e.GET("/user_list", controller.GetAllUserList)
	return e
}
