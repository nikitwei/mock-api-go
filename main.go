package main

import (
	"github.com/gorilla/mux"
	"github.com/nikitwei/mock-api-go/controllers/authcontroller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Register).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Register).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", r))
}
