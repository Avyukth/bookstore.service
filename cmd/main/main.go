package main

import (
	"log"
	"net/http"

	"bookstore.service/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9090", router))
}
