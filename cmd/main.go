package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Omoalfa/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
