package main

import (
	"fmt"

	"github.com/michelkazi/gowebservicedemo/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
