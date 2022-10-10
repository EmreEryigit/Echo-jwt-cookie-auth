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
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Panic(err)
	}
	return string(hash)
}
func VerifyPassword(providedPassword string, storedHash string) bool {
	fmt.Println("before")
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword))
	fmt.Println("after")
	valid := true
	if err != nil {
		valid = false
	}
	return valid
}

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
		repo.Model(&model.User{}).Where("email = ?", user.Email).Count(&count)

		if count > 0 {
			defer cancel()
			return c.JSON(http.StatusConflict, "email already taken")
		}
		hashedPassword := HashPassword(*user.Password)
		user.Password = &hashedPassword
		repo.Save(&user)
		jwtToken, err := helper.GenerateJWT(fmt.Sprint(user.ID), *user.Name, *user.Email)
		if err != nil {
			defer cancel()
			return c.JSON(http.StatusInternalServerError, "error while generating jwt token")
		}
		session, _ := Store.Get(c.Request(), "auth-session")
		session.Values["auth"] = jwtToken
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			defer cancel()
			c.JSON(http.StatusInternalServerError, "error while generating jwt token")
			return err
		}
		c.JSON(http.StatusOK, user.Email)
		defer cancel()
		return err
	}

}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var user model.User
		var foundUser model.User
		if err := c.Bind(&user); err != nil {
			defer cancel()
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		result := repo.Where("email = ?", user.Email).First(&foundUser)
		if result.Error != nil {
			defer cancel()
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		if foundUser.Email == nil {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "user not found")
		}
		isValid := VerifyPassword(*user.Password, *foundUser.Password)
		if !isValid {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "invalid email or password")
		}
		token, err := helper.GenerateJWT(fmt.Sprint(foundUser.ID), *foundUser.Name, *foundUser.Email)
		if err != nil {
			defer cancel()
			return c.JSON(http.StatusInternalServerError, "error while generating token")
		}
		session, _ := Store.Get(c.Request(), "auth-session")
		session.Values["auth"] = token
		err1 := session.Save(c.Request(), c.Response())
		if err1 != nil {
			defer cancel()
			return c.JSON(http.StatusInternalServerError, "could not save the cookie")
		}
		c.JSON(http.StatusOK, foundUser.Email)
		defer cancel()
		return err
	}
}

func WhoAmI() echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("current-user")
		return c.JSON(http.StatusOK, claims)
	}
}
