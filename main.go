package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Golang API!")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "pong"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Explicitly setting 200 OK
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux() // Create a new ServeMux instance
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/ping", pingHandler)

	port := "8080"
	fmt.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux)) // Use the mux here
}
