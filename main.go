package main

import (
	ProductRouter "eulabs/internal/product/router"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("running application..")

	e := echo.New()

	ProductRouter.RegisterRoutes(e)
	e.Start(":8000")
}
