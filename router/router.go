package router

import (
	"github.com/gorilla/mux"
	"github.com/mabittar/go_api/internal/handlers"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/", handlers.HelloHandler).Methods("GET")
	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	return r
}
