package main

import (
	"fmt"
	"net/http"

	"github.com/michelkazi/gowebservicedemo/controllers"
)

func main() {
	port := 3000
	controllers.RegisterControllers()
	fmt.Printf("Server listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
