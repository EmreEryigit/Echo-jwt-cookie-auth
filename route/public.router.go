package route

import (
	"echojwt/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PublicRoute(g *echo.Group) {
	g.POST("/login", controller.Login())
	g.POST("/signup", controller.Signup())
	g.Any("*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "wrong url or method!")
	})
}
