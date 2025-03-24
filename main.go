package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Golang API!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	port := "8080"
	fmt.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
