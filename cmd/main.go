package main

import (
	"eulabs/config"
	Router "eulabs/internal/router"
	"eulabs/pkg/common"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo"
)

func main() {

	log.Info().Msg("Server started")

	e := echo.New()
	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	config.GlobalConfig.LoadVariables()

	Router.RegisterRoutes(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Error().Msg("Error message: " + err.Error())
	}
}
