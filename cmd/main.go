package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wolves/templ-htmx/pkg/view"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(view.Index("Christopher")).ServeHTTP(w, r)
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "dist"))

	FileServer(r, "/assets", filesDir)

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

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
