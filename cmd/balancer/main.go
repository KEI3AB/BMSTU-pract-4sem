package main

import (
	"log"
	"net/http"
	"net/url"

	"go-load-balancer/internal/balancer"
	"go-load-balancer/internal/config"
	"go-load-balancer/internal/health"
	"go-load-balancer/internal/pool"
)

func main() {
	cfg := config.Load()

	urls := []string{
		"http://backend1:8000",
		"http://backend2:8000",
		"http://backend3:8000",
	}

	serverPool := pool.NewPool()

	for _, u := range urls {
		parsedUrl, err := url.Parse(u)
		if err != nil {
			log.Fatal(err)
		}
		serverPool.AddBackend(pool.NewBackend(parsedUrl))
	}

	healthChecker := health.NewChecker(serverPool)
	go healthChecker.Start()

	var strategy balancer.Balancer
	if cfg.Strategy == "ROUND_ROBIN" {
		log.Println("Using Strategy: ROUND_ROBIN")
		strategy = balancer.NewRoundRobin(serverPool)
	} else {
		log.Println("Using Strategy: LEAST_CONN")
		strategy = balancer.NewLeastConn(serverPool)
	}

	proxy := balancer.NewProxy(strategy)

	server := &http.Server{
		Addr:    ":" + cfg.BalancerPort,
		Handler: proxy,
	}

	log.Printf("Load Balancer is running on port %s", cfg.BalancerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
