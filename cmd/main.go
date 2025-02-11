package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/agrawal-2005/bookStore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}