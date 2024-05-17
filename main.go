package main

import (
	"eulabs/common"
	"eulabs/domain"
	ProductRouter "eulabs/internal/product/router"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo"
)

func main() {

	log.Info().Msg("Server started")

	e := echo.New()
	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	domain.GlobalConfig.LoadVariables()

	ProductRouter.RegisterRoutes(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Error().Msg("Error message: " + err.Error())
	}
}
