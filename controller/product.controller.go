package controller

import (
	"context"
	"echojwt/helper"
	"echojwt/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var product model.Product
		if err := c.Bind(&product); err != nil {
			defer cancel()
			return c.JSON(http.StatusInternalServerError, "error extracting json data")
		}
		validationError := validate.Struct(product)
		if validationError != nil {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "invalid content")
		}
		var count int64
		repo.Model(&model.Product{}).Where("title = ?", product.Title).Where("price = ?", product.Price).Count(&count)
		if count > 0 {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "can not create duplicate products")
		}
		currentUser := c.Get("current-user").(*helper.SignedDetails)
		product.UserID = currentUser.UserID
		repo.Save(&product)
		defer cancel()
		return c.JSON(http.StatusCreated, product)
	}
}
