package main

import (
	"echojwt/route"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	e := echo.New()
	authG := e.Group("/auth")
	userG := e.Group("/user")
	route.AuthRoute(authG)
	route.UserRoutes(userG)
	e.Logger.Fatal(e.Start(":" + port))
}
