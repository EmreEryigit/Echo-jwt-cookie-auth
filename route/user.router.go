package route

import (
	"echojwt/controller"
	"echojwt/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(c *echo.Group) {
	c.Use(middleware.Authenticate)
	c.GET("/", controller.GetSelf())
}
