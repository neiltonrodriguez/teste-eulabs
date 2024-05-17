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

	return c.JSON(http.StatusOK, products)
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	product, err := ProductModel.GetById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, product)
}

func Create(c echo.Context) error {
	payload := new(domain.Product)
	err := c.Bind(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u, err := ProductModel.Create(c, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, u)
}

func Update(c echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "Erro ao tentar atualizar dados")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}
