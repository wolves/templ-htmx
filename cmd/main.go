package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3333", r)
	// e := echo.New()
	//
	// e.Use(middleware.Logger())
	//
	// e.Static("/assets", "dist")
	//
	// e.GET("/", func(c echo.Context) error {
	// 	component := view.Index("Christopher")
	// 	return component.Render(context.Background(), c.Response().Writer)
	// })
	//
	// e.GET("/exp/:data", func(c echo.Context) error {
	// 	data := c.Param("data")
	// 	component := view.Exp(data)
	// 	return component.Render(context.Background(), c.Response().Writer)
	// })
	//
	// e.Logger.Fatal(e.Start(":3333"))
	//
}
