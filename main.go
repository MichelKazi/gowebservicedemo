package main

import (
	"net/http"

	"github.com/michelkazi/gowebservicedemo/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
