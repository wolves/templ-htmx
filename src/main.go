package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/wolves/templ-htmx/pkg/view"
)

func main() {
	component := view.Index("Christopher")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
