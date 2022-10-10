package controller

import (
	"context"
	"echojwt/database"
	"echojwt/helper"
	"echojwt/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var repo *gorm.DB = database.ConnectDB()
var validate = validator.New()
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		log.Panic(err)
	}
	return string(hash)
}
func VerifyPassword() {}

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var user model.User
		if err := c.Bind(&user); err != nil {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "invalid request")
		}
		validationError := validate.Struct(user)
		if validationError != nil {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "invalid request")
		}
		var count int64
		repo.Model(&model.User{}).Where("name = ?", user.Email).Count(&count)
		defer cancel()
		if count > 0 {
			return c.JSON(http.StatusConflict, "email already taken")
		}
		hashedPassword := HashPassword(*user.Password)
		user.Password = &hashedPassword
		repo.Save(&user)
		jwtToken, err := helper.GenerateJWT(fmt.Sprint(user.ID), *user.Name, *user.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error while generating jwt token")

		}

		session, _ := Store.Get(c.Request(), "auth-session")
		session.Values["auth"] = jwtToken
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error while generating jwt token")
			return err

		}

		c.JSON(http.StatusOK, user.Email)
		return err
	}

}

/* func Login() *echo.Echo {} */

func GetSelf() echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("current-user")
		return c.JSON(http.StatusOK, claims)
	}
}
