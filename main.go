package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func intialzeRoter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	fmt.Printf("bps")
	InitialMigration()
	intialzeRoter()

}
