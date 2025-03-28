package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mabittar/go_api/router"
)

func main() {
	r := router.InitRouter()
	port := "8080"
	fmt.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
