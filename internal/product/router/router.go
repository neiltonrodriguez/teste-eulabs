package router

import (
	"eulabs/internal/product/handler"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo) {

	g := e.Group("/product")

	g.GET("/", handler.GetAll).Name = "get-all"

	g.GET("/:id", handler.GetById).Name = "get-by-id"

	g.POST("/", handler.Create).Name = "create"

	g.PUT("/:id", handler.Update).Name = "update"

	g.DELETE("/:id", handler.Delete).Name = "delete"
}
