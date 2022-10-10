package route

import (
	"echojwt/controller"
	"echojwt/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group) {
	g.POST("/login", controller.Login())
	g.POST("/signup", controller.Signup())
	g.Use(middleware.Authenticate)
	g.POST("/logout", controller.Logout())
	g.Any("*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "wrong url or method!")
	})
}
