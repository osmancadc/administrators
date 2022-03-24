package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/administrator", GetAdministrators).Methods("GET")
	router.HandleFunc("/administrator", PostAdministrator).Methods("POST")
	router.HandleFunc("/administrator/{id}", UpdateAdministrator).Methods("PUT")

	fmt.Printf("Server listening on port %d", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}
