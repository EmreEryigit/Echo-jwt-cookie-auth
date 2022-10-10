package main

import (
	"echojwt/database"
	"echojwt/middleware"
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
	database.ConnectDB()
	port := os.Getenv("PORT")
	e := echo.New()
	e.Use(middleware.CurrentUser)
	authG := e.Group("/auth")
	userG := e.Group("/user")
	route.PublicRoute(authG)
	route.ProtectedRoute(userG)
	e.Logger.Fatal(e.Start(":" + port))
}
