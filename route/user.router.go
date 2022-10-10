package route

import (
	"echojwt/controller"
	"echojwt/middleware"
	"fmt"

	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group) {
	fmt.Println("route2 fired")
	g.Use(middleware.Authenticate)
	g.GET("/", controller.WhoAmI())
}
