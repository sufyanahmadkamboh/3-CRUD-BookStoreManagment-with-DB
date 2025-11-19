package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sufyanahmadkamboh/3-CRUD-BookStoreManagment-with-DB/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
