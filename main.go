package main

import (
	"eulabs/domain"
	ProductRouter "eulabs/internal/product/router"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	fmt.Println("running application..")

	e := echo.New()
	domain.GlobalConfig.LoadVariables()

	ProductRouter.RegisterRoutes(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
