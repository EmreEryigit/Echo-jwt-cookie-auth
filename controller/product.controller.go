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

func UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var p model.Product
		productId := c.Param("productId")
		//productId, err := strconv.Atoi(c.Param("productId"))
		/* if err != nil {
			defer cancel()
			return c.String(http.StatusBadRequest, "Not a valid id")
		} */
		err := repo.First(&p, productId).Error
		if err != nil {
			defer cancel()
			return c.String(http.StatusBadRequest, "Not a valid id")
		}
		claims, _ := c.Get("current-user").(*helper.SignedDetails)
		if claims.UserID != p.UserID {
			defer cancel()
			return c.JSON(http.StatusUnauthorized, "You don't own this product")
		}
		c.Bind(&p)
		repo.Save(&p)
		defer cancel()
		return c.JSON(http.StatusOK, p)

	}
}

func DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		productId := c.Param("productId")
		var userID uint
		if err := repo.Model(&model.Product{}).Select("user_id").First(&userID, productId).Error; err != nil {
			defer cancel()
			return c.JSON(http.StatusBadRequest, "this product does not exist")
		}

		claims, _ := c.Get("current-user").(*helper.SignedDetails)
		if claims.UserID != userID {
			defer cancel()
			return c.JSON(http.StatusUnauthorized, "You can not delete this product")
		}
		repo.Delete(&model.Product{}, productId)
		defer cancel()
		return c.NoContent(http.StatusOK)

	}
}

func FetchUserProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		userId := c.Param("userId")
		var products []model.Product
		if err := repo.Model(&model.Product{}).Where("user_id = ?", userId).Find(&products).Error; err != nil {
			defer cancel()
			return c.JSON(http.StatusInternalServerError, "error while fetching products")
		}
		defer cancel()
		return c.JSON(http.StatusOK, products)
	}
}
