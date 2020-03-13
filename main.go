package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	// homepage router
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and Running...")
	})
	// posts router
	router.HandleFunc("/posts", getPosts).Methods("GET")
	// add posts router
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
