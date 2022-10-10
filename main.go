package main

import (
	"echojwt/database"
	"echojwt/middleware"
	"echojwt/route"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()
	port := os.Getenv("PORT")
	e := echo.New()
	e.Use(middleware.CurrentUser)
	fmt.Println("main fired")
	authG := e.Group("/auth")
	userG := e.Group("/users")
	productG := e.Group("/products")
	route.ProductRoutes(productG)
	route.AuthRoutes(authG)
	route.UserRoutes(userG)
	e.Any("*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "wrong url or method!")
	})

	e.Logger.Fatal(e.Start(":" + port))
}
