package main

import (
	"eulabs/config"
	Router "eulabs/internal/router"
	"eulabs/pkg/common"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}


func main() {

	log.Info().Msg("Server started")

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	
	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	config.GlobalConfig.LoadVariables()

	Router.RegisterRoutes(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Error().Msg("Error message: " + err.Error())
	}
}


func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
