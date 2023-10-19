package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wolves/templ-htmx/pkg/view"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	// e.Static("/assets", "dist")

	e.GET("/", func(c echo.Context) error {
		component := view.Index("Christopher")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":3333"))

}
