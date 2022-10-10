package route

import (
	"echojwt/controller"
	"echojwt/middleware"

	"github.com/labstack/echo/v4"
)

func ProtectedRoute(g *echo.Group) {
	g.GET("/products/:userId", controller.FetchUserProducts())
	g.Use(middleware.Authenticate)
	g.GET("/", controller.WhoAmI())
	g.POST("/products", controller.CreateProduct())
	g.PATCH("/products/:productId", controller.UpdateProduct())
	g.DELETE("/products/:productId", controller.DeleteProduct())
}
