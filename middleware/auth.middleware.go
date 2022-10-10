package middleware

import (
	"echojwt/controller"
	"echojwt/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := controller.Store.Get(c.Request(), "auth-session")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "no cookie provided")
		}
		strJwt := fmt.Sprintf("%v", session.Values["auth"])
		claims, msg := helper.ValidateToken(strJwt)
		fmt.Println(claims)
		if msg != "" {
			return c.JSON(http.StatusUnauthorized, "invalid token")
		}
		c.Set("current-user", claims)
		return next(c)
	}
}
