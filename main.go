package main

import (
	"fmt"
	"net/http"

	"github.com/michelkazi/gowebservicedemo/controllers"
)

func main() {
	port := 3000
	controllers.RegisterControllers()
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	fmt.Printf("Listening on port %d", port)
}
