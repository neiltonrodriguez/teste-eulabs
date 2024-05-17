package handler

import (
	"eulabs/domain"
	ProductModel "eulabs/models/product"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	products, err := ProductModel.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: len(products),
		},
		Data: products,
	})
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	product, err := ProductModel.GetById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: product,
	})
}

func Create(c echo.Context) error {
	payload := new(domain.Product)
	err := c.Bind(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := ProductModel.Create(c, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: product,
	})
}

func Update(c echo.Context) error {
	id := c.Param("id")
	payload := new(domain.Product)
	err := c.Bind(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = ProductModel.Update(c.Request().Context(), payload, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "success",
	})
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	err := ProductModel.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "success",
	})
}
