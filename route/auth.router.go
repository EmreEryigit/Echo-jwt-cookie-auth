package route

import (
	"echojwt/controller"

	"github.com/labstack/echo/v4"
)

func AuthRoute(c *echo.Group) {

	c.POST("/signup", controller.Signup())
}
