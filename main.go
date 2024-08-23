package main

import (
	"log"
	"net/http"

	"minecraft/handlers"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/create-server", handlers.CreateServer).Methods("POST")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
