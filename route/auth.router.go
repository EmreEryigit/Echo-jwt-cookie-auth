package route

import (
	"echojwt/controller"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group) {
	g.POST("/login", controller.Login())
	g.POST("/signup", controller.Signup())
	fmt.Println("route fired1")
	g.Any("*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "wrong url or method!")
	})
}
