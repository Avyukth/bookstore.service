package main

import (
	"github.com/avyukth/bookstore.service/routes"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.LstenAndServe("localhost:9090", r))
}
