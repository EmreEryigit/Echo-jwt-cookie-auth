package route

import (
	"echojwt/controller"

	"github.com/labstack/echo/v4"
)

func PublicRoute(g *echo.Group) {
	g.POST("/login", controller.Login())
	g.POST("/signup", controller.Signup())
}
