package handler

import (
	"eulabs/internal/domain"
	Model "eulabs/internal/model"
	"eulabs/pkg/common"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func GetAll(c echo.Context) error {
	products, err := Model.GetAll(c.Request().Context())
	if err != nil {
		log.Error().Msg("Error in GetAll of product: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	common.Logger.LogInfo().Msg("GetAll successfully")
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: len(products),
		},
		Data: products,
	})
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	product, err := Model.GetById(c.Request().Context(), id)
	if err != nil {
		log.Error().Msg("Error in GetById of product: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Info().Msg("GetById successfully")
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: product,
	})
}

func Create(c echo.Context) error {
	payload := new(domain.ProductInputDTO)
	err := c.Bind(payload)
	if err != nil {
		log.Error().Msg("bind error in struct: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := Model.Create(c, payload)
	if err != nil {
		log.Error().Msg("Error in Create: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Create with successfully")
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: product,
	})
}

func Update(c echo.Context) error {
	id := c.Param("id")
	payload := new(domain.ProductInputDTO)
	err := c.Bind(payload)
	if err != nil {
		log.Error().Msg("bind error in update: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = Model.Update(c.Request().Context(), payload, id)
	if err != nil {
		log.Error().Msg("Error in Update: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Create with successfully")
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "success",
	})
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	err := Model.Delete(c.Request().Context(), id)
	if err != nil {
		log.Error().Msg("Error in Delete: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Deleted with successfully")
	return c.JSON(http.StatusOK, domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "success",
	})
}
