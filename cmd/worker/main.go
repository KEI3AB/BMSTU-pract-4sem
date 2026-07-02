package main

import (
	"log"
	"net/http"

	"go-load-balancer/internal/config"
	"go-load-balancer/internal/worker"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", worker.HealthHandler)
	mux.HandleFunc("/work", worker.WorkHandler)

	server := &http.Server{
		Addr:    ":" + cfg.WorkerPort,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
