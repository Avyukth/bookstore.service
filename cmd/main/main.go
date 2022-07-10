package main

import (
	"fmt"
	"log"
	"net/http"

	"bookstore.service/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running at port 9090")
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
