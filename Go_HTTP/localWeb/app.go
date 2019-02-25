package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"./Models/Employees.go"
)

func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
