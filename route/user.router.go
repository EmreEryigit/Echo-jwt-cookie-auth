package route

import (
	"echojwt/controller"
	"echojwt/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group) {
	g.Use(middleware.Authenticate)
	g.GET("/", controller.WhoAmI())
}
