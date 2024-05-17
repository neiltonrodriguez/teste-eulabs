package handler

import (
	"eulabs/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "get all")
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	// team := c.QueryParam("team")
	return c.String(http.StatusOK, fmt.Sprintf("get by id, id: %s", id))
}

func Create(c echo.Context) error {
	u := new(domain.Product)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}
