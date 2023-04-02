package controller

import (
	"example/repository"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllUserList(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to the value stated in the {{ define }} statement in "view/home.html"
	userList := repository.GetAllUserList()
	//return c.Render(http.StatusOK, "home", map[string]interface{}{
	//	"name": "HOME",
	//	"msg":  "Hello, Boatswain!",
	//})
	data := make(map[string]interface{})
	data["user_list"] = userList
	return c.Render(http.StatusOK, "user_list", data)
}
