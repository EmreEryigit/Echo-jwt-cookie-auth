package route

import (
	"echojwt/controller"
	"echojwt/middleware"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(g *echo.Group) {
	fmt.Println("route fired")
	g.GET("/", controller.GetAllProducts())
	g.GET("/:userId", controller.FetchUserProducts())
	// authenticated routes
	g.Use(middleware.Authenticate)
	g.POST("/", controller.CreateProduct())
	g.PATCH("/:productId", controller.UpdateProduct())
	g.DELETE("/:productId", controller.DeleteProduct())
}
