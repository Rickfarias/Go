package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CRUD - CREATE, READ, UPDATE and DELETE
func main() {

	router := mux.NewRouter()

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}