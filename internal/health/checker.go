package health

import (
	"net/http"
	"time"

	"go-load-balancer/internal/pool"
)

type Checker struct {
	pool   *pool.Pool
	client *http.Client
}

func NewChecker(p *pool.Pool) *Checker {
	return &Checker{
		pool: p,
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

func (c *Checker) Start() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		for _, b := range c.pool.GetAll() {
			go c.checkBackend(b)
		}
	}
}

func (c *Checker) checkBackend(b *pool.Backend) {
	healthURL := b.URL.String() + "/health"
	resp, err := c.client.Head(healthURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		b.SetAlive(false)
		return
	}
	b.SetAlive(true)
}
